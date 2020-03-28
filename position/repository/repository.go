package repository

import (
	"c19/connector/pgsql"
	"c19/position/model"
)

type PositionRepository interface {
	NewCountry(request model.CountryRequest) error

	NewTown(request model.TownRequest) (string, error)

	NewDistrict(request model.DistrictRequest) (string, error)

	Countries() ([]model.Country, error)

	Localizations() ([]model.Localization, error)
}

type repository struct {
	db *pgsql.DB
}

func NewPositionRepository(db *pgsql.DB) PositionRepository {
	return repository{db: db}
}
