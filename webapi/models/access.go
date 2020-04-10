package models

import (
	"strings"
	"time"

	base "github.com/azhai/gozzo-db/construct"
)

// 操作列表
const (
	ACCESS_VIEW uint16 = 2 << iota
	ACCESS_DISABLE
	ACCESS_REMOVE
	ACCESS_EDIT
	ACCESS_CREATE
	ACCESS_GET
	ACCESS_POST
	ACCESS_GRANT
	ACCESS_ALL
	ACCESS_NONE uint16 = 0 // 无权限
)

var (
	AccessNames = map[uint16]string{
		ACCESS_VIEW: "view", ACCESS_DISABLE: "disable", ACCESS_REMOVE: "remove",
		ACCESS_EDIT: "edit", ACCESS_CREATE: "create", ACCESS_GET: "get", ACCESS_POST: "post",
		ACCESS_GRANT: "grant", ACCESS_ALL: "all", ACCESS_NONE: "",
	}
	AccessTitles = map[uint16]string{
		ACCESS_VIEW: "查看", ACCESS_DISABLE: "禁用", ACCESS_REMOVE: "删除",
		ACCESS_EDIT: "编辑", ACCESS_CREATE: "新建", ACCESS_GET: "GET", ACCESS_POST: "POST",
		ACCESS_GRANT: "授权", ACCESS_ALL: "全部", ACCESS_NONE: "无",
	}
)

// 查询符合条件的所有行
func (m Access) FindAll(filters ...base.FilterFunc) (objs []Access, err error) {
	err = db.Model(m).Scopes(filters...).Find(&objs).Error
	err = IgnoreNotFoundError(err)
	return
}

// 查询符合条件的第一行
func (m Access) GetOne(filters ...base.FilterFunc) (obj *Access, err error) {
	obj = new(Access)
	err = db.Model(m).Scopes(filters...).Take(&obj).Error
	err = IgnoreNotFoundError(err)
	return
}

// 添加权限
func AddAccess(role, res string, perm uint16, args ...string) (access *Access, err error) {
	access = &Access{RoleName: role, ResourceType: res, PermCode: perm}
	_, names := ParsePermNames(access.PermCode)
	access.Actions = strings.Join(names, ",")
	if len(args) > 0 {
		resArgs := strings.Join(args, ",")
		access.ResourceArgs = &resArgs
	}
	access.GrantedAt = time.Now()
	err = db.Create(access).Error
	return
}

// 分解出具体权限
func ParsePermNames(perm uint16) (codes []uint16, names []string) {
	for code, name := range AccessNames {
		if code > 0 && perm&code == code {
			codes = append(codes, code)
			names = append(names, name)
		}
	}
	return
}

// 找出权限的中文名称
func GetPermTitles(codes []uint16) (titles []string) {
	title, ok := "", false
	for _, code := range codes {
		if title, ok = AccessTitles[code]; !ok {
			title = "未知"
		}
		titles = append(titles, title)
	}
	return
}
