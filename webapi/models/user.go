package models

import (
	"encoding/hex"
	"time"

	base "github.com/azhai/gozzo-db/construct"
	"github.com/azhai/gozzo-utils/cryptogy"
	"github.com/jinzhu/gorm"
	"github.com/muyo/sno"
)

// 查询符合条件的所有行
func (m User) FindAll(filters ...base.FilterFunc) (objs []*User, err error) {
	err = db.Model(m).Scopes(filters...).Find(&objs).Error
	err = IgnoreNotFoundError(err)
	return
}

// 查询符合条件的第一行
func (m User) GetOne(filters ...base.FilterFunc) (obj *User, err error) {
	obj = new(User)
	err = db.Model(m).Scopes(filters...).Take(&obj).Error
	err = IgnoreNotFoundError(err)
	return
}

// 是否存在
func (m User) IsExists(field, value string) bool {
	var count int
	err := db.Model(m).Where(field, value).Count(&count).Error
	return err == nil && count > 0
}

func NewUser(username, realname string) *User {
	user := &User{Username: username, Realname: realname}
	user.UID = sno.NewWithTime('U', user.CreatedAt).String()
	user.CreatedAt = time.Now()
	return user
}

// 8位salt值，用$符号分隔开
var saltPasswd = cryptogy.NewSaltPassword(8, "$")

// 设置密码
func (m *User) SetPassword(password string) *User {
	m.Password = saltPasswd.CreatePassword(password)
	return m
}

// 校验密码
func (m User) VerifyPassword(password string) bool {
	return saltPasswd.VerifyPassword(password, m.Password)
}

// 登录
func (m User) Signin(username, password string) (user *User, token string, err error) {
	user, err = m.GetOne(func(query *gorm.DB) *gorm.DB {
		return query.Where("username = ?", username)
	})
	if err == nil && user.VerifyPassword(password) {
		ticket := sno.New('T').Bytes()[:8]
		tailno := cryptogy.RandSalt(2)
		token = hex.EncodeToString(append(ticket, tailno...)) // 生成token
	}
	return
}
