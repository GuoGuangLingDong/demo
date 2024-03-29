// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Poap is the golang structure for table poap.
type Poap struct {
	Id          uint        `json:"id"          ` // pk
	PoapId      string      `json:"poapId"      ` // Poap id
	Miner       string      `json:"miner"       ` // Miner
	PoapName    string      `json:"poapName"    ` // Poap name
	PoapSum     int         `json:"poapSum"     ` // Poap sum
	ReceiveCond int         `json:"receiveCond" ` // Receive condition
	CoverImg    string      `json:"coverImg"    ` // Cover picture
	PoapIntro   string      `json:"poapIntro"   ` // Poap introduction
	CreateAt    *gtime.Time `json:"createAt"    ` // Created Time
	UpdateAt    *gtime.Time `json:"updateAt"    ` // Updated Time
	CollectList string      `json:"collectList" ` //
	MintPlat    int         `json:"mintPlat"    ` // Mint platform
	Status      int         `json:"status"      ` // 状态 0.新建 1.正常（已铸造） 2.审核通过 3.审核不通过
	Type        int         `json:"type"        ` // 类型  1.poap 2.头像nft 3.did
	Seriesid    string      `json:"seriesid"    ` //
}
