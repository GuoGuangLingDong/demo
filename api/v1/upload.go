package v1

import "github.com/gogf/gf/v2/frame/g"

type GetUploadTokenReq struct {
	g.Meta `path:"/upload/token" method:"post" tags:"UploadService" summary:"Get Upload Token"`
	Name   string `json:"name" form:"name"`
}

type GetUploadTokenRes struct {
	UploadKey string `json:"uploadKey"`
	Token     string `json:"token"`
	Cdn       string `json:"cdn"`
}
