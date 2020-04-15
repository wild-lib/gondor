module github.com/astro-bug/gondor

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	gitea.com/lunny/log v0.0.0-20190322053110-01b5df579c4e
	github.com/azhai/gozzo-db v0.6.11
	github.com/azhai/gozzo-utils v0.4.2
	github.com/bxcodec/faker/v3 v3.3.0
	github.com/caddyserver/caddy/v2 v2.0.0-rc.3.0.20200414221146-829e36d535cf
	github.com/denisenkom/go-mssqldb v0.0.0-20200206145737-bbfc9a55622e // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gobwas/glob v0.2.3
	github.com/gofiber/compression v0.0.3
	github.com/gofiber/cors v0.0.3
	github.com/gofiber/fiber v1.9.1-0.20200415080649-5b1a367505cc
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/muyo/sno v1.1.1-0.20200406142550-e5d36f06b5d6
	github.com/satori/go.uuid v1.2.0
	github.com/smallnest/rpcx v0.0.0-20200414114925-bff251b691b9
	github.com/spf13/cobra v1.0.0 // indirect
	github.com/urfave/cli v1.22.4
	golang.org/x/crypto v0.0.0-20200414173820-0848c9571904 // indirect
	gopkg.in/yaml.v2 v2.2.8
	xorm.io/builder v0.3.7
	xorm.io/reverse v0.0.0-20200323084835-27bc196aa762
	xorm.io/xorm v1.0.1
)
