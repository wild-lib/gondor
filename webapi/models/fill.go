package models

import (
	"github.com/azhai/gozzo-db/export"
	"github.com/jinzhu/gorm"
)

// 写入必须的初始化数据
func FillRequiredData(drv string, db *gorm.DB) *gorm.DB {
	export.LoadFileData(db, "data/db_test.toml", ModelInsts, true)
	// 菜单
	if menu, _ := new(Menu).GetOne(); menu.Path == "" {
		_, _ = AddMenu("/dashboard", "面板", "dashboard", nil)
		menu, _ = AddMenu("/permission", "权限", "lock", nil)
		_, _ = AddMenu("role", "角色权限", "", menu)
		menu, _ = AddMenu("/table", "Table", "table", nil)
		_, _ = AddMenu("complex-table", "复杂Table", "", menu)
		_, _ = AddMenu("inline-edit-table", "内联编辑", "", menu)
		menu, _ = AddMenu("/excel", "Excel", "excel", nil)
		_, _ = AddMenu("export-selected-excel", "选择导出", "", menu)
		_, _ = AddMenu("upload-excel", "上传Excel", "", menu)
		menu, _ = AddMenu("/theme/index", "主题", "theme", nil)
		menu, _ = AddMenu("/error/404", "404错误", "404", nil)
		menu, _ = AddMenu("https://cn.vuejs.org/", "外部链接", "link", nil)
	}
	// 权限
	if access, _ := new(Access).GetOne(); access.ID == 0 {
		AddAccess("superuser", "menu", ACCESS_ALL, "*") // 超管可以访问所有菜单
		// 基本用户
		AddAccess("visitor", "menu", ACCESS_VIEW, "/dashboard")
		AddAccess("visitor", "menu", ACCESS_VIEW, "/error/404")
	}
	return db
}
