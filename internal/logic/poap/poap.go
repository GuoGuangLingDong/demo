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

type PoapIdLike struct {
	poapId int
	number int64
}

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

func (S SPoap) GetMainPagePoap(ctx context.Context, in model.GetMainPagePoap) []*v1.PoapDetailPoapRes {
	likes := ([]PoapIdLike)(nil)
	res := ([]*v1.PoapDetailPoapRes)(nil)
	poapRes := ([]*entity.Poap)(nil)
	all, err := dao.Like.Ctx(ctx).Fields("poapId, count(`uid`) total").Group("poap_id").Order("count(`uid`) desc").Limit((int)(in.From), int(in.Count)).All()
	if err != nil {
		return nil
	}
	poapIds := []int64{}
	for _, like := range all {
		poap_id, _ := like.Map()["poap_id"]
		count, _ := like.Map()["total"]
		likes = append(likes, PoapIdLike{
			poapId: poap_id.(int),
			number: count.(int64),
		})
		poapIds = append(poapIds, int64(poap_id.(int)))
	}

	dao.Poap.Ctx(ctx).Where("poap_id in(?)", poapIds).Scan(&poapRes)

	for _, like := range likes {
		for _, poap := range poapRes {
			if uint(like.poapId) == poap.PoapId {
				uid, uname := S.getPoapUser(ctx, int64(poap.PoapId))
				res = append(res, &v1.PoapDetailPoapRes{
					poap,
					int64(like.number),
					uid,
					uname,
				})
			}
		}
	}
	return res
}

func (S SPoap) GetPoapDetails(ctx context.Context, in model.GetPoapDetailsInput) *v1.PoapDetailPoapRes {
	//TODO implement me
	//poapId := in.PoapId
	//res := (*entity.Poap)(nil)
	//dao.Poap.Ctx(ctx).Where("poap_id", poapId).Scan(&res)
	return nil
}

func (S SPoap) getPoapUser(ctx context.Context, poapId int64) (int64, string) {
	holdRes := (*entity.Hold)(nil)
	dao.Hold.Ctx(ctx).Where("poap_id", poapId).Scan(&holdRes)
	userRes := (*entity.User)(nil)
	dao.User.Ctx(ctx).Where("uid", holdRes.Uid).Scan(&userRes)
	return int64(userRes.Uid), userRes.Username
}

func init() {
	service.RegisterPoap(New())
}

func New() *SPoap {
	return &SPoap{}
}
