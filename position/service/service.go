package service

import (
	"c19/position/model"
	"c19/position/repository"
)

type PositionService interface {
	NewCountry(request model.CountryRequest) error

	NewTown(request model.TownRequest) (string, error)

	NewDistrict(request model.DistrictRequest) (string, error)
}

type positionService struct {
	repository repository.PositionRepository
}

func NewPositionService(repo repository.PositionRepository) PositionService {
	return &positionService{repository: repo}
}
