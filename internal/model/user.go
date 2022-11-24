package model

type UserCreateInput struct {
	UId         string
	Username    string
	Password    string
	Nickname    string
	PhoneNumebr string
}

type UserSignInInput struct {
	Username string
	Password string
}
