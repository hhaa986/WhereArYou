package service

import (
	"way_BE/db"
	"way_BE/model"
	repo "way_BE/repository"
)

type WhiskeyService struct {
	Repository *repo.WhiskeyRepository
}

var Service WhiskeyService

func (s *WhiskeyService) InitService() error {
	dbCon := db.GetDbConnection()
	s.Repository = repo.NewWhiskeyRepository(dbCon)
	return nil
}

// Whiskey
func (s *WhiskeyService) GetAllWhiskey() []model.Whiskey {
	return s.Repository.GetAllWhiskeysDb()
}

func (s *WhiskeyService) GetWhiskeyById(id uint) model.Whiskey {
	return s.Repository.GetWhiskeyDbByID(id)
}

func (s *WhiskeyService) GetWhiskeyByName(name string) []model.Whiskey {
	return s.Repository.GetWhiskeyDbByName(name)
}

func (s *WhiskeyService) GetWhiskeyByCategory(cid uint) []model.Whiskey {
	return s.Repository.GetWhiskeyDbByCid(cid)
}

func (s *WhiskeyService) CreateWhiskey(w model.Whiskey) model.Whiskey {
	//Todo cid가 category에 없는거면 반환 에러임!!
	return s.Repository.CreateWhiskeyDb(model.Whiskey{
		Name:   w.Name,
		ABV:    w.ABV,
		Origin: w.Origin,
		CID:    w.CID,
	})
}

func (s *WhiskeyService) UpdateWhiskey(whiskey model.Whiskey) model.Whiskey {
	return s.Repository.UpdateWhiskeyDb(whiskey)
}

func (s *WhiskeyService) DeleteWhiskey(id uint) error {
	return s.Repository.DeleteWhiskeyDb(id)
}

// Category
func (s *WhiskeyService) GetAllWCategory() []model.WCategory {
	return s.Repository.GetAllWCategoryDb()
}

func (s *WhiskeyService) GetWCategory(id uint) model.WCategory {
	return s.Repository.GetWCategoryDbById(id)
}

func (s *WhiskeyService) GetWCategoryByName(name string) []model.WCategory {
	return s.Repository.GetWCategoryDbByName(name)
}

func (s *WhiskeyService) CreateWCategory(w model.WCategory) model.WCategory {
	return s.Repository.CreateWCategoryDb(model.WCategory{
		Name: w.Name,
	})
}

func (s *WhiskeyService) UpdateWCategory(whiskey model.WCategory) model.WCategory {
	return s.Repository.UpdateWCategoryDb(whiskey)
}

func (s *WhiskeyService) DeleteWCategory(id uint) error {
	return s.Repository.DeleteWCategoryDb(id)
}
