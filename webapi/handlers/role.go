package handlers

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/astro-bug/gondor/webapi/fakes"
	"github.com/astro-bug/gondor/webapi/models/db"
	"github.com/gofiber/fiber/v2"
)

// 完整菜单
func AllMenuHandler(ctx *fiber.Ctx) {
	routes := append(append(constantRoutes, permRoute), asyncRoutes...)
	ctx.Type("json").SendBytes([]byte(`{"code":200, "data":[` + fakes.ReduceBlanks(strings.Join(routes, ",")) + `]}`))
}

// 角色列表
func RoleListHandler(ctx *fiber.Ctx) {
	superRoutes := append(append(constantRoutes, permRoute), asyncRoutes...)
	editorRoutes := append(constantRoutes, asyncRoutes...)
	DefaultRoutes := `{
  "path": "",
  "redirect": "dashboard",
  "children": [
    {
      "path": "dashboard",
      "name": "Dashboard",
      "meta": { "title": "面板", "icon": "dashboard" }
    }
  ]
}`
	var buf bytes.Buffer
	roles, _ := db.GetAllRoles()
	for i, role := range roles {
		if i > 0 {
			buf.WriteString(",\n")
		}
		zhName := strings.SplitN(role.Remark, "，", 2)[0]
		routes := DefaultRoutes // 避免当前值带入下一次循环
		if role.Name == "superuser" {
			routes = strings.Join(superRoutes, ",")
		} else if role.Name == "editor" {
			routes = strings.Join(editorRoutes, ",")
		}
		buf.WriteString(fmt.Sprintf(`{
    "key": "%s",
    "name": "%s",
    "description": "%s",
    "routes": [%s]
  }`, role.Name, zhName, role.Remark, routes))
	}
	data := fakes.ReduceBlanks(buf.String())
	ctx.Type("json").SendBytes([]byte(`{"code":200, "data":[` + data + `]}`))
}

// 添加角色
func RoleAddHandler(ctx *fiber.Ctx) {
	result := fiber.Map{
		"code": 200,
		"data": fiber.Map{
			"key": fakes.RandInt(300, 5000),
		},
	}
	ctx.JSON(result)
}

// 修改或删除角色
func RoleModHandler(ctx *fiber.Ctx) {
	result := fiber.Map{
		"code": 200,
		"data": fiber.Map{
			"status": "success",
		},
	}
	ctx.JSON(result)
}
