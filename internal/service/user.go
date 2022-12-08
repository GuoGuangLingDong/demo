// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/model"
	"demo/internal/model/entity"
)

type (
	IUser interface {
		Create(ctx context.Context, in model.UserCreateInput) (err error)
		DidExists(ctx context.Context, in model.DidCreateInput) bool
		IsSignedIn(ctx context.Context) bool
		SignIn(ctx context.Context, in model.UserSignInInput) (err error)
		SignOut(ctx context.Context) error
		UsernameLegalCheck(ctx context.Context, username string) (bool, error)
        GetProfile(ctx context.Context) *entity.User
		GetLink(ctx context.Context, uid string) *v1.Link
		GetFollower(ctx context.Context, uid string) int64
		GetPoapCount(ctx context.Context, uid string) int64
		EditUserProfile(ctx context.Context, in *v1.EditUserProfileReq) (err error)
		GetUserFollow(ctx context.Context, in *v1.GetUserFollowReq) *v1.GetUserFollowRes
		FollowUser(ctx context.Context, in *v1.FollowUserReq) (err error)
		UnfollowUser(ctx context.Context, in *v1.UnfollowUserReq) (err error)
		GetUserScore(ctx context.Context, req *v1.GetUserScoreReq) *v1.GetUserScoreRes
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
