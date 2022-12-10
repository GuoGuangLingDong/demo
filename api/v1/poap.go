package v1

import (
	"demo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type MyPoapReq struct {
	g.Meta `path:"/poap/my_list" method:"get" tags:"PoapService" summary:"Get the poap of me"`
}
type MyPoapRes struct {
	Res []*PoapDetailPoapRes
}

type MainPagePoapReq struct {
	g.Meta    `path:"/poap/mainpage_list" method:"post" tags:"PoapService" summary:"Get the poap of main page"`
	From      int64
	Count     int64
	Condition string
}

type MainPagePoapRes struct {
	Res []*PoapDetailPoapRes `json:"list,omitempty"`
}

type PoapDetailReq struct {
	g.Meta `path:"/poap/details" method:"post" tags:"PoapService" summary:"Get the detail of poap"`
	PoapId string `p:poap_id`
}

type Chain struct {
	PlatForm     string `json:"plat_form"`
	PublishTime  string `json:"publish_time"`
	ContractNo   string `json:"contract_no"`
	ContractAddr string `json:"contract_addr"`
}
type PoapDetailPoapRes struct {
	*entity.Poap
	LikeNum      int         `json:"favour_number,omitempty"`
	FollowMiner  int         `json:"follow_miner"`
	HolderNumber int         `json:"holder_num"`
	Favoured     bool        `json:"favoured"`
	Holders      []*UserInfo `json:"holders,omitempty"`
	Collectable  bool        `json:"collectable"`
	Chain        *Chain      `json:"chain"`
	Avatar       string      `json:"avatar"`
	*Miner
}

type Miner struct {
	MinerUid  string `json:"minerUid"`
	MinerName string `json:"minerName"`
	MinerIcon string `json:"minerIcon"`
}

type PoapCollectReq struct {
	g.Meta `path:"/poap/collect" method:"post" tags:"PoapService" summary:"Collect a poap"`
	PoapId string `json:"poap_id" v:"required"`
}

type PoapCollectRes struct{}

type PoapMintReq struct {
	g.Meta      `path:"/poap/mint" method:"post" tags:"PoapService" summary:"Mint a poap"`
	PoapName    string `json:"poap_name" v:"required|length:2,30"`
	PoapSum     int64  `json:"poap_sum" v:"required|integer|between:1,10000"`
	ReceiveCond int64  `json:"receive_cond" v:"required|integer"`
	CoverImg    string `json:"cover_img" v:"required"`
	PoapIntro   string `json:"poap_intro" v:"required"`
	MintPlat    int    `json:"mint_plat" v:"required"`
	CollectList string `json:"collect_list" v:"required-if:receive_cond,2"`
}

type PoapMintRes struct{}

// ChainCallbackReq 上链回调
type ChainCallbackReq struct {
	g.Meta    `path:"/poap/chain/callback" method:"post" tags:"PoapService" summary:"Poap chain callback"`
	Code      int    `json:"code"`
	Contract  string `json:"contract" v:"required"`
	Hash      string `json:"hash" v:"required"`
	OperateId string `json:"operateId" v:"required"`
	Status    string `json:"status"`
	TokenId   string `json:"tokenId" v:"required"`
	Type      string `json:"type" v:"required"`
	Msg       string `json:"msg"`
}

type ChainCallbackRes struct{}

type FavorReq struct {
	g.Meta `path:"/poap/favor" method:"post" tags:"PoapService" summary:"Favor Poap"`
	PoapId string `json:"poap_id" v:"required"`
}

type FavorRes struct {
}

type GetHoldersReq struct {
	g.Meta `path:"/poap/holders" method:"get" tags:"PoapService" summary:"Get the holders of poap"`
	PoapId string `json:"poap_id"`
	From   int    `json:"from"`
	Count  int    `json:"count"`
}
type HolderInfo struct {
	*UserInfo
	Follow int `json:"follow"`
}
type GetHodlersRes struct {
	Res []*HolderInfo
}
