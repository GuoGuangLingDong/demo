package poap

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
	"fmt"
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
				holders := S.getPoapUser(ctx, int64(poap.PoapId))
				collectable := S.isCollectable(ctx, int64(poap.PoapId))
				res = append(res, &v1.PoapDetailPoapRes{
					poap,
					int(like.number),
					holders,
					collectable,
				})
			}
		}
	}
	return res
}

func (S SPoap) GetPoapDetails(ctx context.Context, in model.GetPoapDetailsInput) *v1.PoapDetailPoapRes {
	poapId := in.PoapId
	res := &v1.PoapDetailPoapRes{}
	dao.Poap.Ctx(ctx).Where("poap_id", poapId).Scan(&res.Poap)
	likeNum, err := dao.Like.Ctx(ctx).Where("poap_id", poapId).Count()
	if err != nil {
		panic(err)
	}
	res.LikeNum = likeNum
	res.Holders = S.getPoapUser(ctx, poapId)
	res.Collectable = S.isCollectable(ctx, poapId)
	return res
}

func (S SPoap) isCollectable(ctx context.Context, poapId int64) bool {
	uid := service.Session().GetUser(ctx).Uid
	holdNum, _ := dao.Hold.Ctx(ctx).Where("uid", uid).Where("poap_id", poapId).Count()
	if holdNum != 0 {
		return false
	}
	poapSum, _ := dao.Poap.Ctx(ctx).Fields("poap_sum").Where("poap_id", poapId).Value()
	poapHold, _ := dao.Hold.Ctx(ctx).Where("poap_id", poapId).Count()
	fmt.Println("poapSum: ", poapSum)
	fmt.Println("poapHold: ", poapHold)

	if poapSum.Int() <= poapHold {
		return false
	}

	return true
}

func (S SPoap) getPoapUser(ctx context.Context, poapId int64) []*v1.UserInfo {
	holderRes, _ := dao.Hold.Ctx(ctx).Fields("DISTINCT uid").Where("poap_id", poapId).All()

	holders := ([]*v1.UserInfo)(nil)
	holderIds := ([]int64)(nil)
	for _, holder := range holderRes {
		holderId, _ := holder.Map()["uid"]
		holderIds = append(holderIds, int64(holderId.(int)))
	}
	dao.User.Ctx(ctx).Where("uid in (?)", holderIds).Scan(&holders)
	return holders
}

func (S SPoap) CollectPoap(ctx context.Context, in model.CollectPoapInput) (err error) {
	collect := &entity.Hold{
		Uid:    service.Session().GetUser(ctx).Uid,
		PoapId: uint(in.PoapId),
	}
	_, err = dao.Hold.Ctx(ctx).Insert(collect)
	return err
}

func (S SPoap) MintPoap(ctx context.Context, in model.MintPoapInput) (err error) {
	newPoap := &entity.Poap{
		PoapId:      S.generatePoapId(ctx),
		Miner:       service.Session().GetUser(ctx).Uid,
		PoapName:    in.PoapName,
		PoapSum:     int(in.PoapSum),
		ReceiveCond: int(in.ReceiveCond),
		CoverImg:    in.CoverImg,
		PoapIntro:   in.PoapIntro,
	}
	_, err = dao.Poap.Ctx(ctx).Insert(newPoap)

	return
}
func (S SPoap) generatePoapId(ctx context.Context) uint {
	return 0
}

func init() {
	service.RegisterPoap(New())
}

func New() *SPoap {
	return &SPoap{}
}
