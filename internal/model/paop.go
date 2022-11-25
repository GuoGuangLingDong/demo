package model

type GetMyPoapInput struct {
	UId string
}

type GetPoapDetailsInput struct {
	PoapId int64
}

type GetMainPagePoap struct {
	From  int64
	Count int64
}
