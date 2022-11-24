// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Poap is the golang structure for table poap.
type Poap struct {
	PoapId           uint   `json:"poapId"           ` // Poap id
	Miner            string `json:"miner"            ` // Miner
	PoapName         string `json:"poapName"         ` // Poap name
	PoapNumber       int    `json:"poapNumber"       ` // Poap number
	ReceiveCondition int    `json:"receiveCondition" ` // Receive condition
	CoverPic         string `json:"coverPic"         ` // Cover picture
	PoapIntro        string `json:"poapIntro"        ` // Poap introduction
	FavourNumber     uint   `json:"favourNumber"     ` // Favour_number
}
