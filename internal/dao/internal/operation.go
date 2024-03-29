// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OperationDao is the data access object for table operation.
type OperationDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns OperationColumns // columns contains all the column names of Table for convenient usage.
}

// OperationColumns defines and stores column names for table operation.
type OperationColumns struct {
	Id        string // pk
	Uid       string // User ID
	OptType   string // Operate Type
	Score     string // Score
	CreateAt  string // Created Time
	UpdateAt  string // Updated Time
	OverdueAt string //
}

// operationColumns holds the columns for table operation.
var operationColumns = OperationColumns{
	Id:        "id",
	Uid:       "uid",
	OptType:   "opt_type",
	Score:     "score",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	OverdueAt: "overdue_at",
}

// NewOperationDao creates and returns a new DAO object for table data access.
func NewOperationDao() *OperationDao {
	return &OperationDao{
		group:   "default",
		table:   "operation",
		columns: operationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OperationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OperationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OperationDao) Columns() OperationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OperationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OperationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OperationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
