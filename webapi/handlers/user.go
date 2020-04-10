package handlers

import (
	"net/http"

	"github.com/azhai/gondor/webapi/models"
	"github.com/azhai/gondor/webapi/utils"
	"github.com/azhai/gozzo-db/session"
	"github.com/gofiber/fiber"
)

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserInfo(user *models.User) map[string]interface{} {
	info := map[string]interface{}{
		"username": user.Username,
	}
	if user.Realname != "" {
		info["realname"] = user.Realname
	}
	if user.Mobile != "" {
		info["mobile"] = user.Mobile
	}
	if user.Avatar != nil {
		info["avatar"] = *user.Avatar
	}
	if user.Introduction != nil {
		info["introduction"] = *user.Introduction
	}
	return info
}

// 用户登录
func UserLoginHandler(ctx *fiber.Ctx) {
	// 获取参数
	var data UserData
	if err := ctx.BodyParser(&data); err != nil {
		ctx.JSON(fiber.Map{
			"code":    510,
			"message": err.Error(),
		})
	}
	// 查询数据
	user, token, err := new(models.User).Signin(data.Username, data.Password)
	if err != nil || token == "" {
		ctx.JSON(fiber.Map{
			"code":    510,
			"message": "失败，密码不正确！",
		})
		return
	}
	// 写入Session
	roles := new(models.UserRole).GetUserRoles(user.UID)
	sess := models.Session(token)
	sess.BindRoles(user.UID, roles, true)
	sess.SaveMap(GetUserInfo(user), false)
	ctx.JSON(fiber.Map{
		"code": 200,
		"data": fiber.Map{
			"token": token,
		},
	})
}

// 用户退出
func UserLogoutHandler(ctx *fiber.Ctx) {
	token := ctx.Cookies("access_token")
	if token == "" {
		utils.Abort(ctx, http.StatusUnauthorized)
		return
	}
	models.Registry().DelSession(token)
	result := fiber.Map{
		"code": 200,
		"data": "成功退出",
	}
	ctx.JSON(result)
}

// 用户资料
func UserInfoHandler(ctx *fiber.Ctx) {
	token := ctx.Query("token")
	sess := models.Session(token)
	if uid, err := sess.GetString("uid"); err != nil || uid == "" {
		result := fiber.Map{
			"code":    508,
			"message": "失败，用户不存在！",
		}
		ctx.JSON(result)
	} else {
		data, _ := sess.GetAllString()
		result := fiber.Map{
			"code": 200,
			"data": fiber.Map{
				"roles":        session.SessListSplit(data["roles"]),
				"name":         data["name"],
				"avatar":       data["avatar"],
				"introduction": data["introduction"],
			},
		}
		ctx.JSON(result)
	}
}
