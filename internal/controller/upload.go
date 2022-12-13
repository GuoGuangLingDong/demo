package controller

import (
	"context"
	v1 "demo/api/v1"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"path"
)

var Upload = cUpload{}

type cUpload struct{}

func (c *cUpload) GetUploadToken(ctx context.Context, req *v1.GetUploadTokenReq) (res *v1.GetUploadTokenRes, err error) {
	conf, err := g.Cfg().Get(ctx, "upload")
	confMap := conf.MapStrStr()

	putPolicy := storage.PutPolicy{
		Scope: confMap["bucket"],
	}

	putPolicy.Expires = 7200 //2小时有效期
	mac := qbox.NewMac(confMap["accessKey"], confMap["secretKey"])

	upToken := putPolicy.UploadToken(mac)

	key := fmt.Sprintf("did/%s", uuid.NewString()+path.Ext(req.Name))

	res = &v1.GetUploadTokenRes{
		UploadKey: key,
		Token:     upToken,
		Cdn:       confMap["cdn"],
	}
	return
}
