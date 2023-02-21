package repository

import (
	"way_BE/db"
	"way_BE/model"
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
	result := wd.DbCon.Db.Create(&whiskey)
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

func (wd *WhiskeyRepository) GetWhiskeyDbByID(id uint) model.Whiskey {
	var whiskey model.Whiskey
	wd.DbCon.Db.Where("id = ?", id).Find(&whiskey)
	return whiskey
}

func (wd *WhiskeyRepository) GetWhiskeyDbByName(name string) []model.Whiskey {
	var whiskeys []model.Whiskey
	wd.DbCon.Db.Where("name = ?", name).Find(&whiskeys)
	return whiskeys
}

func (wd *WhiskeyRepository) GetWhiskeyDbByCid(cid uint) []model.Whiskey {
	var whiskeys []model.Whiskey
	wd.DbCon.Db.Where("c_id = ?", cid).Find(&whiskeys)
	return whiskeys
}

func (wd *WhiskeyRepository) UpdateWhiskeyDb(w model.Whiskey) model.Whiskey {
	var wsk map[string]interface{}
	wsk = map[string]interface{}{
		"Name":         w.Name,
		"Origin":       w.Origin,
		"AlcoholLevel": w.AlcoholLevel,
		"CID":          w.CID,
	}
	wd.DbCon.Db.Model(model.Whiskey{}).Where("ID = ?", w.Model.ID).Updates(&wsk)
	wskk := wd.GetWhiskeyDbByID(w.ID)
	return wskk
}

func (wd *WhiskeyRepository) DeleteWhiskeyDb(id uint) error {
	var whiskey model.Whiskey
	wd.DbCon.Db.Delete(&whiskey, id)
	return nil
}

func (wd *WhiskeyRepository) CreateWCategoryDb(whiskey model.WCategory) model.WCategory {
	result := wd.DbCon.Db.Create(&whiskey)
	if result.Error != nil {
		panic("Errrr fail to create whiskey")
	}
	return whiskey
}

func (wd *WhiskeyRepository) GetWCategoryDbById(id uint) model.WCategory {
	var wc model.WCategory
	wd.DbCon.Db.Where("id = ?", id).Find(&wc)
	return wc
}
func (wd *WhiskeyRepository) GetWCategoryDbByName(name string) []model.WCategory {
	var wcs []model.WCategory
	wd.DbCon.Db.Where("name = ?", name).Find(&wcs)
	return wcs
}

func (wd *WhiskeyRepository) GetAllWCategoryDb() []model.WCategory {
	var wCategorys []model.WCategory
	wd.DbCon.Db.Find(&wCategorys)
	return wCategorys
}

func (wd *WhiskeyRepository) UpdateWCategoryDb(w model.WCategory) model.WCategory {
	var wsk map[string]interface{}
	wsk = map[string]interface{}{
		"Name": w.Name,
	}
	wd.DbCon.Db.Model(model.WCategory{}).Where("ID = ?", w.Model.ID).Updates(&wsk)
	wcc := wd.GetWCategoryDbById(w.ID)
	return wcc
}

func (wd *WhiskeyRepository) DeleteWCategoryDb(id uint) error {
	var wc model.WCategory
	wd.DbCon.Db.Delete(&wc, id)
	return nil
}
