package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"demo/internal/model/entity"
)

type UserProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"UserService" summary:"Get the profile of current user"`
}
type poapsDetail []*PoapDetailPoapRes
type UserProfileRes struct {
	*entity.User
	FollowCount int64
	PoapCount   int64
	Links       *Link
	poapsDetail
}

type EditUserProfileReq struct {
	g.Meta       `path:"/user/profile" method:"post" tags:"UserService" summary:"Edit the profile of current user"`
	UserName     string `json:"user_name,omitempty" json:"user_name,omitempty"`
	Introduction string `json:"introduction,omitempty" json:"introduction,omitempty"`
	Links        *Link  `json:"links,omitempty"`
	Avatar       string `json:"avatar,omitempty" json:"avatar,omitempty"`
}

type EditUserProfileRes struct {
}

type Link struct {
	TiktokLink   string `json:"tiktok_link,omitempty"`
	InsLink      string `json:"ins_link,omitempty"`
	WeiboLink    string `json:"weibo_link,omitempty"`
	RedLink      string `json:"red_link,omitempty"`
	WechatLink   string `json:"wechat_link,omitempty"`
	TelLink      string `json:"tel_link,omitempty"`
	TweetLink    string `json:"tweet_link,omitempty"`
	FacebookLink string `json:"facebook_link,omitempty"`
	LinkedinLink string `json:"linkedin_link,omitempty"`
}

type UserSignUpReq struct {
	g.Meta      `path:"/user/sign-up" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	Username    string `v:"required|length:6,16"`
	Password    string `v:"required|length:6,16"`
	Password2   string `v:"required|length:6,16|same:Password"`
	Nickname    string
	PhoneNumber string
}
type UserSignUpRes struct{}

type UserSignInReq struct {
	g.Meta   `path:"/user/sign-in" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	Username string `v:"required"`
	Password string `v:"required"`
}
type UserSignInRes struct{}

type UserCheckPassportReq struct {
	g.Meta   `path:"/user/check-passport" method:"post" tags:"UserService" summary:"Check passport available"`
	Passport string `v:"required"`
}
type UserCheckPassportRes struct{}

type UserCheckNickNameReq struct {
	g.Meta   `path:"/user/check-passport" method:"post" tags:"UserService" summary:"Check nickname available"`
	Nickname string `v:"required"`
}
type UserCheckNickNameRes struct{}

type UserIsSignedInReq struct {
	g.Meta `path:"/user/is-signed-in" method:"post" tags:"UserService" summary:"Check current user is already signed-in"`
}
type UserIsSignedInRes struct {
	OK bool `dc:"True if current user is signed in; or else false"`
}

type UserSignOutReq struct {
	g.Meta `path:"/user/sign-out" method:"post" tags:"UserService" summary:"Sign out current user"`
}
type UserSignOutRes struct{}
