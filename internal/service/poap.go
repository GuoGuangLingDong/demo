package service

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/model"
)

type IPoap interface {
	GetMyPoap(ctx context.Context, in model.GetMyPoapInput) []*v1.PoapDetailPoapRes
	GetMainPagePoap(ctx context.Context, in model.GetMainPagePoap) []*v1.PoapDetailPoapRes
	GetPoapsDetail(ctx context.Context, in model.GetPoapsDetailsInput) []*v1.PoapDetailPoapRes
	CollectPoap(ctx context.Context, in model.CollectPoapInput) (err error)
	MintPoap(ctx context.Context, in model.MintPoapInput) (poapId string, err error)
	ChainCallback(ctx context.Context, in *v1.ChainCallbackReq) (err error)
	Generate(ctx context.Context, in model.GenerateTokenReq) (err error)
	Favor(ctx context.Context, in *v1.FavorReq) (err error)
	GetHolders(ctx context.Context, in *v1.GetHoldersReq) []*v1.HolderInfo
}

var (
	poap IPoap
)

func Poap() IPoap {
	if poap == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return poap
}

func RegisterPoap(i IPoap) {
	poap = i
}
