package v1

import "github.com/gogf/gf/v2/frame/g"

type GetImageVerifyReq struct {
	g.Meta `path:"/code/image" method:"post" tags:"CodeService" summary:"Get image verify"`
}

type GetImageVerifyRes struct {
	Id     string `json:"id"`
	Base64 string `json:"base64"`
}

type CodeSendReq struct {
	g.Meta `path:"/code/send" method:"post" tags:"CodeService" summary:"Get code"`
	Phone  string `json:"phone" v:"required"`
	From   string `json:"from" v:"required"`
	//ImageVerifyId string `json:"imageVerifyId" v:"required"`
	//ImageVerify   string `json:"imageVerify" v:"required"`
}

type CodeSendRes struct {
}
