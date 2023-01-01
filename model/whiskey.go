package model

import "gorm.io/gorm"

type Whiskey struct {
	gorm.Model
	Name         string    //품명
	AlcoholLevel int64     //도수
	Origin       string    //원산지
	CID          uint      //위스키 카테고리 no
	Reviews      []WReview `gorm:"foreignKey:WID"`
}

type WCategory struct {
	gorm.Model
	Name string //분류명
}

type WReview struct {
	gorm.Model
	Price    float32 //가격
	Location string  //위치
	WID      uint    //위스키 no
}
