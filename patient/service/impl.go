package service

import (
	"encoding/json"
	"errors"
	"github.com/mmcloughlin/geohash"

	"c19/connector/es"
	"c19/patient/model"
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
	ID, err := ps.repository.AddHealthConstant(request)
	if err != nil {
		return "", err
	}
	patient, _ := ps.Patient(request.PatientID)
	patientJson, err := json.Marshal(patient)
	constantJson, err := json.Marshal(request)
	pData := string(patientJson)
	cData := string(constantJson)
	doc := es.Document{
		ID:    ID,
		Index: "patientConstants",
		Json:  cData[:len(cData)-1] + "," + pData[1:],
	}

	return ID, ps.esClient.Insert(doc, true)
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
		Localization:       LatLonConverter(patientInfo.Latitude, patientInfo.Longitude),
		CreatedAt:          patientInfo.CreatedAt,
		DistrictID:         patientInfo.DistrictID,
		DistrictName:       patientInfo.DistrictName,
		TownID:             patientInfo.TownID,
		TownName:           patientInfo.TownName,
		CountryCode:        patientInfo.CountryCode,
		CountryName:        patientInfo.CountryName,
	}

	return patient, nil
}

func (ps *patientService) PatientHealthConstants(predicate string) (model.Patient, error) {
	patient, err := ps.Patient(predicate)
	if err != nil {
		return model.Patient{}, err
	}
	if patient.ID == "" {
		return model.Patient{}, nil
	}
	constantInfos, err := ps.repository.HealthConstant(patient.ID)

	if err != nil {
		return model.Patient{}, err
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
	patient.HealthConstants = healthConstants
	return patient, nil
}

func (ps *patientService) Connect(phoneNumber string) (model.Login, error) {
	p, err := ps.Patient(phoneNumber)
	if err != nil {
		return model.Login{}, err
	}
	return model.Login{
		ID:          p.ID,
		PhoneNumber: p.PhoneNumber,
	}, nil
}

func (ps *patientService) LatLonConverter (latitude float64, longitude float64) (model.Localization) {
	return model.Localization {
		Lon: longitude,
		Lat: latitude
	}
}
