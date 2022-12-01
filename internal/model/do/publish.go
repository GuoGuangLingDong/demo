// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Publish is the golang structure of table publish for DAO operations like Where/Data.
type Publish struct {
	g.Meta       `orm:"table:publish, do:true"`
	Id           interface{} // 主键ID
	PoapId       interface{} // Poap id
	TokenId      interface{} // Poap tokenId
	Status       interface{} // 状态 disable:未使用 used.已使用
	No           interface{} // 编号
	ChainStatus  interface{} // 上链状态  0.未上链 1.已上链  2.上链中 3.上链失败
	ChainHash    interface{} // 链上hash
	IsError      interface{} // 是否异常
	ErrorMessage interface{} // 错误信息
	LockFlag     interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}