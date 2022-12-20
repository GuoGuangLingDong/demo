package poap

import (
	"bytes"
	"context"
	v1 "demo/api/v1"
	"demo/internal/dao"
	"demo/internal/model/entity"
	"demo/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"io/ioutil"
	"net/http"
	"time"
)

// ChainCallback  上链回调
func (S SPoap) ChainCallback(ctx context.Context, params *v1.ChainCallbackReq) (err error) {
	if params.Code == 0 && params.Status == "SUCCESS" && params.Type == "BATCH_TOKEN" {
		_, err = g.DB().Model(dao.Publish.Table()).Where("token_id = ?", params.TokenId).Data(map[string]interface{}{
			"chain_status": 1,
			"chain_hash":   params.Hash,
		}).Update()
		if err != nil {
			return
		}
	}
	return
}

// UpChain 上链
func UpChain(poapId string) (err error) {
	var poapInfo *entity.Poap
	dao.Poap.Ctx(nil).Where("poap_id = ?", poapId).Scan(&poapInfo)
	if poapInfo == nil {
		err = errors.New("未查找到poap")
		return
	}

	page := 1
	pageSize := 5000

	for {
		tokenIds, err := g.DB().Model(dao.Publish.Table()).Where("poap_id = ? and chain_status = 0", poapId).
			Fields("token_id").Limit(pageSize).Offset((page - 1) * pageSize).Array()
		page++
		if err != nil {
			g.Log().Error(nil, err)
			continue
		}
		if len(tokenIds) == 0 {
			break
		}

		// 构建上链参数
		conf := getChainConf()

		tokenInfo := make([]map[string]interface{}, 0)
		for _, v := range tokenIds {
			tokenInfo = append(tokenInfo, map[string]interface{}{
				"tokenId": v.String(),
				"url":     fmt.Sprintf(conf.TokenUrl, gconv.String(poapInfo.PoapId), v.String()),
			})
		}

		nonce := utils.RandStringBytesMaskImprSrcUnsafe(9)
		operateId := fmt.Sprintf("%s:%s", gconv.String(poapInfo.PoapId), nonce)

		timestamp := time.Now().UnixNano() / 1e6

		signData := map[string]interface{}{
			"appId":     conf.AppId,
			"contract":  conf.ChainAddr,
			"nonce":     nonce,
			"operateId": operateId,
			"remark":    "did",
			"sign":      "",
			"tokenInfo": tokenInfo,
			"timestamp": timestamp,
		}

		sign := utils.Sign(conf.AppSecret, signData)
		signData["sign"] = sign
		data, _ := json.Marshal(signData)

		//调用上链接口
		for i := 1; i <= 3; i++ {
			err := callUpChain(data)
			g.Log().Infof(nil, "UpChain params: %+v, 结果:%+v", signData, err)
			if err == nil {
				break
			}
			time.Sleep(time.Second)
		}
	}

	return
}

func callUpChain(signData []byte) (err error) {
	conf := getChainConf()
	resp, err := http.Post(conf.UpChainUrl, "application/json", bytes.NewBuffer(signData))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		var errmsg map[string]interface{}
		_ = json.Unmarshal(result, &errmsg)
		g.Log().Error(nil, errmsg)
		err = errors.New("内部接口错误")
		return
	}
	var res = struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{}
	err = json.Unmarshal(result, &res)
	if err != nil {
		return
	}
	if res.Code != 0 {
		err = errors.New(res.Message)
		return
	}
	return
}

type chainConf struct {
	Name       string `json:"name"`
	AppId      string `json:"appId"`
	AppSecret  string `json:"appSecret"`
	ChainAddr  string `json:"chainAddr"`
	TokenUrl   string `json:"tokenUrl"`
	UpChainUrl string `json:"upChainUrl"`
}

func getChainConf() (conf *chainConf) {
	ret, err := g.Cfg().Get(nil, "chain")
	if err != nil {
		return
	}
	err = ret.Struct(&conf)
	if err == nil {
		return
	}
	return
}
