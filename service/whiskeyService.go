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

func (s *WhiskeyService) GetAllWhiskey() []model.Whiskey {
	whiskeys := s.Repository.GetAllWhiskeysDb()
	return whiskeys
}
