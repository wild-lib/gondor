package main

import (
	"flag"
	"fmt"

	"github.com/azhai/gondor/webapi"
	"github.com/gofiber/compression"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

var (
	servAddr         string
	port             uint   // 运行端口
	verbose          bool   // 详细输出
)

func init() {
	flag.UintVar(&port, "p", 8000, "运行端口")
	flag.BoolVar(&verbose, "v", false, "输出详细信息")
	flag.Parse()
	servAddr = fmt.Sprintf(":%d", port)
}

func main() {
	app := fiber.New()
	app.Use(compression.New()).Use(cors.New())
	webapi.AddRoutes(app.Group("/api"))
	app.Listen(servAddr)
}
