// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Operation is the golang structure of table operation for DAO operations like Where/Data.
type Operation struct {
	g.Meta   `orm:"table:operation, do:true"`
	Id       interface{} // pk
	Uid      interface{} // User ID
	OptType  interface{} // Operate Type
	Score    interface{} // Score
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
