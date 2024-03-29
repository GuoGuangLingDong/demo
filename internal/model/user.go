package model

type UserCreateInput struct {
	UId         string
	Did         string
	UserName    string
	NickName    string
	Password    string
	PhoneNumber string
	InviteCode  string
}

type UserSignInInput struct {
	PhoneNumber string
	Password    string
}

type DidCreateInput struct {
	Did string
}

type ResetPasswordInput struct {
	Password    string
	PhoneNumber string
}
