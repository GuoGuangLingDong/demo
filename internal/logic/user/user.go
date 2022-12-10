package user

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	SUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() *SUser {
	return &SUser{}
}

// Create creates user account.
func (s *SUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {
	// If Nickname is not specified, generate one
	//if in.Nickname == "" {
	//	in.Nickname = fmt.Sprintf("wesoul-%v", in.UId[:6])
	//}
	// Username checks.
	//available, err := s.UsernameLegalCheck(ctx, in.Username)
	//if err != nil {
	//	return err
	//}
	//if !available {
	//	return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	//}
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Uid:         in.UId,
			PhoneNumber: in.PhoneNumber,
			Password:    in.Password,
			Did:         in.Did,
			Username:    in.Did,
			Nickname:    in.NickName,
		}).Insert()
		return err
	})
}

func (s *SUser) DidExists(ctx context.Context, in model.DidCreateInput) bool {
	count, _ := dao.User.Ctx(ctx).Where("did", in.Did).Count()
	if count != 0 {
		fmt.Println(in.Did, " Exits")
		return false
	}

	return true
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *SUser) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// SignIn creates session for given user account.
func (s *SUser) SignIn(ctx context.Context, in model.UserSignInInput) (err error, sessionId string) {

	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		PhoneNumber: in.PhoneNumber,
		Password:    in.Password,
	}).Scan(&user)
	if err != nil {
		return err, ""
	}

	if user == nil {
		return gerror.New(`Passport or Password not correct`), ""
	}
	if err, sessionId = service.Session().SetUser(ctx, user); err != nil {
		return err, sessionId
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
	})
	return nil, sessionId
}

func (s *SUser) ResetPassword(ctx context.Context, in model.ResetPasswordInput) (err error) {
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		PhoneNumber: in.PhoneNumber,
	}).Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New(`User not exist`)
	}
	if err, _ = service.Session().SetUser(ctx, user); err != nil {
		return err
	}

	_, err = dao.User.Ctx(ctx).Where("uid", user.Uid).Data(g.Map{"password": in.Password}).Update()
	if err != nil {
		return err
	}
	return
}

// SignOut removes the session for current signed-in user.
func (s *SUser) SignOut(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}

// UsernameLegalCheck checks and returns given nickname is available for signing up.
func (s *SUser) UsernameLegalCheck(ctx context.Context, username string) (bool, error) {
	res, err := dao.User.Ctx(ctx).One(do.User{
		Username: username,
	})
	if err != nil {
		return false, err
	}
	return len(res) == 0, nil
}

// GetProfile retrieves and returns current user info in session.
func (s *SUser) GetProfile(ctx context.Context) *entity.User {
	return service.Session().GetUser(ctx)
}

func (s *SUser) GetLink(ctx context.Context, uid string) []*entity.Userlink {
	links := ([]*entity.Userlink)(nil)
	dao.Userlink.Ctx(ctx).Where("uid", uid).Scan(&links)
	return links
}

func (s *SUser) GetFollower(ctx context.Context, uid string) int64 {
	count, err := dao.Follow.Ctx(ctx).Where("followee", uid).Count()
	if err != nil {
		return 0
	}
	return int64(count)
}

func (s *SUser) GetFollowee(ctx context.Context, uid string) int64 {
	count, err := dao.Follow.Ctx(ctx).Where("follower", uid).Count()
	if err != nil {
		return 0
	}
	return int64(count)
}

func (s *SUser) GetPoapCount(ctx context.Context, uid string) int64 {
	count, err := dao.Hold.Ctx(ctx).Where("uid", uid).Count()
	if err != nil {
		return 0
	}
	return int64(count)
}

func (s *SUser) EditUserProfile(ctx context.Context, in *v1.EditUserProfileReq) (err error) {
	user := service.Session().GetUser(ctx)

	dao.Userlink.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(g.Map{
			"username":     in.UserName,
			"introduction": in.Introduction,
			"avatar":       in.Avatar,
		}).Where("uid", user.Uid).Update()
		if err != nil {
			return err
		}
		user.Username = in.UserName
		user.Introduction = in.Introduction
		user.Avatar = in.Avatar
		err, _ = service.Session().SetUser(ctx, user)
		if err != nil {
			return err
		}

		_, err = dao.Userlink.Ctx(ctx).Where("uid = ?", user.Uid).Delete()
		if err != nil {
			return err
		}
		for _, link := range in.Links {
			_, err = dao.Userlink.Ctx(ctx).Data(do.Userlink{
				Uid:       user.Uid,
				Link:      link.Link,
				LinkType:  link.LinkType,
				LinkTitle: link.LinkTitle,
			}).Insert()
			if err != nil {
				return err
			}
		}
		// 铸造头像nft
		if user.Avatar != in.Avatar {
			err = service.Poap().MintPoap(ctx, model.MintPoapInput{
				PoapName:    fmt.Sprintf("%d.did Avatar PFP", user.Uid),
				PoapSum:     1,
				ReceiveCond: 1,
				CoverImg:    in.Avatar,
				PoapIntro:   fmt.Sprintf("%d.did于%s时间更新头像", user.Uid, time.Now().Format(time.RFC3339)),
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

func (s *SUser) GetUserFollowers(ctx context.Context, in *v1.GetUserFollowerReq) *v1.GetUserFollowerRes {
	user := service.Session().GetUser(ctx)
	followers := []*v1.FollowInformation{}
	follower := ([]*entity.Follow)(nil)
	dao.Follow.Ctx(ctx).Where("follower", user.Uid).Scan(&follower) //关注的人
	for _, f := range follower {
		followers = append(followers, &v1.FollowInformation{
			Username:    s.GetUserInfo(ctx, f.Followee).Username,
			Uid:         f.Followee,
			FollowCount: int(s.GetFollowee(ctx, f.Followee)),
			PoapCount:   int(s.GetPoapCount(ctx, f.Followee)),
			Avatar:      s.GetUserInfo(ctx, f.Followee).Avatar,
			Did:         s.GetUserInfo(ctx, f.Followee).Did,
		})
	}
	return &v1.GetUserFollowerRes{
		Follower: followers,
	}
}
func (s *SUser) GetUserFollowees(ctx context.Context, in *v1.GetUserFolloweeReq) *v1.GetUserFolloweeRes {
	user := service.Session().GetUser(ctx)
	followees := []*v1.FollowInformation{}
	followee := ([]*entity.Follow)(nil)
	dao.Follow.Ctx(ctx).Where("followee", user.Uid).Scan(&followee) //粉丝
	for _, f := range followee {
		followees = append(followees, &v1.FollowInformation{
			Username:    s.GetUserInfo(ctx, f.Followee).Username,
			Uid:         f.Followee,
			FollowCount: int(s.GetFollowee(ctx, f.Followee)),
			PoapCount:   int(s.GetPoapCount(ctx, f.Followee)),
			Avatar:      s.GetUserInfo(ctx, f.Followee).Avatar,
			Did:         s.GetUserInfo(ctx, f.Followee).Did,
		})
	}
	return &v1.GetUserFolloweeRes{
		Followee: followees,
	}

}
func (s *SUser) GetUserFollow(ctx context.Context, in *v1.GetUserFollowReq) *v1.GetUserFollowRes {
	user := service.Session().GetUser(ctx)
	followees := []*v1.FollowInformation{}
	followers := []*v1.FollowInformation{}
	followee := ([]*entity.Follow)(nil)
	follower := ([]*entity.Follow)(nil)
	dao.Follow.Ctx(ctx).Where("follower", user.Uid).Scan(&followee) //关注的人
	dao.Follow.Ctx(ctx).Where("followee", user.Uid).Scan(&follower) //粉丝
	for _, f := range followee {
		followees = append(followees, &v1.FollowInformation{
			Username:    s.GetUserInfo(ctx, f.Followee).Username,
			Uid:         f.Followee,
			FollowCount: int(s.GetFollowee(ctx, f.Followee)),
			PoapCount:   int(s.GetPoapCount(ctx, f.Followee)),
			Avatar:      s.GetUserInfo(ctx, f.Followee).Avatar,
		})
	}

	for _, f := range follower {
		followers = append(followers, &v1.FollowInformation{
			Username:    s.GetUserInfo(ctx, f.Follower).Username,
			Uid:         f.Follower,
			FollowCount: int(s.GetFollowee(ctx, f.Follower)),
			PoapCount:   int(s.GetPoapCount(ctx, f.Follower)),
			Avatar:      s.GetUserInfo(ctx, f.Follower).Avatar,
		})
	}

	return &v1.GetUserFollowRes{
		Followee: followees,
		Follower: followers,
	}
}

func (s *SUser) GetUserInfo(ctx context.Context, uid string) *entity.User {
	var user *entity.User
	dao.User.Ctx(ctx).Where("uid", uid).Scan(&user)
	return user
}

func (s *SUser) FollowUser(ctx context.Context, req *v1.FollowUserReq) (err error) {
	user := service.Session().GetUser(ctx)
	fmt.Println("followee", req.Uid)
	_, err = dao.Follow.Ctx(ctx).Data(g.Map{
		"followee": req.Uid,
		"follower": user.Uid,
	}).Insert()
	return
}

func (s *SUser) UnfollowUser(ctx context.Context, req *v1.UnfollowUserReq) (err error) {
	user := service.Session().GetUser(ctx)
	_, err = dao.Follow.Ctx(ctx).Where("followee", req.Uid).Where("follower", user.Uid).Delete()
	return
}

func (s *SUser) GetUserScore(ctx context.Context, req *v1.GetUserScoreReq) *v1.GetUserScoreRes {
	user := service.Session().GetUser(ctx)
	sum, _ := dao.Operation.Ctx(ctx).Where("uid", user.Uid).Sum("score")
	opts := ([]*entity.Operation)(nil)
	dao.Operation.Ctx(ctx).Where("uid", user.Uid).Scan(&opts)
	operations := ([]*v1.Operation)(nil)
	for _, opt := range opts {
		optName := ""
		switch {
		case opt.OptType == 1:
			optName = "铸造"
		case opt.OptType == 2:
			optName = "兑换DID"
		case opt.OptType == 3:
			optName = "点赞"
		case opt.OptType == 4:
			optName = "建立连接（关注）"
		case opt.OptType == 5:
			optName = "领取POAP"
		case opt.OptType == 6:
			optName = "新用户赠送"
		}
		operations = append(operations, &v1.Operation{
			Operation: opt,
			Opt:       optName,
			Opt_time:  opt.CreateAt.String(),
		})
	}
	return &v1.GetUserScoreRes{
		Score:     int64(sum),
		Oprations: operations,
	}

}

func (s *SUser) GetUserByDid(ctx context.Context, did string) *entity.User {
	user := (*entity.User)(nil)
	dao.User.Ctx(ctx).Where("did", did).Scan(&user)
	return user
}

func (s *SUser) GetPoapList(ctx context.Context, uid string, from int, count int) []*v1.PoapDetailPoapRes {
	res := service.Poap().GetMyPoap(ctx, model.GetMyPoapInput{UId: uid, From: from, Count: count})
	return res
}

type Operation struct {
	Uid     string
	OptType int
	Score   int
}

func (s *SUser) SignUpScore(ctx context.Context, uid string) (err error) {
	_, err = dao.Operation.Ctx(ctx).Data(Operation{
		Uid:     uid,
		OptType: 6,
		Score:   800,
	}).Insert()
	return
}
