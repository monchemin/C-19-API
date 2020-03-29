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
	countryInfos, err := p.repository.Countries()
	if (err != nil) {
		return nil, err
	} else {
		var countries[]Country
		if len(countryInfos) == 0 {
			return nil, errors.EmptyResultData()
		} else {
			for index, info := range countryInfos {
				countries[index] = Country{
					ID:            info.ID,
					Name:          info.Name,
					IsoCode:       info.IsoCode,
				}
			}
			return countries, err// p.repository.Countries()
		}
	}
}

func (p positionService) Localizations() ([]Localization, error) {
	return p.repository.Localizations()
}
