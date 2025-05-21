// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserGroupDao is the data access object for the table user_group.
type UserGroupDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserGroupColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserGroupColumns defines and stores column names for the table user_group.
type UserGroupColumns struct {
	UserId  string //
	GroupId string //
}

// userGroupColumns holds the columns for the table user_group.
var userGroupColumns = UserGroupColumns{
	UserId:  "user_id",
	GroupId: "group_id",
}

// NewUserGroupDao creates and returns a new DAO object for table data access.
func NewUserGroupDao(handlers ...gdb.ModelHandler) *UserGroupDao {
	return &UserGroupDao{
		group:    "default",
		table:    "user_group",
		columns:  userGroupColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserGroupDao) Columns() UserGroupColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserGroupDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
