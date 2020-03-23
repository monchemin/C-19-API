package service

import (
	"c19/errors"
	"c19/position/model"
)

func (p positionService) NewCountry(request model.CountryRequest) error {
	if !request.IsValid() {
		return errors.InvalidRequestData()
	}

	return p.repository.NewCountry(request)
}

func (p positionService) NewTown(request model.TownRequest) (string, error) {
	if !request.IsValid() {
		return "", errors.InvalidRequestData()
	}

	return p.repository.NewTown(request)
}

func (p positionService) NewDistrict(request model.DistrictRequest) (string, error) {
	if !request.IsValid() {
		return "", errors.InvalidRequestData()
	}

	return p.repository.NewDistrict(request)
}
