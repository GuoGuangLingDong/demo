package poap

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

type (
	SPoap struct{}
)

func (S SPoap) GetMyPoap(ctx context.Context, in model.GetMyPoapInput) []*entity.Poap {
	//TODO implement me
	uid := in.UId

	holds, err := dao.Hold.Ctx(ctx).Where(do.Hold{Uid: uid}).All()
	if err != nil {
		return []*entity.Poap{}
	}
	poap_ids := []int64{}
	for _, hold := range holds {
		poap_id, _ := hold.Map()["poap_id"]
		poap_ids = append(poap_ids, int64(poap_id.(int)))
	}
	res := ([]*entity.Poap)(nil)
	dao.Poap.Ctx(ctx).Where("poap_id in(?)", poap_ids).Scan(&res)
	return res
}

func (S SPoap) GetMainPagePoap(ctx context.Context, in model.GetMainPagePoap) []*entity.Poap {
	res := ([]*entity.Poap)(nil)
	dao.Poap.Ctx(ctx).Order("favour_number desc").Limit((int)(in.From), int(in.Count)).Scan(&res)
	return res
}

func (S SPoap) GetPoapDetails(ctx context.Context, in model.GetPoapDetailsInput) *entity.Poap {
	//TODO implement me
	panic("implement me")
}

func init() {
	service.RegisterPoap(New())
}

func New() *SPoap {
	return &SPoap{}
}
