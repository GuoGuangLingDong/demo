package model

type GetMyPoapInput struct {
	UId int64
}

type GetPoapDetailsInput struct {
	PoapId int64
}

type GetMainPagePoap struct {
	From  int64
	Count int64
}