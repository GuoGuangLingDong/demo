// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta       `orm:"table:user, do:true"`
	Id           interface{} // pk
	Uid          interface{} // User ID
	Password     interface{} // User Password
	Username     interface{} // User Name
	Nickname     interface{} // User Nickname
	CreateAt     *gtime.Time // Created Time
	UpdateAt     *gtime.Time // Updated Time
	PhoneNumber  interface{} // Phone Number
	WechatNumber interface{} // Wechat Number
	InviteCode   interface{} // Invite Code
	Introduction interface{} // Introduction
	Avatar       interface{} // 头像
}
