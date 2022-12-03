package model

type GetMyPoapInput struct {
	UId string
}

type GetPoapDetailsInput struct {
	PoapId string
}

type GetMainPagePoap struct {
	From  int64
	Count int64
}

type CollectPoapInput struct {
	PoapId string
}

type MintPoapInput struct {
	PoapName    string
	PoapSum     int64
	ReceiveCond int64
	CoverImg    string
	PoapIntro   string
}
