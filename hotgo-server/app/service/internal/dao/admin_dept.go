// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/bufanyun/hotgo/app/service/internal/dao/internal"
	"github.com/gogf/gf/v2/errors/gerror"
)

// adminDeptDao is the data access object for table hg_admin_dept.
// You can define custom methods on it to extend its functionality as you wish.
type adminDeptDao struct {
	*internal.AdminDeptDao
}

var (
	// AdminDept is globally public accessible object for table hg_admin_dept operations.
	AdminDept = adminDeptDao{
		internal.NewAdminDeptDao(),
	}
)

//
//  @Title  判断名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   id
//  @Param   name
//  @Return  bool
//  @Return  error
//
func (dao *adminDeptDao) IsUniqueName(ctx context.Context, id int64, name string) (bool, error) {
	var data *entity.AdminDept
	m := dao.Ctx(ctx).Where("name", name)

	if id > 0 {
		m = m.WhereNot("id", id)
	}

	if err := m.Scan(&data); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return false, err
	}

	if data == nil {
		return true, nil
	}

	return false, nil
}

//
//  @Title  获取最上级pid
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   data
//  @Return  int64
//  @Return  error
//
func (dao *adminDeptDao) TopPid(ctx context.Context, data *entity.AdminDept) (int64, error) {
	var pidData *entity.AdminDept
	if data.Pid == 0 {
		return data.Id, nil
	}
	err := dao.Ctx(ctx).Where("id", data.Pid).Scan(&pidData)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return 0, err
	}

	return dao.TopPid(ctx, pidData)
}

// Fill with you ideas below.
