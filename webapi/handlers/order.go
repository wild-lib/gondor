package handlers

import (
	"strings"

	"github.com/astro-bug/gondor/webapi/fakes"
	"gitee.com/azhai/fiber-u8l/v2"
)

// 订单列表
func OrderListHandler(ctx *fiber.Ctx) (err error) {
	result := fakes.ReduceBlanks(`{"code":200, "total":20, "data":[` + fakes.GenOrder() + `]}`)
	err = ctx.Type("json").Send([]byte(result))
	return
}

// 查找用户名
func SearchUserHandler(ctx *fiber.Ctx) (err error) {
	var names []string
	match := strings.ToLower(ctx.Query("name"))
	for _, name := range fakes.FakeUsers {
		lowerName := strings.ToLower(name)
		if strings.Contains(lowerName, match) {
			names = append(names, name)
		}
	}
	result := fiber.Map{
		"code": 200,
		"data": fiber.Map{
			"items": names,
		},
	}
	err = ctx.JSON(result)
	return
}
