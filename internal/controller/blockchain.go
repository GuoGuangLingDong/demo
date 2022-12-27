package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/dao"
	"demo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

var BlockChain = cBlockChain{}

type cBlockChain struct{}

func (c *cBlockChain) GetMetaData(ctx context.Context, req *v1.BlockChainDataReq) (res *v1.BlockChainDataRes, err error) {
	var poapInfo entity.Poap
	g.Model(dao.Poap.Table()).Where("poap_id", req.PoapId).Scan(&poapInfo)
	if poapInfo.PoapId == "" {
		return
	}
	res = &v1.BlockChainDataRes{
		Name:        poapInfo.PoapName,
		Description: poapInfo.PoapIntro,
		Image:       poapInfo.CoverImg,
	}
	return
}
