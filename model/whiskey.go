package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Whiskey struct {
	gorm.Model
	Name    string         //품명(한국어)
	NameEng string         //품명(English)
	ABV     int64          //도수
	Cask    string         //캐스크
	Origin  string         //원산지
	CID     uint           //위스키 카테고리 No
	Price   uint           //가격 , 우선은 한개..
	CIDs    pq.StringArray `gorm:"type:text[]"`
}

type WCategory struct {
	gorm.Model
	Name string //분류명
}

type WImage struct {
	gorm.Model
	WID       uint   //whiskeyID
	ImageLink string //image link
}
