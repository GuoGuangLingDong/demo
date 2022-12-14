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
	"strings"
	"sync"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type (
	SPoap struct{}
)

func (S SPoap) Favor(ctx context.Context, in *v1.FavorReq) (err error) {
	uid := service.Session().GetUser(ctx).Uid
	count, err := dao.Like.Ctx(ctx).Where("uid = ? and poap_id = ?", uid, in.PoapId).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		_, err := dao.Like.Ctx(ctx).Data(do.Like{
			Uid:    uid,
			PoapId: in.PoapId,
		}).Insert()
		if err != nil {
			return err
		}
		//更新缓存
		key := fmt.Sprintf("poapid-%s-uid-%s", in.PoapId, uid)
		cmd, err := g.Redis().Do(ctx, "EXISTS", key)
		if err != nil {
			err = fmt.Errorf("查询缓存失败")
		}
		exists := cmd.Int64()
		if exists == 1 {
			//在内存中，使缓存失效
			_, err := g.Redis().Do(ctx, "DEL", key)
			if err != nil {
				err = fmt.Errorf("删除缓存失败")
			}

		}
	}
	return nil
}

type PoapIdLike struct {
	poapId string
	number int64
}

var lock sync.Mutex
var wg sync.WaitGroup

func (S SPoap) GetMyPoap(ctx context.Context, in model.GetMyPoapInput) []*v1.PoapDetailPoapRes {
	//TODO implement me
	uid := in.UId

	holds, err := dao.Hold.Ctx(ctx).Where(do.Hold{Uid: uid}).Limit(in.From, in.Count).All()
	if err != nil {

		return []*v1.PoapDetailPoapRes{}
	}
	poap_ids := []string{}
	for _, hold := range holds {
		poap_id, _ := hold.Map()["poap_id"]
		poap_ids = append(poap_ids, poap_id.(string))
	}
	res := []*v1.PoapDetailPoapRes{}
	res = S.GetPoapsDetail(ctx, model.GetPoapsDetailsInput{PoapIds: poap_ids, Uid: in.UId})
	return res
}

func (S SPoap) GetMainPagePoap(ctx context.Context, in model.GetMainPagePoap) []*v1.PoapDetailPoapRes {
	res := ([]*v1.PoapDetailPoapRes)(nil)
	all, err := dao.Poap.Ctx(ctx).LeftJoin("`like` l", "poap.poap_id=l.poap_id").Fields("poap.poap_id").Where("poap_name like ?", "%"+in.Condition+"%").Group("poap_id").Order("count(`uid`) desc").Limit((int)(in.From), int(in.Count)).All()
	if err != nil {
		return nil
	}
	poap_ids := []string{}
	for _, like := range all {
		poap_id, _ := like.Map()["poap_id"]
		poap_ids = append(poap_ids, poap_id.(string))
	}
	res = S.GetPoapsDetail(ctx, model.GetPoapsDetailsInput{PoapIds: poap_ids})
	return res
}

// todo 点赞、持有等关系变化需要设置缓存失效，或者缓存粒度设粗
func (S SPoap) GetPoapsDetail(ctx context.Context, in model.GetPoapsDetailsInput) []*v1.PoapDetailPoapRes {
	res := []*v1.PoapDetailPoapRes{}

	var uid string
	if in.Uid != "" && in.Uid != "tempUser" {
		uid = in.Uid
	} else if service.Session().GetUser(ctx) != nil {
		uid = service.Session().GetUser(ctx).Uid
	} else {
		uid = "tempUser"
	}
	fmt.Println("FinalUid: ", uid)
	for _, poapId := range in.PoapIds {
		key := fmt.Sprintf("poapid-%s-uid-%s", poapId, uid)
		cmd, err := g.Redis().Do(ctx, "EXISTS", key)
		if err != nil {
			return res
		}
		exists := cmd.Int64()
		if exists == 1 {
			//在内存中从内存中取
			// fmt.Println("查redis，key：", key)
			gv, err := g.Redis().Do(ctx, "GET", key)
			if err != nil {
				return res
			}
			var curPoap *v1.PoapDetailPoapRes
			err = gv.Scan(&curPoap)
			if err != nil {
				return res
			} else {
				res = append(res, curPoap)
				_, err = g.Redis().Do(ctx, "SET", key, curPoap, "ex", 3*24*3600)
				if err != nil {
					// g.Log().Errorf("发送验证码失败：%v", err)
					err = fmt.Errorf("设置缓存失败")
				}
			}

		} else {
			//不在内存从数据库里查
			fmt.Println("查数据库")
			curPoap := S.GetPoapDetail(ctx, poapId, uid)
			res = append(res, curPoap)
			_, err = g.Redis().Do(ctx, "SET", key, curPoap, "ex", 3*24*3600)
			if err != nil {
				// g.Log().Errorf("发送验证码失败：%v", err)
				err = fmt.Errorf("设置缓存失败")
			}
		}
	}

	return res

}

func (S SPoap) GetPoapDetail(ctx context.Context, poapId, uid string) *v1.PoapDetailPoapRes {
	res := &v1.PoapDetailPoapRes{}
	dao.Poap.Ctx(ctx).Where("poap_id", poapId).Scan(&res.Poap)
	res.LikeNum, _ = dao.Like.Ctx(ctx).Where("poap_id", poapId).Count()
	res.HolderNumber, _ = dao.Hold.Ctx(ctx).Where("poap_id", poapId).Count()
	avatar, _ := dao.User.Ctx(ctx).Fields("avatar").Where("uid", uid).Value()
	res.Avatar = avatar.String()
	chainConf := getChainConf()
	res.Chain = &v1.Chain{
		PlatForm:     chainConf.Name,
		PublishTime:  res.Poap.CreateAt.Format("Y-m-d H:i:s"),
		ContractNo:   res.Poap.PoapId,
		ContractAddr: chainConf.ChainAddr,
	}
	var miner *entity.User
	dao.User.Ctx(ctx).Where("uid", res.Poap.Miner).Scan(&miner)
	res.Miner = &v1.Miner{
		MinerUid:  miner.Uid,
		MinerName: miner.Username,
		MinerIcon: miner.Avatar,
	}

	var favour int
	var follow int
	if uid == "tempUser" {
		res.Collectable = false
		favour = 0
		follow = 0
	} else {
		res.Collectable = S.isCollectable(ctx, poapId, uid)
		favour, _ = dao.Like.Ctx(ctx).Where("poap_id", poapId).Where("uid", uid).Count()
		follow, _ = dao.Follow.Ctx(ctx).Where("followee", miner.Uid).Where("follower", uid).Count()
	}
	fmt.Println("follow：", follow)
	if favour == 0 {
		res.Favoured = false
	} else {
		res.Favoured = true
	}
	if follow == 0 {
		res.FollowMiner = 0
	} else {
		res.FollowMiner = 1
	}
	return res
}

func (S SPoap) isCollectable(ctx context.Context, poapId, uid string) bool {
	if uid == "" {
		uid = service.Session().GetUser(ctx).Uid
	}
	poap := &entity.Poap{}
	dao.Poap.Ctx(ctx).Where("poap_id", poapId).Scan(&poap)
	receiveCond := poap.ReceiveCond
	if receiveCond == 1 { //所有人可领取 未持有 有剩余
		holdNum, _ := dao.Hold.Ctx(ctx).Where("uid", uid).Where("poap_id", poapId).Count()
		if holdNum != 0 {
			return false
		}
		poapSum, _ := dao.Poap.Ctx(ctx).Fields("poap_sum").Where("poap_id", poapId).Value()
		poapHold, _ := dao.Hold.Ctx(ctx).Where("poap_id", poapId).Count()
		if poapSum.Int() <= poapHold {
			return false
		}
		return true
	} else if receiveCond == 2 { //指定人可领取 未持有 有剩余
		holdNum, _ := dao.Hold.Ctx(ctx).Where("uid", uid).Where("poap_id", poapId).Count()
		if holdNum != 0 {
			return false
		}
		collectList, _ := dao.Poap.Ctx(ctx).Fields("collect_list").Where("poap_id", poapId).Value()
		// fmt.Println("phonenumber：", service.Session().GetUser(ctx).PhoneNumber)
		if strings.Contains(collectList.String(), service.Session().GetUser(ctx).Did) {
			return true
		} else if strings.Contains(collectList.String(), service.Session().GetUser(ctx).PhoneNumber) {
			return true
		}
		return false
	} else if receiveCond == 3 { //凭口令领取

	} else if receiveCond == 4 { //我的链接
		follow, _ := dao.Follow.Ctx(ctx).Where("follower", poap.Miner).Where("followee", uid).Count()
		if follow == 0 {
			return false
		}
		return true
	} else if receiveCond == 5 { //链接我的
		followee, _ := dao.Follow.Ctx(ctx).Where("follower", uid).Where("followee", poap.Miner).Count()
		if followee == 0 {
			return false
		}
		return true

	} else if receiveCond == 6 { //付费领取

	}
	return false
}

func (S SPoap) getPoapUser(ctx context.Context, poapId string, from int, count int) []*v1.UserInfo {
	holderRes := gdb.Result{}
	if from == -1 {
		holderRes, _ = dao.Hold.Ctx(ctx).Fields("DISTINCT uid").Where("poap_id", poapId).All()
	} else {
		holderRes, _ = dao.Hold.Ctx(ctx).Fields("DISTINCT uid").Limit(from, count).Where("poap_id", poapId).All()
	}

	holders := ([]*v1.UserInfo)(nil)
	holderIds := ([]string)(nil)
	for _, holder := range holderRes {
		holderId, _ := holder.Map()["uid"]
		holderIds = append(holderIds, holderId.(string))
	}
	dao.User.Ctx(ctx).Where("uid in (?)", holderIds).Scan(&holders)
	return holders
}
func (S SPoap) GetHolders(ctx context.Context, in *v1.GetHoldersReq) []*v1.HolderInfo {
	userId := service.Session().GetUser(ctx).Uid
	holdersInit := S.getPoapUser(ctx, in.PoapId, in.From, in.Count)
	holders := ([]*v1.HolderInfo)(nil)
	// miner, _ := dao.Poap.Ctx(ctx).Fields("miner").Where("poap_id", in.PoapId).Value()
	for _, user := range holdersInit {
		temp := &v1.HolderInfo{}
		temp.UserInfo = user
		follow, _ := dao.Follow.Ctx(ctx).Where("followee", user.Uid).Where("follower", userId).Count()
		if follow == 0 {
			temp.Follow = 0
		} else {
			temp.Follow = 1
		}
		holders = append(holders, temp)
	}
	return holders

}
func (S SPoap) CollectPoap(ctx context.Context, in model.CollectPoapInput) (err error) {
	userId := service.Session().GetUser(ctx).Uid
	err = S.publishPoap(ctx, userId, in.PoapId, 1)
	//更新缓存
	key := fmt.Sprintf("poapid-%s-uid-%s", in.PoapId, userId)
	cmd, err := g.Redis().Do(ctx, "EXISTS", key)
	if err != nil {
		err = fmt.Errorf("查询缓存失败")
	}
	exists := cmd.Int64()
	if exists == 1 {
		//在内存中，使缓存失效
		_, err := g.Redis().Do(ctx, "DEL", key)
		if err != nil {
			err = fmt.Errorf("删除缓存失败")
		}

	}
	return err
}

func (S SPoap) MintPoap(ctx context.Context, in model.MintPoapInput) (err error) {
	user := service.Session().GetUser(ctx)
	if user == nil {
		err = gerror.New("获取用户信息失败")
		return
	}
	poapId := S.generatePoapId(ctx)
	newPoap := &entity.Poap{
		PoapId: poapId,
		Miner:  user.Uid,
		//Miner:       "aea7cbc430cb4a7893896e64a5dc2b9c",
		PoapName:    in.PoapName,
		PoapSum:     int(in.PoapSum),
		ReceiveCond: int(in.ReceiveCond),
		CoverImg:    in.CoverImg,
		PoapIntro:   in.PoapIntro,
		MintPlat:    in.MintPlat,
		CollectList: in.CollectList,
	}
	_, err = dao.Poap.Ctx(ctx).Insert(newPoap)
	if err != nil {
		return
	}

	_ = S.Generate(ctx, model.GenerateTokenReq{
		PoapId: poapId,
		Num:    uint(in.PoapSum),
	})
	return
}
func (S SPoap) generatePoapId(ctx context.Context) string {
	uid := uuid.NewString()
	uid = strings.Replace(uid, "-", "", -1)
	return uid
}

// publishPoap 发放POAP
func (S SPoap) publishPoap(ctx context.Context, userId string, poapId string, num int) (err error) {
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
func (S SPoap) Generate(ctx context.Context, req model.GenerateTokenReq) (err error) {
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
