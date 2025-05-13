// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserInfoDao is the data access object for the table user_info.
type UserInfoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserInfoColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserInfoColumns defines and stores column names for the table user_info.
type UserInfoColumns struct {
	UserId   string //
	Username string //
	IdNumber string // 学号/工号
	Avatar   string //
}

// userInfoColumns holds the columns for the table user_info.
var userInfoColumns = UserInfoColumns{
	UserId:   "user_id",
	Username: "username",
	IdNumber: "id_number",
	Avatar:   "avatar",
}

// NewUserInfoDao creates and returns a new DAO object for table data access.
func NewUserInfoDao(handlers ...gdb.ModelHandler) *UserInfoDao {
	return &UserInfoDao{
		group:    "default",
		table:    "user_info",
		columns:  userInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserInfoDao) Columns() UserInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserInfoDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
