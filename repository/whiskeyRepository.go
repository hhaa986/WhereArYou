package repository

import (
	"gorm.io/gorm"
	"wayBack/db"
	"wayBack/model"
)

type IWhiskeyRepository interface {
	CreateWhiskeyDb(whiskey model.Whiskey) model.Whiskey
}

type WhiskeyRepository struct {
	DbCon *db.Connection
}

func NewWhiskeyRepository(dbCon *db.Connection) *WhiskeyRepository {
	dbCon.Db.AutoMigrate(&model.WCategory{}, &model.Whiskey{}, &model.WReview{})
	return &WhiskeyRepository{DbCon: dbCon}
}

// CreateWhiskeyDb Whiskey
func (wd *WhiskeyRepository) CreateWhiskeyDb(whiskey model.Whiskey) model.Whiskey {
	result := wd.DbCon.Db.Save(&whiskey)
	if result.Error != nil {
		panic("Errrr fail to create whiskey")
	}
	return whiskey
}

func (wd *WhiskeyRepository) GetAllWhiskeysDb() []model.Whiskey {
	var whiskeys []model.Whiskey
	wd.DbCon.Db.Find(&whiskeys)
	return whiskeys
}

func (wd *WhiskeyRepository) GetWhiskeyDb(id uint) model.Whiskey {
	var whiskey = model.Whiskey{Model: gorm.Model{ID: id}}
	wd.DbCon.Db.Find(&whiskey)
	return whiskey
}

func (wd *WhiskeyRepository) CreateWhiskeyCDb(whiskey model.WCategory) model.WCategory {
	result := wd.DbCon.Db.Save(&whiskey)
	if result.Error != nil {
		panic("Errrr fail to create whiskey")
	}
	return whiskey
}
