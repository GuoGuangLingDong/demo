// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Hold is the golang structure of table hold for DAO operations like Where/Data.
type Hold struct {
	g.Meta   `orm:"table:hold, do:true"`
	Id       interface{} // pk
	Uid      interface{} // User ID
	PoapId   interface{} // Poap id
	TokenId  interface{} // Poap tokenId
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
