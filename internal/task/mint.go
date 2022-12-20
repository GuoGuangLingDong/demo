package task

import (
	"demo/internal/dao"
	"demo/internal/logic/poap"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func Mint() {
	for {
		taskList := getMintList()
		if len(taskList) > 0 {
			for _, v := range taskList {
				// 中心化铸造
				_ = service.Poap().Generate(nil, model.GenerateTokenReq{
					PoapId: v.PoapId,
					Num:    uint(v.PoapSum),
				})
				// 头像NFT、DID铸造后领取
				if v.Type == 2 || v.Type == 3 {
					_ = service.Poap().PublishPoap(nil, v.Miner, v.PoapId, 1)
				}
				// 上链
				_ = poap.UpChain(v.PoapId)
			}
		}
		time.Sleep(time.Second * 3)
	}
}

func getMintList() (ret []entity.Poap) {
	g.DB().Model(dao.Poap.Table()).Where("status", 2).OrderAsc("id").Limit(5).Scan(&ret)
	return
}
