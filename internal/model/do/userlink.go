// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Userlink is the golang structure of table userlink for DAO operations like Where/Data.
type Userlink struct {
	g.Meta    `orm:"table:userlink, do:true"`
	Id        interface{} // pk
	Uid       interface{} // User ID
	Link      interface{} // Link
	LinkType  interface{} // Link type
	CreateAt  *gtime.Time // Created Time
	UpdateAt  *gtime.Time // Updated Time
	LinkTitle interface{} // Link Title
}
