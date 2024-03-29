package model

type GetMyPoapInput struct {
	UId   string
	From  int
	Count int
}

type GetPoapsDetailsInput struct {
	PoapIds []string
	Uid     string
}

type GetMainPagePoap struct {
	From      int64
	Count     int64
	Condition string
}

type CollectPoapInput struct {
	PoapId     string
	Endorse    string
	EndorsePic string
}

type MintPoapInput struct {
	PoapName    string
	PoapSum     int64
	ReceiveCond int64
	CoverImg    string
	PoapIntro   string
	MintPlat    int
	CollectList string
	Type        int
	Status      int
	Miner       string
	SeriesId    string
}
