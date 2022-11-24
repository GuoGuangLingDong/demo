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
	g.Meta      `orm:"table:operation, do:true"`
	Uid         interface{} // User ID
	OperateCode interface{} // Operate Code
	OperateTime *gtime.Time // Operate Time
	Score       interface{} // Score
}
