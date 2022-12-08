package model

type UserCreateInput struct {
	UId         string
	Did         string
	Username    string
	Password    string
	Nickname    string
	PhoneNumebr string
	InviteCode  string
}

type UserSignInInput struct {
	Username string
	Password string
}

type DidCreateInput struct {
	Did string
}
