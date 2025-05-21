// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProblemDao is the data access object for the table problem.
type ProblemDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ProblemColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ProblemColumns defines and stores column names for the table problem.
type ProblemColumns struct {
	Id                  string //
	Title               string // 简单的题目
	Topic               string // 详细题目
	Category            string // 1 html
	Difficulty          string // 1 入门 2简单 3中等 4难 5困难
	Answer              string //
	PassingNumber       string // 通过人数
	ParticipationNumber string // 参与人数
	CreatedAt           string //
	UpdatedAt           string //
	DeletedAt           string //
}

// problemColumns holds the columns for the table problem.
var problemColumns = ProblemColumns{
	Id:                  "id",
	Title:               "title",
	Topic:               "topic",
	Category:            "category",
	Difficulty:          "difficulty",
	Answer:              "answer",
	PassingNumber:       "passing_number",
	ParticipationNumber: "participation_number",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	DeletedAt:           "deleted_at",
}

// NewProblemDao creates and returns a new DAO object for table data access.
func NewProblemDao(handlers ...gdb.ModelHandler) *ProblemDao {
	return &ProblemDao{
		group:    "default",
		table:    "problem",
		columns:  problemColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ProblemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ProblemDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ProblemDao) Columns() ProblemColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ProblemDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ProblemDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ProblemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
