package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/model"
	"demo/internal/service"
	"fmt"
)

var Poap = cPoap{}

type cPoap struct{}

// SignUp is the API for get user poap list
func (c *cPoap) GetMyPoap(ctx context.Context, req *v1.MyPoapReq) (res *v1.MyPoapRes, err error) {
	user := service.Session().GetUser(ctx)
	res = &v1.MyPoapRes{Res: nil}
	res.Res = service.Poap().GetMyPoap(ctx, model.GetMyPoapInput{UId: int64(user.Uid)})
	fmt.Println(res.Res)
	return
}

// Get main page poaps
func (c *cPoap) GetMainPagePoap(ctx context.Context, req *v1.MainPagePoapReq) (res *v1.MainPagePoapRes, err error) {
	res = &v1.MainPagePoapRes{Res: nil}
	if req.Count == 0 {
		req.Count = 20
	}
	res.Res = service.Poap().GetMainPagePoap(ctx, model.GetMainPagePoap{
		From:  req.From,
		Count: req.Count,
	})
	return
}

// Get poap info
func (c *cPoap) GetPoapDetails(ctx context.Context, req *v1.PoapDetailReq) (res *v1.PoapDetailPoapRes, err error) {
	res = &v1.PoapDetailPoapRes{nil}
	res.Poap = service.Poap().GetPoapDetails(ctx, model.GetPoapDetailsInput{PoapId: req.PoapId})
	return
}
