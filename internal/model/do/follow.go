// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Follow is the golang structure of table follow for DAO operations like Where/Data.
type Follow struct {
	g.Meta     `orm:"table:follow, do:true"`
	Followee   interface{} // Followee ID
	Follower   interface{} // Follower id
	FollowTime *gtime.Time // Follow Time
}
