// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMenuOld is the golang structure of table hg_admin_menu_old for DAO operations like Where/Data.
type AdminMenuOld struct {
	g.Meta    `orm:"table:hg_admin_menu_old, dto:true"`
	Id        interface{} // 菜单ID
	Pid       interface{} // 父菜单ID
	Name      interface{} // 菜单名称
	Icon      interface{} // 菜单图标
	Type      interface{} // 菜单类型（M目录 C菜单 F按钮）
	Perms     interface{} // 权限标识
	Path      interface{} // 路由地址
	Component interface{} // 组件路径
	Query     interface{} // 路由参数
	IsFrame   interface{} // 是否为外链（0是 1否）
	IsCache   interface{} // 是否缓存（0缓存 1不缓存）
	IsVisible interface{} // 菜单状态（0显示 1隐藏）
	Remark    interface{} // 备注
	Level     interface{} // 级别
	Tree      interface{} // 树
	Sort      interface{} // 排序
	Status    interface{} // 菜单状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
