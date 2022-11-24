package controller

import (
	"context"
	v1 "demo/api/v1"
)

var Poap = cPoap{}

type cPoap struct{}

// SignUp is the API for get user poap list
func (c *cPoap) GetUserPoap(ctx context.Context, req *v1.MyPoapReq) (res *v1.MyPoapRes, err error) {
	//err = service.User().Create(ctx, model.UserCreateInput{
	//	Passport: req.Passport,
	//	Password: req.Password,
	//	Nickname: req.Nickname,
	//})
	//
	//err = service.
	return
}
