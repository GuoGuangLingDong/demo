package poap

import (
	"context"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

// generateToken 生成token
func generateToken(ctx context.Context, req model.GenerateTokenReq) (gens []*model.GenerateTokenRes, err error) {
	tokenIdType := model.TokenId("")
	tokens := make(chan uint64)
	errChan := make(chan error)
	no := make(chan uint64)

	err = generateMultiRedisTokenInt(ctx, tokens, errChan, req.Num)

	if err != nil {
		return
	}

	err = generateMultiNoInt(req.PoapId, no, req.Num)

	if err != nil {
		return
	}

	var i uint

	for {
		i++

		noData := <-no

		select {
		case err := <-errChan:
			gens = append(gens, &model.GenerateTokenRes{
				No:           noData,
				ErrorMessage: err.Error(),
			})
			g.Log().Error(ctx, "SelfGenerate:req:%+v-==-err:%+v", req, err)
		case token := <-tokens:
			gens = append(gens, &model.GenerateTokenRes{
				No:      noData,
				TokenId: tokenIdType.Uint64ToTokenId(token).ToStr(),
			})
		}

		if i == req.Num {
			close(tokens)
			close(errChan)
			close(no)
			break
		}
	}

	return
}

// generateMultiRedisTokenInt 生成tokenId 通过redis
func generateMultiRedisTokenInt(ctx context.Context, tokens chan uint64, errChan chan error, num uint) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("redis链接失败")
		}
	}()

	cmd, err := g.Redis().Do(ctx, "EXISTS", consts.TokenKey)
	if err != nil {
		return
	}
	exists := cmd.Int64()

	for i := num; i > 0; i-- {
		if exists == 0 {
			init, err := getTplNextTokenIdV2()
			if err != nil {
				return err
			}
			nowTokenId, err := g.Redis().Do(ctx, "INCRBY", consts.TokenKey, int64(int(init+1)))

			go func(nowTokenId *gvar.Var, err error) {
				if err != nil {
					fmt.Println("err")
					errChan <- err
				} else {
					fmt.Println("suc")
					tokens <- nowTokenId.Uint64()
				}
			}(nowTokenId, err)

		} else {
			go func() {
				nowTokenId, err := g.Redis().Do(ctx, "INCR", consts.TokenKey)
				if err != nil {
					errChan <- err
				} else {
					tokens <- nowTokenId.Uint64()
				}
				return
			}()
		}
	}

	return nil
}

// generateMultiNoInt 生成 编号
func generateMultiNoInt(poapId uint, tokens chan uint64, num uint) error {
	// 获取起始值
	first, err := getTplNextNo(poapId)

	if err != nil {
		return err
	}

	go func() {
		i := uint(0)
		for {
			first++
			tokens <- first
			i++

			if i == num {
				break
			}
		}
	}()

	return nil
}

// GetTplNextNo 获取下一个编号
func getTplNextNo(poapId uint) (uint64, error) {
	val, err := g.DB().Model(dao.Publish.Table()).
		Where("poap_id = ?", poapId).
		Fields("MAX(no) max").
		Value()
	if err != nil {
		return 0, err
	}
	if val != nil && val.Int64() > 0 {
		return uint64(val.Int64()), nil
	}

	return 0, nil
}

// getTplNextTokenIdV2 获取下一个token_id
func getTplNextTokenIdV2() (uint64, error) {
	val, err := g.DB().Model(dao.Publish.Table()).
		Fields("MAX(token_id + 0) max").
		Where("token_id < ?", 121000000000).
		Value()
	if err != nil {
		return 0, err
	}
	if val != nil {
		return uint64(val.Float64()), nil
	}
	return 1000002000, nil
}
