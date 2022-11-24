// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"

	"demo/internal/model"
	"demo/internal/model/entity"
)

type IUser interface {
	Create(ctx context.Context, in model.UserCreateInput) (err error)
	IsSignedIn(ctx context.Context) bool
	SignIn(ctx context.Context, in model.UserSignInInput) (err error)
	SignOut(ctx context.Context) error
	UsernameLegalCheck(ctx context.Context, nickname string) (bool, error)
	GetProfile(ctx context.Context) *entity.User
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
