package model

type WhiskeyRequest struct {
	WName        string `json:"name" example:"Black Label"`
	AlcoholLevel int64  `json:"alcoholLevel" example:"40"`
	Origin       string `json:"origin" example:"Germany"'`
	CID          uint   `json:"cid" example:"1"`
}
type WhiskeyUpdateRequest struct {
	Id           uint   `json:"id"`
	WName        string `json:"name" example:"Black Label"`
	AlcoholLevel int64  `json:"alcoholLevel" example:"40"`
	Origin       string `json:"origin" example:"Germany"'`
	CID          uint   `json:"cid" example:"1"`
}

type WCategoryRequest struct {
	WName string `json:"name" example:"Black Label"`
}
type WCategoryUpdateRequest struct {
	Id    uint   `json:"id"`
	WName string `json:"name" example:"Black Label"`
}
