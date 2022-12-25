package service

import (
	"wayBack/db"
	"wayBack/model"
	repo "wayBack/repository"
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

// GetAllWhiskey WhiskeyService godoc
// @Accept json
// @Produce json
func (s *WhiskeyService) GetAllWhiskey() []model.Whiskey {
	whiskeys := s.Repository.GetAllWhiskeysDb()
	return whiskeys
}

// GetWhiskey WhiskeyService godoc
// @Accept json
// @Produce json
// @Success 200 {} model.Whiskey
// @Router /whiskey [get]
func (s *WhiskeyService) GetWhiskey(name string) model.Whiskey {
	var a model.Whiskey
	a.Name = name
	return a
}

func (s *WhiskeyService) CreateWhiskey(cw model.CreateWRequest) model.Whiskey {
	w := model.Whiskey{
		Name:         cw.WName,
		AlcoholLevel: cw.AlcoholLevel,
		Origin:       cw.Origin,
		CID:          cw.CID,
	}
	//Todo cid가 category에 없는거면 반환 에러임!!
	whiskey := s.Repository.CreateWhiskeyDb(w)
	return whiskey
}
