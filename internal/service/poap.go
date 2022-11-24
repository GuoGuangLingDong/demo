package service

import (
	"context"
	"demo/internal/model"
	"demo/internal/model/entity"
)

type IPoap interface {
	GetMyPoap(ctx context.Context, in model.GetMyPoapInput) []*entity.Poap
	GetMainPagePoap(ctx context.Context) []*entity.Poap
	GetPoapDetails(ctx context.Context, in model.GetPoapDetailsInput) *entity.Poap
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
