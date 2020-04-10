package models

import (
	base "github.com/azhai/gozzo-db/construct"
)

// 查询符合条件的所有行
func (m Menu) FindAll(filters ...base.FilterFunc) (objs []*Menu, err error) {
	err = db.Model(m).Scopes(filters...).Find(&objs).Error
	err = IgnoreNotFoundError(err)
	return
}

// 查询符合条件的第一行
func (m Menu) GetOne(filters ...base.FilterFunc) (obj *Menu, err error) {
	obj = new(Menu)
	err = db.Model(m).Scopes(filters...).Take(&obj).Error
	err = IgnoreNotFoundError(err)
	return
}

// 添加一个菜单
func (m *Menu) AddTo(parent *Menu) (err error) {
	var parentNode *base.NestedModel
	if parent != nil {
		parentNode = parent.NestedModel
	}
	if m.NestedModel == nil {
		m.NestedModel = new(base.NestedModel)
	}
	query := db.Table(m.TableName())
	err = m.NestedModel.AddToParent(parentNode, query)
	if err == nil {
		err = db.Create(m).Error
	}
	return
}

// 添加菜单
func AddMenu(path, title string, icon string, parent *Menu) (menu *Menu, err error) {
	menu = &Menu{Path: path, Title: title, Icon: icon}
	err = menu.AddTo(parent)
	return
}
