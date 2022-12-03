package controller

import (
	"context"
	v1 "demo/api/v1"
	vcodeService "demo/internal/service/vcode"
	"demo/internal/utils"
)

var Code = cCode{}

type cCode struct{}

// GetImageVerify 获取图形验证码
func (c *cCode) GetImageVerify(ctx context.Context, req *v1.GetImageVerifyReq) (res *v1.GetImageVerifyRes, err error) {
	id, base64, err := utils.ImageCode.CreateCode()
	if err != nil {
		return
	}
	res = &v1.GetImageVerifyRes{
		Id:     id,
		Base64: base64,
	}
	return
}

// CodeSend 发送验证码
func (c *cCode) CodeSend(ctx context.Context, req *v1.CodeSendReq) (res *v1.CodeSendRes, err error) {
	//err = utils.ImageCode.VerifyCaptcha(req.ImageVerifyId, req.ImageVerify)
	//if err != nil {
	//	return
	//}
	err = vcodeService.Send(req.Phone, req.From)
	if err != nil {
		return
	}
	return
}
