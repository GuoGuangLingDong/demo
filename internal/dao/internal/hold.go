// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HoldDao is the data access object for table hold.
type HoldDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns HoldColumns // columns contains all the column names of Table for convenient usage.
}

// HoldColumns defines and stores column names for table hold.
type HoldColumns struct {
	Id       string // pk
	Uid      string // User ID
	PoapId   string // Poap id
	TokenId  string // Poap tokenId
	CreateAt string // Created Time
	UpdateAt string // Updated Time
}

// holdColumns holds the columns for table hold.
var holdColumns = HoldColumns{
	Id:       "id",
	Uid:      "uid",
	PoapId:   "poap_id",
	TokenId:  "token_id",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewHoldDao creates and returns a new DAO object for table data access.
func NewHoldDao() *HoldDao {
	return &HoldDao{
		group:   "default",
		table:   "hold",
		columns: holdColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HoldDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HoldDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HoldDao) Columns() HoldColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HoldDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HoldDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HoldDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
