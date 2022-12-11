package controller

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model/entity"
	vcodeService "demo/internal/service/vcode"
	"demo/internal/utils"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"

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

	// check code
	if env, _ := g.Cfg().Get(ctx, "system.env"); env.String() != "test" {
		err = vcodeService.VerifyCode(req.PhoneNumber, req.VerifyCode, vcodeService.REGIST_CODE)
		if err != nil {
			return
		}
	}

	userDid := c.generateDid()
	for !service.User().DidExists(ctx, model.DidCreateInput{Did: userDid}) {
		userDid = c.generateDid()
	}

	err = service.User().Create(ctx, model.UserCreateInput{
		UId:         uid,
		Did:         userDid,
		UserName:    userDid,
		NickName:    userDid,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		InviteCode:  req.InviteCode,
	})
	if err == nil {
		service.User().RecordScore(ctx, 200, 6, uid)
	}

	if err == nil {
		vcodeService.DeleteVcode(req.PhoneNumber, vcodeService.REGIST_CODE)
	}

	return
}

func (c *cUser) generateDid() string {
	return fmt.Sprintf("%05v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000))
}

// SignIn is the API for user sign in.
func (c *cUser) SignIn(ctx context.Context, req *v1.UserSignInReq) (res *v1.UserSignInRes, err error) {
	err = utils.ImageCode.VerifyCaptcha(req.ImageVerifyId, req.ImageVerify)
	if err != nil {
		return nil, err
	}
	res = &v1.UserSignInRes{SessionId: ""}
	err, res.SessionId = service.User().SignIn(ctx, model.UserSignInInput{
		PhoneNumber: req.Phonenumber,
		Password:    req.Password,
	})
	return
}

// ResetPassword is the API for user reset password
func (c *cUser) ResetPassword(ctx context.Context, req *v1.UserResetPasswordReq) (res *v1.UserResetPasswordRes, err error) {
	// check code
	if env, _ := g.Cfg().Get(ctx, "system.env"); env.String() != "test" {
		err = vcodeService.VerifyCode(req.PhoneNumber, req.VerifyCode, vcodeService.RESET_CODE)
		if err != nil {
			return
		}
	}
	err = service.User().ResetPassword(ctx, model.ResetPasswordInput{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
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
	var user *entity.User
	if service.User() != nil {
		user = service.User().GetProfile(ctx)
	}
	if req.Did != "" {
		user = GetUserByDid(ctx, req.Did)
	}
	res = &v1.UserProfileRes{
		UserInfo: &v1.UserInfo{
			Id:           user.Id,
			Uid:          user.Uid,
			Username:     user.Username,
			Nickname:     user.Nickname,
			Introduction: user.Introduction,
			Avatar:       user.Avatar,
			Did:          user.Did,
		},
		FollowCount: service.User().GetFollower(ctx, user.Uid),
		PoapCount:   service.User().GetPoapCount(ctx, user.Uid),
		Links:       service.User().GetLink(ctx, user.Uid),
		PoapList:    service.User().GetPoapList(ctx, user.Uid, req.From, req.Count),
	}
	return
}

// ShareInfo is the API for user share info
func (c *cUser) ShareInfo(ctx context.Context, req *v1.UserShareReq) (res *v1.UserShareRes, err error) {
	user := service.User().GetProfile(ctx)
	res = &v1.UserShareRes{
		Did:         user.Did,
		Username:    user.Username,
		UserDesc:    user.Introduction,
		Avatar:      user.Avatar,
		FollowCount: service.User().GetFollower(ctx, user.Uid),
		NftCount:    service.User().GetPoapCount(ctx, user.Uid),
	}
	return
}

func (c *cUser) EditProfile(ctx context.Context, req *v1.EditUserProfileReq) (res *v1.EditUserProfileRes, err error) {
	err = service.User().EditUserProfile(ctx, req)
	return
}

func (c *cUser) GetUserFollow(ctx context.Context, req *v1.GetUserFollowReq) (res *v1.GetUserFollowRes, err error) {
	res = service.User().GetUserFollow(ctx, req)
	return
}
func (*cUser) GetUserFollower(ctx context.Context, req *v1.GetUserFollowerReq) (res *v1.GetUserFollowerRes, err error) {
	res = service.User().GetUserFollowers(ctx, req)
	return
}
func (*cUser) GetUserFollowee(ctx context.Context, req *v1.GetUserFolloweeReq) (res *v1.GetUserFolloweeRes, err error) {
	res = service.User().GetUserFollowees(ctx, req)
	return
}

func (c *cUser) FollowUser(ctx context.Context, req *v1.FollowUserReq) (res *v1.FollowUserRes, err error) {
	err = service.User().FollowUser(ctx, req)
	return
}

func (c *cUser) UnfollowUser(ctx context.Context, req *v1.UnfollowUserReq) (res *v1.UnfollowUserRes, err error) {
	err = service.User().UnfollowUser(ctx, req)
	return
}

func (c *cUser) GetUserScore(ctx context.Context, req *v1.GetUserScoreReq) (res *v1.GetUserScoreRes, err error) {
	res = service.User().GetUserScore(ctx, req)
	return
}

func legalCheck(ctx context.Context, phoneNumber string) error {
	if len(phoneNumber) == 0 {
		return gerror.New("phone number is empty")
	}
	//TODO phone number check
	return nil
}

func GetUserByDid(ctx context.Context, did string) *entity.User {
	user := (*entity.User)(nil)
	dao.User.Ctx(ctx).Where("did", did).Scan(&user)
	return user
}
