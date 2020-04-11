package service

import (
	"c19/connector/es"
	"c19/patient/model"
	"encoding/json"
	"errors"
)

func (ps *patientService) Add(request model.PatientRequest) (string, error) {
	if !request.IsValid() {
		return "", errors.New("invalid request data")
	}
	return ps.repository.NewPatient(request)
}

func (ps *patientService) AddHealthConstant(request model.HealthConstantRequest) (string, error) {
	if !request.IsValid() {
		return "", errors.New("invalid request data")
	}
	ID, DT, err := ps.repository.AddHealthConstant(request)
	if err != nil {
		return "", err
	}
	request.DateTime = DT
	patient, _ := ps.Patient(request.PatientID)
	patientJson, err := json.Marshal(patient)
	constantJson, err := json.Marshal(request)
	pData := string(patientJson)
	cData := string(constantJson)
	doc := es.Document{
		ID:    ID,
		Index: "patientconstants",
		Json:  cData[:len(cData)-1] + "," + pData[1:],
	}
	_ = ps.esClient.Insert(doc, true)
	return ID, nil
}

func (ps *patientService) Patient(predicate string) (model.Patient, error) {
	if len(predicate) == 0 {
		return model.Patient{}, errors.New("invalid data")
	}
	result, err := ps.repository.Patient(predicate)
	if err != nil {
		return model.Patient{}, err
	}
	if result == nil {
		return model.Patient{}, nil
	}
	patientInfo := result[0]
	patient := model.Patient{
		ID:                 patientInfo.ID,
		PhoneNumber:        patientInfo.PhoneNumber,
		Age:                patientInfo.Age,
		Weight:             patientInfo.Weight,
		IsDiabetic:         patientInfo.IsDiabetic,
		IsHypertensive:     patientInfo.IsHypertensive,
		IsAsthmatic:        patientInfo.IsAsthmatic,
		IsCardioIschemic:   patientInfo.IsCardioIschemic,
		HasLungDisease:     patientInfo.HasLungDisease,
		HasKidneyDisease:   patientInfo.HasKidneyDisease,
		IsSmoker:           patientInfo.IsSmoker,
		IsReturnFromTravel: patientInfo.IsReturnFromTravel,
		Longitude:          patientInfo.Longitude,
		Latitude:           patientInfo.Latitude,
		Localization:       ps.geoPointConverter(patientInfo.Latitude, patientInfo.Longitude),
		CreatedAt:          patientInfo.CreatedAt,
		DistrictID:         patientInfo.DistrictID,
		DistrictName:       patientInfo.DistrictName,
		TownID:             patientInfo.TownID,
		TownName:           patientInfo.TownName,
		CountryCode:        patientInfo.CountryCode,
		CountryName:        patientInfo.CountryName,
		CountryIsoCode:     patientInfo.CountryIsoCode,
		Gender:             patientInfo.Gender,
		TownLocalization:   ps.geoPointConverter(patientInfo.TownLatitude, patientInfo.TownLongitude),
		IsAtRisk:           patientInfo.IsAtRisk,
		ShouldBeTested:     patientInfo.ShouldBeTested,
		IsTested:           patientInfo.IsTested,
		Height:             patientInfo.Height,
		AtRiskDate:         patientInfo.AtRiskDate,
		InfectedDate:       patientInfo.InfectedDate,
		HealingDate:        patientInfo.HealingDate,
		DeathDate:          patientInfo.DeathDate,
	}

	return patient, nil
}

func (ps *patientService) PatientHealthConstants(patientID string) ([]model.HealthConstant, error) {

	constantInfos, err := ps.repository.HealthConstant(patientID)

	if err != nil {
		return nil, err
	}
	healthConstants := make([]model.HealthConstant, len(constantInfos))
	for index, info := range constantInfos {
		healthConstants[index] = model.HealthConstant{
			PatientID:            info.PatientID,
			DateTime:             info.DateTime,
			Temperature:          info.Temperature,
			IsTired:              info.IsTired,
			HasDryCough:          info.HasDryCough,
			HasShortnessOfBreath: info.HasShortnessOfBreath,
			HasHeadache:          info.HasHeadache,
			HasRunnyNose:         info.HasRunnyNose,
			HasNasalCongestion:   info.HasNasalCongestion,
			HasSoreThroat:        info.HasSoreThroat,
			HasMusclePain:        info.HasMusclePain,
			HasDiarrhea:          info.HasDiarrhea,
		}
	}

	return healthConstants, nil
}

func (ps *patientService) Connect(phoneNumber string) (model.Login, error) {
	result, err := ps.repository.Connect(phoneNumber)
	if err != nil {
		return model.Login{}, err
	}
	if result == nil {
		return model.Login{}, err
	}
	connection := result[0]
	hc, _ := ps.PatientHealthConstants(connection.ID)
	return model.Login{
		ID:               connection.ID,
		PhoneNumber:      connection.PhoneNumber,
		DailyInformation: hc,
	}, nil
}

func (ps *patientService) geoPointConverter(latitude float64, longitude float64) model.GeoPoint {
	return model.GeoPoint{
		Lon: longitude,
		Lat: latitude,
	}
}
