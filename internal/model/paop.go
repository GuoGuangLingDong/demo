package model

type GetMyPoapInput struct {
	UId   string
	From  int
	Count int
}

type GetPoapDetailsInput struct {
	PoapId string
}

type GetMainPagePoap struct {
	From      int64
	Count     int64
	Condition string
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
