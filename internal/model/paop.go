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

type CollectPoapInput struct {
	PoapId int64
}

type MintPoapInput struct {
	PoapName    string
	PoapSum     int64
	ReceiveCond int64
	CoverImg    string
	PoapIntro   string
}
