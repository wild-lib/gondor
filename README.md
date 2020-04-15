# gondor （Go + Vue 管理后台）

## 项目概述

* 后端使用 Go 语言 caddy + fiber + gorm （将来可能换成 xorm）
* 前端使用 NodeJS 的 vue-element-admin
* 用到 Mariadb 数据库和 Redis 缓存

需要先安装 go nodejs mariadb/mysql redis ；

请修改 settings.toml 中的数据库和缓存配置，将 db_test.sql 导入到对应数据库中；

如果修改服务的端口，请修改 Caddyfile  默认 Web(gondor) 使用 8080 端口， API(rohan) 使用 8000 端口。

账号为 **admin** 密码是 **654321**

## 编译重要提示

由于 etcd 和 grpc 的版本冲突，后面两节中的 make all 不可用，需要使用以下步骤代替：

0. 更新依赖库（可选）

```bash
# 将大部分依赖库升级到最新发行版本，即 latest
# caddy, fiber, rpcx, sno 升级到最新主干版本，即 master
# grpc 固定在 v1.26.0 版本，因为 etcd 和它的新版本不兼容
gawk -f go_mod_up.awk go.mod
```

1. 打开 go.mod 文件，将 replace 这一段先注释掉

> // replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

2. 接着编译 gondor 

```bash
go build -ldflags="-s -w" ./cmd/gondor/
```

3. 再将注释去掉，然后编译 moria 和 rohan （双击运行 make.bat 或者命令行下执行 make ）

## Windows 下的安装与编译

（默认项目文件夹 D:\gondor ）

1. 第一次运行，在 Dos 下进入项目文件夹，执行 make all 生成 gondor.exe 以及 moria.exe 和 rohan.exe ，

	以后只需要在 Windows 下双击运行 make.bat ，重新生成 moria.exe 和 rohan.exe；
	
2. 同样，在 Dos 下进入 website 文件夹，第一次执行 npm install 安装依赖包，可能还需要运行 npm audit fix ，
	
	在 Windows 下双击运行 build.bat ，生成静态文件在 website/dist 文件夹；
	
3. 进入 bin 文件夹，修改 winsw.xml 中的项目文件夹位置，双击运行 install.bat ，

    生成开发用的ssl证书，并将 gondor.exe 安装为 Windows 服务和启动服务；

4. 核对和修改 settings.yml 中的数据库、缓存连接参数，在 Dos 下运行接口程序 rohan.exe ，

    然后在浏览器中打开 https://127.0.0.1:8080/ （注意您是否有更换端口）

## Linux/MacOS 下的安装与编译

（默认项目文件夹 /var/projects/gondor ）

1. 使用 [mkcert](https://github.com/FiloSottile/mkcert) 在 bin/certs/ 下生成开发用ssl证书

2. 将 gondor 设置为系统服务，可保持服务在后台运行，并且方便管理

   2.1 Linux systemd 服务

```bash
cat > /etc/systemd/system/gondor.service <<EOD
[Unit]
Description=Gondor Web Server
After=syslog.target network.target

[Service]
ExecStart=/var/projects/gondor/gondor run
WorkingDirectory=/var/projects/gondor
#PIDFile=/var/run/gondor.pid
LimitNOFILE=819200
LimitNPROC=819200
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=gondor
Restart=always

[Install]
WantedBy=multi-user.target
EOD

systemctl daemon-reload
```

​	2.2 MacOS launchd 服务

```zsh
sudo cat > /Library/LaunchAgents/my.gondor.plist <<EOD
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>Label</key>
  <string>my.gondor</string>
  <key>WorkingDirectory</key>
  <string>/var/projects/gondor</string>
  <key>ProgramArguments</key>
  <array>
    <string>/var/projects/gondor/gondor</string>
    <string>run</string>
  </array>
  <key>RunAtLoad</key>
  <true/>
  <key>KeepAlive</key>
  <false/>
</dict>
</plist>
EOD

sudo chown root:wheel /Library/LaunchAgents/my.gondor.plist
sudo chmod 600 /Library/LaunchAgents/my.gondor.plist
sudo launchctl load /Library/LaunchAgents/my.gondor.plist
```

3. 在项目文件夹生成 Web Server 和 Web API 程序，启动 rohan 之前，

   先核对和修改 settings.yml 中的数据库、缓存连接参数
```bash
#每次重新编译 gondor 执行
make serv && systemctl restart gondor
#每次重新编译 rohan 执行
make && ./rohan -p 8000 -v
```

4. 在 website 文件夹下安装依赖包和生成静态文件
```bash
cd /var/projects/gondor/website/
npm install
npm audit fix #可能需要运行
npm run build:prod
```
   在浏览器中打开 https://127.0.0.1:8080/ （注意您是否有更换端口）

## Linux （以 CentOS 7 为例）下 可选设置

### 1. 使用 openresty/nginx 代替 gondor 作为 Web Server 

配置文件中单个站点配置如下：

```
upstream rohan_api {
	server 127.0.0.1:8000 weight=10;
}
map $http_upgrade $connection_upgrade {
	default upgrade;
	'' close;
}

server {
    listen           8080;
    server_name      127.0.0.1;
    root             /var/projects/gondor/website/dist;
    index            index.html;
    access_log       off;
    error_page  404  /404.html;
    
    location / {
		try_files  $uri /index.html @websocket;
		access_log logs/gondor.access.log  main;
    }
    location ~ \.(svn|git|hg|bzr|cvs) {
        return 404;
    }
    
    ## 随机图片， nginx 需要安装相关的 redis 模块才能使用
    #location = /image/random/ {
    #    access_by_lua_block {
	#		local redis = require "resty.redis-util"
	#		local red = redis.new()
	#		local img_url, err = red:srandmember("posters:1090:300")
	#		return ngx.redirect("/posters/" .. img_url, 302)
    #    }
    #}
    
    location = /ws {
		try_files  $uri @websocket;
		access_log logs/rohan.access.log  main;
    }
    location ~ ^/(api|ws)/ {
		try_files  $uri @websocket;
		access_log logs/rohan.access.log  main;
    }
    location @websocket {
		proxy_pass http://rohan_api;
		proxy_http_version 1.1;
		proxy_set_header Upgrade $http_upgrade;
		proxy_set_header Connection $connection_upgrade;
		proxy_set_header Origin "";
		proxy_set_header Cookie $http_cookie;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

### 2. syslog Linux 系统日志配置

```
cat > /etc/rsyslog.d/daemon.conf <<EOD
#*.*;daemon.none,auth,authpriv.none     /var/log/syslog
#daemon.*                               -/var/log/daemon.log
:app-name, isequal, "gondor"           -/var/log/gondor.log
:app-name, isequal, "rohan"            -/var/log/rohan.log
EOD
```

### 3. 加大 CentOS7 系统资源配置

```
cat >> /etc/security/limits.conf <<EOD

#<domain>      <type>  <item>         <value>
*        soft    nofile        819200
*        hard    nofile        819200
root     soft    nofile        819200
root     hard    nofile        819200

#<domain>      <type>  <item>         <value>
*        soft    nproc         819200
*        hard    nproc         819200
root     soft    nproc         819200
root     hard    nproc         819200

#<domain>      <type>  <item>         <value>
*        soft    sigpending         409600
*        hard    sigpending         409600
root     soft    sigpending         409600
root     hard    sigpending         409600
EOD

ulimit -SHn 819200 && ulimit -SHu 819200 && ulimit -SHi 409600


cat >> /etc/sysctl.conf <<EOD

fs.file-max = 819200
vm.max_map_count = 819200
kernel.pid_max = 204800
kernel.sysrq = 1

net.core.netdev_max_backlog = 32000
net.core.rmem_max = 16777216
net.core.somaxconn = 8192
net.core.wmem_max = 16777216

net.ipv4.conf.all.arp_announce=2
net.ipv4.conf.all.rp_filter=0
net.ipv4.conf.all.send_redirects = 1
net.ipv4.conf.default.arp_announce = 2
net.ipv4.conf.default.rp_filter=0
net.ipv4.conf.default.send_redirects = 1
net.ipv4.conf.lo.arp_announce=2

net.ipv4.ip_forward = 1
net.ipv4.ip_local_port_range = 5001  65535
net.ipv4.icmp_echo_ignore_broadcasts = 1 # 避免放大攻击
net.ipv4.icmp_ignore_bogus_error_responses = 1 # 开启恶意icmp错误消息保护

net.ipv4.tcp_fin_timeout = 30
net.ipv4.tcp_keepalive_time = 1800
net.ipv4.tcp_max_syn_backlog = 1024
net.ipv4.tcp_max_syn_backlog = 8192
net.ipv4.tcp_max_tw_buckets = 5000
net.ipv4.tcp_rmem = 4096 87380 16777216

net.ipv4.tcp_synack_retries = 2
net.ipv4.tcp_syncookies = 1
net.ipv4.tcp_timestamps = 1
#net.ipv4.tcp_tw_recycle = 1
net.ipv4.tcp_tw_reuse = 1
#net.ipv4.tcp_tw_timeout = 3
net.ipv4.tcp_wmem = 4096 65536 16777216
EOD

sysctl -p
```