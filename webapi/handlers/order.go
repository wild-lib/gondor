package handlers

import (
	"fmt"
	"strings"

	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber"
	"github.com/satori/go.uuid"
)

func genOrder() string {
	order_no := uuid.NewV4().String()
	orderTime := randTime(365*86400, false).Unix() * 1000
	salesman := fmt.Sprintf("%s %s", faker.FirstName(), faker.LastName())
	price := randFloat(500, 90000, true)
	status := randItem([]string{"pending", "payed", "finish"})
	return ReduceBlanks(fmt.Sprintf(`{
	"order_no": "%s",
	"timestamp": %d,
	"username": "%s",
	"price": %.1f,
	"status": "%s"
}`, order_no, orderTime, salesman, price, status))
}

// 订单列表
func OrderListHandler(ctx *fiber.Ctx) {
	ctx.Type("json").SendBytes([]byte(ReduceBlanks(`{"code":200, "total":20, "data":[` + genOrder() + `]}`)))
}

// 查找用户名
func SearchUserHandler(ctx *fiber.Ctx) {
	var names []string
	match := strings.ToLower(ctx.Query("name"))
	for _, name := range fakeUsers {
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
	ctx.JSON(result)
}
