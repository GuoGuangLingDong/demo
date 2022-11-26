// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Favor is the golang structure of table favor for DAO operations like Where/Data.
type Favor struct {
	g.Meta    `orm:"table:favor, do:true"`
	Uid       interface{} // User ID
	PoapId    interface{} // Poap id
	FavorTime *gtime.Time // Favor Time
}