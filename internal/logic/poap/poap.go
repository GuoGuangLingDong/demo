package poap

import (
	"context"
	v1 "demo/api/v1"
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
	likes := ([]int64)(nil)
	res := ([]*entity.Poap)(nil)
	dao.Like.Ctx(ctx).Fields("poapId").Group("poap_id").Order("count(`uid`) desc").Limit((int)(in.From), int(in.Count)).Scan(&res)
	dao.Poap.Ctx(ctx).Where("poap_id in(?)", likes).Scan(&res)
	return res
}

func (S SPoap) GetPoapDetails(ctx context.Context, in model.GetPoapDetailsInput) *v1.PoapDetailPoapRes {
	//TODO implement me
	poapId := in.PoapId
	res := &v1.PoapDetailPoapRes{}
	dao.Poap.Ctx(ctx).Where("poap_id", poapId).Scan(&res.Poap)
	likeNum, err := dao.Like.Ctx(ctx).Where("poap_id", poapId).Count()
	if err != nil {
		panic(err)
	}
	res.LikeNum = likeNum
	holderRes, _ := dao.Hold.Ctx(ctx).Fields("DISTINCT uid").Where("poap_id", poapId).All()

	holders := ([]*v1.UserInfo)(nil)
	holderIds := ([]int64)(nil)
	for _, holder := range holderRes {
		holderId, _ := holder.Map()["uid"]
		holderIds = append(holderIds, int64(holderId.(int)))
	}
	dao.User.Ctx(ctx).Where("uid in (?)", holderIds).Scan(&holders)
	res.Holders = holders
	return res
}

func init() {
	service.RegisterPoap(New())
}

func New() *SPoap {
	return &SPoap{}
}
