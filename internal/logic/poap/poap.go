package poap

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"time"
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
	userId := service.Session().GetUser(ctx).Uid
	err = S.publishPoap(ctx, userId, in.PoapId, 1)
	return err
}

func (S SPoap) MintPoap(ctx context.Context, in model.MintPoapInput) (err error) {
	poapId := S.generatePoapId(ctx)
	newPoap := &entity.Poap{
		PoapId: poapId,
		Miner:  service.Session().GetUser(ctx).Uid,
		//Miner:       1,
		PoapName:    in.PoapName,
		PoapSum:     int(in.PoapSum),
		ReceiveCond: int(in.ReceiveCond),
		CoverImg:    in.CoverImg,
		PoapIntro:   in.PoapIntro,
	}
	_, err = dao.Poap.Ctx(ctx).Insert(newPoap)
	if err != nil {
		return
	}

	_ = S.generate(ctx, model.GenerateTokenReq{
		PoapId: poapId,
		Num:    uint(in.PoapSum),
	})
	return
}
func (S SPoap) generatePoapId(ctx context.Context) uint {
	return uint(time.Now().Unix())
}

// publishPoap 发放POAP
func (S SPoap) publishPoap(ctx context.Context, userId uint, poapId int64, num int) (err error) {
	var asset []entity.Publish
	m := g.DB().Model("publish")
	m.Where("poap_id", poapId)
	m.Where("status", "disable")
	m.Where("lock_flag", 1)
	m.Order("no ASC")
	m.Limit(num)
	err = m.Scan(&asset)
	if len(asset) == 0 {
		g.Log().Errorf(ctx, "未查询可用资产：poapId:%s", poapId)
		err = fmt.Errorf("出错了，请稍后再试")
		return
	}
	if len(asset) != num {
		g.Log().Errorf(ctx, "查询可用资产数量不足：poapId:%s", poapId)
		err = fmt.Errorf("出错了，请稍后再试")
		return
	}

	hold := g.Slice{}
	ids := make([]int64, 0)
	for _, v := range asset {
		ids = append(ids, v.Id)
		hold = append(hold, g.Map{
			"uid":      userId,
			"poap_id":  poapId,
			"token_id": v.TokenId,
		})
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		result, err := tx.Model(dao.Publish.Table()).Data(g.Map{
			"lock_flag": gdb.Raw("lock_flag - 1"),
			"status":    "used",
		}).WhereIn("id", ids).Update()
		if err != nil {
			return err
		}
		ra, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if ra == 0 {
			err = fmt.Errorf("RowsAffected = 0")
			return err
		}
		_, err = tx.Model(dao.Hold.Table()).Insert(hold)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		g.Log().Errorf(ctx, "领取poap失败：poapId:%s, err:%v", poapId, err.Error())
		err = fmt.Errorf("出错了，请稍后再试")
		return
	}
	return
}

// generate poap铸造发行
func (S SPoap) generate(ctx context.Context, req model.GenerateTokenReq) (err error) {
	// 生成token
	tokens, err := generateToken(ctx, req)
	if err != nil {
		return
	}
	lockFlag := uint(1)
	unlockFlag := uint(0)
	var publish []*entity.Publish
	for _, v := range tokens {
		if v.ErrorMessage != "" {
			publish = append(publish, &entity.Publish{
				PoapId:       req.PoapId,
				TokenId:      v.TokenId,
				No:           v.No,
				ErrorMessage: v.ErrorMessage,
				LockFlag:     unlockFlag,
				Status:       consts.PublishStatusError,
			})
		} else {
			publish = append(publish, &entity.Publish{
				PoapId:   req.PoapId,
				TokenId:  v.TokenId,
				No:       v.No,
				LockFlag: lockFlag,
				Status:   consts.PublishStatusDisable,
			})
		}
	}
	// insert publish
	_, err = dao.Publish.Ctx(ctx).Batch(5000).Data(publish).Insert()
	if err != nil {
		return
	}

	// 上链
	_ = upChain(req.PoapId)
	return
}

func init() {
	service.RegisterPoap(New())
}

func New() *SPoap {
	return &SPoap{}
}
