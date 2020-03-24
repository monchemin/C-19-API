package service

import (
	"c19/errors"
	. "c19/position/model"
)

func (p positionService) NewCountry(request CountryRequest) error {
	if !request.IsValid() {
		return errors.InvalidRequestData()
	}

	return p.repository.NewCountry(request)
}

func (p positionService) NewTown(request TownRequest) (string, error) {
	if !request.IsValid() {
		return "", errors.InvalidRequestData()
	}

	return p.repository.NewTown(request)
}

func (p positionService) NewDistrict(request DistrictRequest) (string, error) {
	if !request.IsValid() {
		return "", errors.InvalidRequestData()
	}

	return p.repository.NewDistrict(request)
}

func (p positionService) Countries() ([]Country, error) {
	return p.repository.Countries()
}

func (p positionService) Localisations() ([]Localisation, error) {
	return p.repository.Localisations()
}
