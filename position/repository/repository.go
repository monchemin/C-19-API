package repository

import (
	"github.com/monchemin/C-19-API/connector/pgsql"
	"github.com/monchemin/C-19-API/position/model"
)

type PositionRepository interface {
	NewCountry(request model.CountryRequest) error

	NewTown(request model.TownRequest) (string, error)

	NewDistrict(request model.DistrictRequest) (string, error)

	Countries() ([]CountryResult, error)

	Localizations() ([]model.Localization, error)
}

type repository struct {
	db *pgsql.DB
}

func NewPositionRepository(db *pgsql.DB) PositionRepository {
	return repository{db: db}
}
