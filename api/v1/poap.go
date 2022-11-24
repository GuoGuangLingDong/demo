package v1

import (
	"demo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type MyPoapReq struct {
	g.Meta `path:"/poap/my_list" method:"get" tags:"PoapService" summary:"Get the poap of me"`
}
type MyPoapRes struct {
	res []*entity.Poap
}

type MainPagePoapReq struct {
	g.Meta `path:"/poap/mainpage_list" method:"get" tags:"PoapService" summary:"Get the poap of main page"`
}

type MainPagePoapRes struct {
	res []*entity.Poap
}

type PoapDetailReq struct {
	g.Meta `path:"/poap/details" method:"get" tags:"PoapService" summary:"Get the detail of poap"`
}

type PoapDetailPoapRes struct {
	*entity.Poap
}
