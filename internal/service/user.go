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
		IsSignedIn(ctx context.Context) bool
		SignIn(ctx context.Context, in model.UserSignInInput) (err error)
		SignOut(ctx context.Context) error
		UsernameLegalCheck(ctx context.Context, username string) (bool, error)
        GetProfile(ctx context.Context) *entity.User
		GetLink(ctx context.Context, uid uint) *v1.Link
		GetFollow(ctx context.Context, uid uint) int64
		GetPoapCount(ctx context.Context, uid uint) int64
		EditUserProfile(ctx context.Context, in *v1.EditUserProfileReq) (err error)
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
