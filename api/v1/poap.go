package v1

import (
	"demo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type MyPoapReq struct {
	g.Meta `path:"/poap/my_list" method:"get" tags:"PoapService" summary:"Get the poap of me"`
}
type MyPoapRes struct {
	Res []*entity.Poap
}

type MainPagePoapReq struct {
	g.Meta `path:"/poap/mainpage_list" method:"post" tags:"PoapService" summary:"Get the poap of main page"`
	From   int64
	Count  int64
}

type MainPagePoapRes struct {
	Res []*PoapDetailPoapRes
}

type PoapDetailReq struct {
	g.Meta `path:"/poap/details" method:"post" tags:"PoapService" summary:"Get the detail of poap"`
	PoapId int64 `p:poapid`
}

type UserInfo struct {
	Uid      string `json:"uid,omitempty"`
	Username string `json:"username,omitempty"`
}
type PoapDetailPoapRes struct {
	*entity.Poap `json:"poap,omitempty"`
	LikeNum      int         `json:"like_num,omitempty"`
	Holders      []*UserInfo `json:"holders,omitempty"`
}
