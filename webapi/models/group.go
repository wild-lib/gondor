package models

import (
	"time"

	base "github.com/azhai/gozzo-db/construct"
	"github.com/muyo/sno"
)

// 查询符合条件的所有行
func (m Group) FindAll(filters ...base.FilterFunc) (objs []Group, err error) {
	err = db.Model(m).Scopes(filters...).Find(&objs).Error
	err = IgnoreNotFoundError(err)
	return
}

// 查询符合条件的第一行
func (m Group) GetOne(filters ...base.FilterFunc) (obj *Group, err error) {
	obj = new(Group)
	err = db.Model(m).Scopes(filters...).Take(&obj).Error
	err = IgnoreNotFoundError(err)
	return
}

func NewGroup(title string) *Group {
	group := &Group{Title: title, CreatedAt: time.Now()}
	group.GID = sno.NewWithTime('G', group.CreatedAt).String()
	return group
}
