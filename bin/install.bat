@ECHO OFF

REM 安装开发用ssl证书
mkcert.exe -install
mkcert.exe -key-file certs\key.pem -cert-file certs\cert.pem localhost 127.0.0.1

REM 安装并启动Windows服务
winsw.exe install
winsw.exe start

PAUSE