package vcodeService

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/parnurzeal/gorequest"
)

type sendCodeRet struct {
	BusinessData struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"BusinessData"`
}

type sendCodeReq struct {
	AccessKey            string            `json:"accessKey"`
	AccessSecret         string            `json:"accessSecret"`
	ClassificationSecret string            `json:"classificationSecret"`
	SignCode             string            `json:"signCode"`
	TemplateCode         string            `json:"templateCode"`
	Phone                string            `json:"phone"`
	Params               map[string]string `json:"params"`
}

var (
	ipMap sync.Map
)

const UCENTER_CODE = "UCENTER_CODE"
const REGIST_CODE = "register"

func DeleteVcode(phone, from string) {
	key := UCENTER_CODE + phone + from
	//删掉验证码
	g.Redis().Do(nil, "DEL", key)
}

func VerifyCode(phone, vcode, from string) (err error) {
	key := UCENTER_CODE + phone + from
	var gv *gvar.Var
	gv, err = g.Redis().Do(nil, "GET", key)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}
	if gv.IsEmpty() {
		err = fmt.Errorf("验证码已过期")
		return
	}
	if gv.String() != vcode {
		err = fmt.Errorf("验证码无效")
		return
	}
	return
}

func Send(phone, from string) (err error) {
	key := UCENTER_CODE + phone + from
	var gv *gvar.Var
	gv, err = g.Redis().Do(nil, "GET", key)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}
	if !gv.IsEmpty() {
		err = fmt.Errorf("获取验证码频繁")
		return
	}
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	smsConf, err := g.Cfg().Get(nil, "sms")
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}
	smsConfMap := smsConf.MapStrStr()
	req := sendCodeReq{
		AccessKey:            smsConfMap["accessKey"],
		AccessSecret:         smsConfMap["accessSecret"],
		ClassificationSecret: smsConfMap["classificationSecret"],
		SignCode:             smsConfMap["signCode"],
		TemplateCode:         smsConfMap["templateCode"],
		Phone:                phone,
		Params:               map[string]string{"code": code},
	}

	ret := sendCodeRet{}
	sendbByte, _ := json.Marshal(req)
	_, _, errs := gorequest.New().Post(smsConfMap["url"]).Timeout(time.Second*10).AppendHeader("Content-Type", "application/json; encoding=utf-8").SendString(string(sendbByte)).EndStruct(&ret)
	if len(errs) != 0 {
		err = fmt.Errorf("失败-1")
		// g.Log().Errorf("发送验证码失败：%v", errs[0])
		return
	}
	if ret.BusinessData.Code != 10000 {
		// g.Log().Errorf("发送验证码失败：%s", ret.BusinessData.Msg)
		GetFailCount(ret.BusinessData.Msg, phone)
		err = fmt.Errorf("获取验证码次数超限，0点后可重新获取")
		return
	}
	_, err = g.Redis().Do(nil, "SET", key, code, "ex", 900)
	if err != nil {
		// g.Log().Errorf("发送验证码失败：%v", err)
		err = fmt.Errorf("失败-3")
		return
	}
	return
}

func GetFailCount(msg, phone string) {
	count, _ := ipMap.Load(phone)
	if count == nil {
		count = 0
	}
	count = count.(int) + 1
	ipMap.Store(phone, count)

	g.Log().Infof(nil, "发送验证码失败：%s,phone:%s,次数:%v", msg, phone, count)

}
