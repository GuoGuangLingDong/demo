// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Like is the golang structure for table like.
type Like struct {
	Id       uint        `json:"id"       ` // pk
	Uid      string      `json:"uid"      ` // User ID
	PoapId   string      `json:"poapId"   ` // Poap id
	CreateAt *gtime.Time `json:"createAt" ` // Created Time
	UpdateAt *gtime.Time `json:"updateAt" ` // Updated Time
}
