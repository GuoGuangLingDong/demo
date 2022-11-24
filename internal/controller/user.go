package controller

import (
	"context"
	"github.com/google/uuid"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "demo/api/v1"
	"demo/internal/model"
	"demo/internal/service"
)

// User is the controller for user.
var User = cUser{}

type cUser struct{}

// SignUp is the API for user sign up.
func (c *cUser) SignUp(ctx context.Context, req *v1.UserSignUpReq) (res *v1.UserSignUpRes, err error) {
	uid := uuid.NewString()
	uid = strings.Replace(uid, "-", "", -1)

	if err = legalCheck(ctx, req.PhoneNumber); err != nil {
		return
	}

	err = service.User().Create(ctx, model.UserCreateInput{
		UId:         uid,
		Username:    req.Username,
		Password:    req.Password,
		Nickname:    req.Nickname,
		PhoneNumebr: req.PhoneNumber,
	})
	return
}

// SignIn is the API for user sign in.
func (c *cUser) SignIn(ctx context.Context, req *v1.UserSignInReq) (res *v1.UserSignInRes, err error) {
	err = service.User().SignIn(ctx, model.UserSignInInput{
		Username: req.Username,
		Password: req.Password,
	})
	return
}

// IsSignedIn checks and returns whether the user is signed in.
func (c *cUser) IsSignedIn(ctx context.Context, req *v1.UserIsSignedInReq) (res *v1.UserIsSignedInRes, err error) {
	res = &v1.UserIsSignedInRes{
		OK: service.User().IsSignedIn(ctx),
	}
	return
}

// SignOut is the API for user sign out.
func (c *cUser) SignOut(ctx context.Context, req *v1.UserSignOutReq) (res *v1.UserSignOutRes, err error) {
	err = service.User().SignOut(ctx)
	return
}

// CheckUserName checks and returns whether the user nickname is available.
func (c *cUser) CheckUserName(ctx context.Context, req *v1.UserCheckNickNameReq) (res *v1.UserCheckNickNameRes, err error) {
	available, err := service.User().UsernameLegalCheck(ctx, req.Nickname)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Nickname "%s" is already token by others`, req.Nickname)
	}
	return
}

// Profile returns the user profile.
func (c *cUser) Profile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	res = &v1.UserProfileRes{
		User: service.User().GetProfile(ctx),
	}
	return
}

func legalCheck(ctx context.Context, phoneNumber string) error {
	if len(phoneNumber) == 0 {
		return gerror.New("phone number is empty")
	}
	//TODO phone number check
	return nil
}
