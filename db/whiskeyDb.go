package db

import (
	"wayBack/model"
)

type IWhiskeyDb interface {
	CreateWhiskeyDb(whiskey model.Whiskey) model.Whiskey
}

type WhiskeyDb struct {
	DbCon *Connection
}

func NewWhiskeyDb(dbCon *Connection) *WhiskeyDb {
	dbCon.Db.AutoMigrate(&model.Whiskey{}, &model.WCategory{}, &model.WReview{})
	return &WhiskeyDb{DbCon: dbCon}
}

// create
func (wd *WhiskeyDb) CreateWhiskeyDb(whiskey model.Whiskey) model.Whiskey {
	result := wd.DbCon.Db.Save(&whiskey)
	if result.Error != nil {
		panic("Errrr fail to create whiskey")
	}
	return whiskey
}
