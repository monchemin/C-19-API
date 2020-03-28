package repository

import (
	"c19/errors"
	"c19/position/model"
)

func (r repository) NewCountry(request model.CountryRequest) error {
	if !request.IsValid() {
		return errors.InvalidRequestData()
	}
	_, err := r.db.NamedQuery(insertNewCountry, request)
	return err
}

func (r repository) NewTown(request model.TownRequest) (string, error) {
	if !request.IsValid() {
		return "", errors.InvalidRequestData()
	}
	row, err := r.db.NamedQuery(insertNewTown, request)
	if err != nil {
		return "", err
	}
	var ID string
	if row.Next() {
		row.Scan(&ID)
	}
	return ID, err
}

func (r repository) NewDistrict(request model.DistrictRequest) (string, error) {
	if !request.IsValid() {
		return "", errors.InvalidRequestData()
	}
	row, err := r.db.NamedQuery(insertNewDistrict, request)
	if err != nil {
		return "", err
	}
	var ID string
	if row.Next() {
		row.Scan(&ID)
	}
	return ID, err
}

func (r repository) Countries() ([]model.Country, error) {
	var result []model.Country
	err := r.db.Select(&result, getCountries)
	return result, err
}

func (r repository) Localizations() ([]model.Localization, error) {
	var result []model.Localization
	err := r.db.Select(&result, getLocalizations)
	return result, err
}
