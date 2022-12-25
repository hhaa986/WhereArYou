package model

type CreateWRequest struct {
	WName        string `json:"name" example:"Black Label"`
	AlcoholLevel int64  `json:"alcoholLevel" example:"40"`
	Origin       string `json:"origin" example:"Germany"'`
	CID          uint   `json:"cid" example:"1"`
}
