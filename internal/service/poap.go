package service

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/model"
	"demo/internal/model/entity"
)

type IPoap interface {
	GetMyPoap(ctx context.Context, in model.GetMyPoapInput) []*entity.Poap
	GetMainPagePoap(ctx context.Context, in model.GetMainPagePoap) []*v1.PoapDetailPoapRes
	GetPoapDetails(ctx context.Context, in model.GetPoapDetailsInput) *v1.PoapDetailPoapRes
	CollectPoap(ctx context.Context, in model.CollectPoapInput) (err error)
	MintPoap(ctx context.Context, in model.MintPoapInput) (err error)
	ChainCallback(ctx context.Context, in *v1.ChainCallbackReq) (err error)
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
