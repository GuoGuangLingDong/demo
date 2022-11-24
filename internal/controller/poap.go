package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/model"
	"demo/internal/service"
	"fmt"
)

var Poap = cPoap{}

type cPoap struct{}

// SignUp is the API for get user poap list
func (c *cPoap) GetMyPoap(ctx context.Context, req *v1.MyPoapReq) (res *v1.MyPoapRes, err error) {
	//err = service.User().Create(ctx, model.UserCreateInput{
	//	Passport: req.Passport,
	//	Password: req.Password,
	//	Nickname: req.Nickname,
	//})
	//
	//err = service.
	user := service.Session().GetUser(ctx)
	res = &v1.MyPoapRes{Res: nil}
	res.Res = service.Poap().GetMyPoap(ctx, model.GetMyPoapInput{UId: user.Uid})
	fmt.Println(res.Res)
	return
}
