// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminMenuOldDao is the data access object for table hg_admin_menu_old.
type AdminMenuOldDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns AdminMenuOldColumns // columns contains all the column names of Table for convenient usage.
}

// AdminMenuOldColumns defines and stores column names for table hg_admin_menu_old.
type AdminMenuOldColumns struct {
	Id        string // 菜单ID
	Pid       string // 父菜单ID
	Name      string // 菜单名称
	Icon      string // 菜单图标
	Type      string // 菜单类型（M目录 C菜单 F按钮）
	Perms     string // 权限标识
	Path      string // 路由地址
	Component string // 组件路径
	Query     string // 路由参数
	IsFrame   string // 是否为外链（0是 1否）
	IsCache   string // 是否缓存（0缓存 1不缓存）
	IsVisible string // 菜单状态（0显示 1隐藏）
	Remark    string // 备注
	Level     string // 级别
	Tree      string // 树
	Sort      string // 排序
	Status    string // 菜单状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

//  adminMenuOldColumns holds the columns for table hg_admin_menu_old.
var adminMenuOldColumns = AdminMenuOldColumns{
	Id:        "id",
	Pid:       "pid",
	Name:      "name",
	Icon:      "icon",
	Type:      "type",
	Perms:     "perms",
	Path:      "path",
	Component: "component",
	Query:     "query",
	IsFrame:   "is_frame",
	IsCache:   "is_cache",
	IsVisible: "is_visible",
	Remark:    "remark",
	Level:     "level",
	Tree:      "tree",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAdminMenuOldDao creates and returns a new DAO object for table data access.
func NewAdminMenuOldDao() *AdminMenuOldDao {
	return &AdminMenuOldDao{
		group:   "default",
		table:   "hg_admin_menu_old",
		columns: adminMenuOldColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminMenuOldDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminMenuOldDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminMenuOldDao) Columns() AdminMenuOldColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminMenuOldDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminMenuOldDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminMenuOldDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
