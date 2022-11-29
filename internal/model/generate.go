package model

import "strconv"

// GenerateTokenReq 铸造TOKEN
type GenerateTokenReq struct {
	PoapId uint `json:"poapId"`
	Num    uint `json:"num"`
}

// GenerateTokenRes 铸造出的token
type GenerateTokenRes struct {
	No           uint64
	ErrorMessage string
	TokenId      string
}

type TokenId string

// Uint64ToTokenId uint64 转 TokenId
func (TokenId) Uint64ToTokenId(id uint64) TokenId {
	return TokenId(strconv.FormatUint(id, 10))
}

func (TokenId) StringToTokenId(id string) TokenId {
	return TokenId(id)
}

func (id TokenId) ToStr() string {
	return string(id)
}
