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
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
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
	if in.Nickname == "" {
		in.Nickname = fmt.Sprintf("wesoul-%v", in.UId[:6])
	}
	// Username checks.
	available, err := s.UsernameLegalCheck(ctx, in.Username)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	}
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Uid:         in.UId,
			PhoneNumber: in.PhoneNumebr,
			Username:    in.Username,
			Password:    in.Password,
			Nickname:    in.Nickname,
		}).Insert()
		return err
	})
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *SUser) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// SignIn creates session for given user account.
func (s *SUser) SignIn(ctx context.Context, in model.UserSignInInput) (err error) {
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		Username: in.Username,
		Password: in.Password,
	}).Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New(`Passport or Password not correct`)
	}
	if err = service.Session().SetUser(ctx, user); err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
	})
	return nil
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

func (s *SUser) GetLink(ctx context.Context, uid string) *v1.Link {
	links := ([]*entity.Userlink)(nil)
	dao.Userlink.Ctx(ctx).Where("uid", uid).Scan(&links)
	res := &v1.Link{
		TiktokLink:   "",
		InsLink:      "",
		WeiboLink:    "",
		RedLink:      "",
		WechatLink:   "",
		TelLink:      "",
		TweetLink:    "",
		FacebookLink: "",
		LinkedinLink: "",
	}
	for _, link := range links {
		switch link.LinkType {
		case 1:
			res.TiktokLink = link.Link
		case 2:
			res.WeiboLink = link.Link
		case 3:
			res.RedLink = link.Link
		case 4:
			res.WechatLink = link.Link
		case 5:
			res.TelLink = link.Link
		case 6:
			res.InsLink = link.Link
		case 7:
			res.TweetLink = link.Link
		case 8:
			res.FacebookLink = link.Link
		case 9:
			res.LinkedinLink = link.Link
		}
	}
	return res
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

// todo 加事务
func (s *SUser) EditUserProfile(ctx context.Context, in *v1.EditUserProfileReq) (err error) {
	user := service.Session().GetUser(ctx)

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
	err = service.Session().SetUser(ctx, user)
	if err != nil {
		return err
	}
	if in.Links != nil {
		if in.Links.TiktokLink != "" {
			_, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.TiktokLink, "link_type": 1}).Insert()
			if err != nil {
				_, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.TiktokLink, "link_type": 1}).Update()
			}
		}
		if in.Links.WeiboLink != "" {
			_, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.WeiboLink, "link_type": 2}).Insert()
			if err != nil {
				_, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.WeiboLink, "link_type": 2}).Update()
			}
		}
		if in.Links.RedLink != "" {
			if _, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.RedLink, "link_type": 3}).Insert(); err != nil {
				_, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.RedLink, "link_type": 3}).Update()
			}
		}
		if in.Links.WechatLink != "" {
			if _, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.WechatLink, "link_type": 4}).Insert(); err != nil {
				_, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.WechatLink, "link_type": 4}).Update()
			}
		}
		if in.Links.TelLink != "" {
			if _, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.TelLink, "link_type": 5}).Insert(); err != nil {
				_, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.TelLink, "link_type": 5}).Update()
			}
		}
		if in.Links.InsLink != "" {
			if _, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.InsLink, "link_type": 6}).Insert(); err != nil {
				_, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.InsLink, "link_type": 6}).Update()
			}
		}
		if in.Links.TweetLink != "" {
			if _, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.TweetLink, "link_type": 7}).Insert(); err != nil {
				_, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.TweetLink, "link_type": 7}).Update()
			}
		}
		if in.Links.FacebookLink != "" {
			if _, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.FacebookLink, "link_type": 8}).Insert(); err != nil {
				_, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.FacebookLink, "link_type": 8}).Update()
			}
		}
		if in.Links.LinkedinLink != "" {
			if _, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.LinkedinLink, "link_type": 9}).Insert(); err != nil {
				_, err = dao.Userlink.Ctx(ctx).Data(g.Map{"uid": user.Uid, "link": in.Links.LinkedinLink, "link_type": 9}).Update()
			}
		}
	}

	// 铸造头像nft
	if user.Avatar != in.Avatar {
		_ = service.Poap().MintPoap(ctx, model.MintPoapInput{
			PoapName:    fmt.Sprintf("%d.did Avatar PFP", user.Uid),
			PoapSum:     1,
			ReceiveCond: 1,
			CoverImg:    in.Avatar,
			PoapIntro:   fmt.Sprintf("%d.did于%s时间更新头像", user.Uid, time.Now().Format(time.RFC3339)),
		})
	}

	return nil
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
		})
	}

	for _, f := range follower {
		followers = append(followers, &v1.FollowInformation{
			Username:    s.GetUserInfo(ctx, f.Follower).Username,
			Uid:         f.Follower,
			FollowCount: int(s.GetFollowee(ctx, f.Follower)),
			PoapCount:   int(s.GetPoapCount(ctx, f.Follower)),
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
	return &v1.GetUserScoreRes{
		Score:     int64(sum),
		Oprations: opts,
	}
}
