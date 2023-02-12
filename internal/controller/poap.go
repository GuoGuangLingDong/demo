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
	res.Res = service.Poap().GetMyPoap(ctx, model.GetMyPoapInput{UId: user.Uid})
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
		From:      req.From,
		Count:     req.Count,
		Condition: req.Condition,
	})
	return
}

// Get poap info
func (c *cPoap) GetPoapDetail(ctx context.Context, req *v1.PoapDetailReq) (res *v1.PoapDetailPoapRes, err error) {
	res = service.Poap().GetPoapsDetail(ctx, model.GetPoapsDetailsInput{PoapIds: []string{req.PoapId}})[0]
	return
}

// CollectPoap poap领取
func (c *cPoap) CollectPoap(ctx context.Context, req *v1.PoapCollectReq) (res *v1.PoapCollectRes, err error) {
	res = &v1.PoapCollectRes{}
	err = service.Poap().CollectPoap(ctx, model.CollectPoapInput{PoapId: req.PoapId, Endorse: req.Endorse, EndorsePic: req.EndorsePic})
	return
}

// MintPoap poap铸造
func (c *cPoap) MintPoap(ctx context.Context, req *v1.PoapMintReq) (res *v1.PoapMintRes, err error) {
	res = &v1.PoapMintRes{}
	_, err = service.Poap().MintPoap(ctx, model.MintPoapInput{
		PoapName:    req.PoapName,
		PoapSum:     req.PoapSum,
		ReceiveCond: req.ReceiveCond,
		CoverImg:    req.CoverImg,
		PoapIntro:   req.PoapIntro,
		MintPlat:    req.MintPlat,
		CollectList: req.CollectList,
		Type:        1,
		SeriesId:    req.SeriesId,
	})
	return
}

// ChainCallback 上链回调
func (c *cPoap) ChainCallback(ctx context.Context, req *v1.ChainCallbackReq) (res *v1.ChainCallbackRes, err error) {
	res = &v1.ChainCallbackRes{}
	err = service.Poap().ChainCallback(ctx, req)
	return
}

func (c *cPoap) Favor(ctx context.Context, req *v1.FavorReq) (res *v1.FavorRes, err error) {
	res = &v1.FavorRes{}
	err = service.Poap().Favor(ctx, req)
	return
}

func (c *cPoap) GetHolders(ctx context.Context, req *v1.GetHoldersReq) (res *v1.GetHodlersRes, err error) {
	res = &v1.GetHodlersRes{}
	res.Res = service.Poap().GetHolders(ctx, req)
	return
}

func (c *cPoap) CreatePoapSeries(ctx context.Context, req *v1.CreatePoapSeriesReq) (res *v1.CreatePoapSeriesRes, err error) {
	res = &v1.CreatePoapSeriesRes{}
	user := service.Session().GetUser(ctx)
	err = service.Poap().CreatePoapSeries(ctx, req, user.Uid)
	return
}

func (c *cPoap) GetPoapSeries(ctx context.Context, req *v1.GetPoapSeriesReq) (res *v1.GetPoapSeriesRes, err error) {
	res = &v1.GetPoapSeriesRes{}
	user := service.Session().GetUser(ctx)
	res.Res = service.Poap().GetPoapSeries(ctx, req, user.Uid)
	return
}

func (c *cPoap) GetPoapSeriesDetail(ctx context.Context, req *v1.GetPoapSeriesDetailReq) (res *v1.GetPoapSeriesDetailRes, err error) {
	res = &v1.GetPoapSeriesDetailRes{}
	res.SeriesDeatil = service.Poap().GetPoapSeriesDetail(ctx, req)
	return
}

func (c *cPoap) GetEndorse(ctx context.Context, req *v1.GetEndorseReq) (res *v1.GetEndorseRes, err error) {
	res = &v1.GetEndorseRes{}
	user := service.Session().GetUser(ctx)
	res.Res = service.Poap().GetEndorse(ctx, req, user.Uid)
	return
}

func (c *cPoap) LikeEndorse(ctx context.Context, req *v1.LikeEndorseReq) (res *v1.LikeEndorseRes, err error) {
	user := service.Session().GetUser(ctx)
	res = &v1.LikeEndorseRes{}
	err = service.Poap().LikeEndorse(ctx, req, user.Uid)
	return
}
