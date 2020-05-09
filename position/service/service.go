package service

import (
	. "github.com/monchemin/C-19-API/position/model"
	"github.com/monchemin/C-19-API/position/repository"
)

type PositionService interface {
	NewCountry(request CountryRequest) error

	NewTown(request TownRequest) (string, error)

	NewDistrict(request DistrictRequest) (string, error)

	Countries()([]Country, error)

	Localizations()([]Localization, error)
}

type positionService struct {
	repository repository.PositionRepository
}

func NewPositionService(repo repository.PositionRepository) PositionService {
	return &positionService{repository: repo}
}
