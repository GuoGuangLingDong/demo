// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Follow is the golang structure for table follow.
type Follow struct {
	Followee   string      `json:"followee"   ` // Followee ID
	Follower   string      `json:"follower"   ` // Follower id
	FollowTime *gtime.Time `json:"followTime" ` // Follow Time
}
