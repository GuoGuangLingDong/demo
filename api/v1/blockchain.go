package v1

import "github.com/gogf/gf/v2/frame/g"

type BlockChainDataReq struct {
	g.Meta  `path:"/blockchain/metadata/:poapId/:tokenId" method:"get" tags:"CodeService" summary:"Get image verify"`
	PoapId  string `json:"poapId"`
	TokenId int    `json:"tokenId"`
}

type BlockChainDataRes struct {
	Name            string                   `json:"name"`
	BatchNo         string                   `json:"batch_no"`
	Description     string                   `json:"description"`
	ExternalUrl     string                   `json:"external_url"`
	Image           string                   `json:"image"`
	ImageData       string                   `json:"image_data"`
	AnimationUrl    string                   `json:"animation_url"`
	BackgroundColor string                   `json:"background_color"`
	Attributes      []map[string]interface{} `json:"attributes"`
}
