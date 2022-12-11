package v1

import (
	"demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type UserProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"UserService" summary:"Get the profile of current user" `
	From   int    `json:"from,omitempty"`
	Count  int    `json:"count,omitempty"`
	Uid    string `json:"uid,omitempty"`
}
type poapsDetail []*PoapDetailPoapRes
type UserProfileRes struct {
	*UserInfo
	FollowCount int64              `json:"follow_count,omitempty"`
	PoapCount   int64              `json:"poap_count,omitempty"`
	Links       []*entity.Userlink `json:"links,omitempty"`
	PoapList    poapsDetail        `json:"poap_list,omitempty"`
}

type UserInfo struct {
	Id           uint   `json:"id"           ` // pk
	Uid          string `json:"uid"          ` // User ID
	Username     string `json:"username"     ` // User Name
	Nickname     string `json:"nickname"     ` // User Nickname
	PhoneNumber  string `json:"phoneNumber"  ` // Phone Number
	WechatNumber string `json:"wechatNumber" ` // Wechat Number
	InviteCode   string `json:"inviteCode"   ` // Invite Code
	Introduction string `json:"introduction" ` // Introduction
	Avatar       string `json:"avatar"       ` // 头像
	Did          string `json:"did"          ` // User DID
}

type EditUserProfileReq struct {
	g.Meta       `path:"/user/profile" method:"post" tags:"UserService" summary:"Edit the profile of current user"`
	UserName     string  `json:"username,omitempty" json:"user_name,omitempty"`
	Introduction string  `json:"introduction,omitempty" json:"introduction,omitempty"`
	Links        []*Link `json:"links,omitempty"`
	Avatar       string  `json:"avatar,omitempty" json:"avatar,omitempty"`
}

type EditUserProfileRes struct {
}

type Link struct {
	LinkTitle string `json:"link_title,omitempty"`
	Link      string `json:"link,omitempty"`
	LinkType  string `json:"link_type,omitempty"`
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
type UserSignInRes struct {
	SessionId string `json:"sessionId,omitempty"`
}

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
	g.Meta `path:"/user/follow_all" method:"get" tags:"UserService" summary:"Get the follow information of current user"`
}

type GetUserFollowerReq struct {
	g.Meta `path:"/user/followers" method:"get" tags:"UserService" summary:"Get the follower information of current user"`
	From   int `json:"from"`
	Count  int `json:"count"`
}

type GetUserFollowerRes struct {
	Follower []*FollowInformation `json:"list,omitempty"`
}

type GetUserFolloweeReq struct {
	g.Meta `path:"/user/followees" method:"get" tags:"UserService" summary:"Get the followee information of current user"`
	From   int `json:"from"`
	Count  int `json:"count"`
}
type GetUserFolloweeRes struct {
	Followee []*FollowInformation `json:"list,omitempty"`
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
	Avatar      string `json:"avatar"`
	Did         string `json:"did"`
}

type FollowUserReq struct {
	g.Meta `path:"/user/follow" method:"post" tags:"UserService" summary:"Follow current user"`
	Uid    string `json:"uid,omitempty"`
}

type FollowUserRes struct {
}

type UnfollowUserReq struct {
	g.Meta `path:"/user/unfollow" method:"post" tags:"UserService" summary:"Unfollow current user"`
	Uid    string `json:"uid,omitempty"`
}

type UnfollowUserRes struct {
}

type GetUserScoreReq struct {
	g.Meta `path:"/user/score" method:"get" tags:"UserService" summary:"Get the score of current user"`
}

type GetUserScoreRes struct {
	Score     int64        `json:"score,omitempty"`
	Oprations []*Operation `json:"oprations,omitempty"`
}
type Operation struct {
	*entity.Operation
	Opt      string `json:"opt"`
	Opt_time string `json:"opt_time"`
}

type UserShareReq struct {
	g.Meta `path:"/user/share" method:"get" tags:"UserService" summary:"User share info"`
}

type UserShareRes struct {
	Did         string `json:"did"`
	Username    string `json:"username"`
	UserDesc    string `json:"user_desc"`
	Avatar      string `json:"avatar"`
	FollowCount int64  `json:"follow_count"`
	NftCount    int64  `json:"nft_count"`
}
