package models

import (
	base "github.com/azhai/gozzo-db/construct"
	"github.com/jinzhu/gorm"
)

// 查询符合条件的所有行
func (m UserRole) FindAll(filters ...base.FilterFunc) (objs []*UserRole, err error) {
	err = db.Model(m).Scopes(filters...).Find(&objs).Error
	err = IgnoreNotFoundError(err)
	return
}

// 查询符合条件的第一行
func (m UserRole) GetOne(filters ...base.FilterFunc) (obj *UserRole, err error) {
	obj = new(UserRole)
	err = db.Model(m).Scopes(filters...).Take(&obj).Error
	err = IgnoreNotFoundError(err)
	return
}

// 查询某个用户的所有角色名
func (m UserRole) GetUserRoles(uid string) (roles []string) {
	_ = db.Model(m).Where("user_uid = ?", uid).Pluck("role_name", &roles).Error
	return
}

// 查询属于某个角色的所有用户
func (m UserRole) GetRoleUsers(roleName string) (users []*User) {
	var uids []string
	_ = db.Model(m).Where("role_name", roleName).Pluck("user_uid", &uids).Error
	if len(uids) > 0 {
		users, _ = new(User).FindAll(func(query *gorm.DB) *gorm.DB {
			return query.Where("uid IN (?)", uids)
		})
	}
	return
}
