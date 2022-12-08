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
	Password    string `v:"required|length:6,16"`
	Password2   string `v:"required|length:6,16|same:Password"`
	PhoneNumber string
	VerifyCode  string `v:"required"`
	InviteCode  string
}
type UserSignUpRes struct{}

type UserSignInReq struct {
	g.Meta        `path:"/user/sign-in" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	Phonenumber   string `v:"required"`
	Password      string `v:"required"`
	ImageVerify   string `v:"required"`
	ImageVerifyId string `v:"required"`
}
type UserSignInRes struct{}

type UserResetPasswordReq struct {
	g.Meta      `path:"/user/reset-password" method:"post" tags:"UserService" summary:"Reset user's password"`
	Password    string `v:"required|length:6,16"`
	Password2   string `v:"required|length:6,16|same:Password"`
	PhoneNumber string
	VerifyCode  string `v:"required"`
}
type UserResetPasswordRes struct{}

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

type GetUserFollowReq struct {
	g.Meta `path:"/user/follow" method:"get" tags:"UserService" summary:"Get the follow information of current user"`
}

type GetUserFollowRes struct {
	Followee []*FollowInformation `json:"followee,omitempty"`
	Follower []*FollowInformation `json:"follower,omitempty"`
}

type FollowInformation struct {
	Username    string `json:"username,omitempty"`
	Uid         string `json:"uid,omitempty"`
	FollowCount int    `json:"follow_count,omitempty"`
	PoapCount   int    `json:"poap_count,omitempty"`
}

type FollowUserReq struct {
	g.Meta `path:"/user/follow" method:"post" tags:"UserService" summary:"Follow current user"`
	Uid    uint `json:"uid,omitempty"`
}

type FollowUserRes struct {
}

type UnfollowUserReq struct {
	g.Meta `path:"/user/unfollow" method:"post" tags:"UserService" summary:"Unfollow current user"`
	Uid    uint `json:"uid,omitempty"`
}

type UnfollowUserRes struct {
}

type GetUserScoreReq struct {
	g.Meta `path:"/user/score" method:"get" tags:"UserService" summary:"Get the score of current user"`
}

type GetUserScoreRes struct {
	Score     int64               `json:"score,omitempty"`
	Oprations []*entity.Operation `json:"oprations,omitempty"`
}

type UserShareReq struct {
	g.Meta `path:"/user/share" method:"get" tags:"UserService" summary:"User share info"`
}

type UserShareRes struct {
	Uid         string `json:"uid"`
	Username    string `json:"username"`
	UserDesc    string `json:"user_desc"`
	Avatar      string `json:"avatar"`
	FollowCount int64  `json:"follow_count"`
	NftCount    int64  `json:"nft_count"`
}
