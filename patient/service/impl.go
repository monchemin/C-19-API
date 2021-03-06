package service

import (

	"encoding/json"
	"errors"
	"log"

	"github.com/monchemin/C-19-API/connector/es"
	"github.com/monchemin/C-19-API/patient/model"
	"github.com/monchemin/C-19-API/patient/repository"
)

const indexationDelta = 72

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
		log.Println(err)
		return model.Patient{}, err
	}
	if result == nil {
		log.Println("no patient")
		return model.Patient{}, nil
	}
	return ps.patientResultToPatient(result[0]), nil
}

func (ps *patientService) PatientHealthConstants(patientID string) ([]model.HealthConstant, error) {

	constantInfos, err := ps.repository.HealthConstant(patientID)

	if err != nil {
		return nil, err
	}
	healthConstants := make([]model.HealthConstant, len(constantInfos))
	for index, info := range constantInfos {
		healthConstants[index] = ps.ConstantResultToHealthConstant(info)
	}

	return healthConstants, nil
}

func (ps *patientService) NewTestResult(request model.TestResultRequest) (string, error) {
	if !request.IsValid() {
		return "", errors.New("invalid request data")
	}
	ID, DT, err := ps.repository.NewTestResult(request)
	if err != nil {
		return "", err
	}
	request.DateTime = DT
	augmentedRequest := ps.augmentTestResultRequestForEs(request)
	patient, _ := ps.Patient(request.PatientID)
	patientJson, err := json.Marshal(patient)
	testResultJson, err := json.Marshal(augmentedRequest)
	pData := string(patientJson)
	trData := string(testResultJson)
	doc := es.Document{
		ID:    ID,
		Index: "patientresults",
		Json:  trData[:len(trData)-1] + "," + pData[1:],
	}
	_ = ps.esClient.Insert(doc, true)
	return ID, nil
}

func (ps *patientService) PatientTestResult(patientID string) ([]model.TestResult, error) {

	testResultInfos, err := ps.repository.TestResult(patientID)

	if err != nil {
		return nil, err
	}
	testResults := make([]model.TestResult, len(testResultInfos))
	for index, info := range testResultInfos {
		testResults[index] = ps.TestResultResultToTestResult(info)
	}

	return testResults, nil
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

func (ps *patientService) patientResultToPatient(rawPatient repository.PatientResult) model.Patient {
	return model.Patient{
		ID:                 rawPatient.ID,
		PhoneNumber:        rawPatient.PhoneNumber,
		Age:                rawPatient.Age,
		Weight:             rawPatient.Weight,
		IsDiabetic:         rawPatient.IsDiabetic,
		IsHypertensive:     rawPatient.IsHypertensive,
		IsAsthmatic:        rawPatient.IsAsthmatic,
		IsCardioIschemic:   rawPatient.IsCardioIschemic,
		HasLungDisease:     rawPatient.HasLungDisease,
		HasKidneyDisease:   rawPatient.HasKidneyDisease,
		IsSmoker:           rawPatient.IsSmoker,
		IsReturnFromTravel: rawPatient.IsReturnFromTravel,
		Longitude:          rawPatient.Longitude,
		Latitude:           rawPatient.Latitude,
		Localization:       ps.geoPointConverter(rawPatient.Latitude, rawPatient.Longitude),
		CreatedAt:          rawPatient.CreatedAt,
		DistrictID:         rawPatient.DistrictID,
		DistrictName:       rawPatient.DistrictName,
		TownID:             rawPatient.TownID,
		TownName:           rawPatient.TownName,
		CountryCode:        rawPatient.CountryCode,
		CountryName:        rawPatient.CountryName,
		CountryIsoCode:     rawPatient.CountryIsoCode,
		Gender:             rawPatient.Gender,
		TownLocalization:   ps.geoPointConverter(rawPatient.TownLatitude, rawPatient.TownLongitude),
		IsAtRisk:           rawPatient.IsAtRisk,
		ShouldBeTested:     rawPatient.ShouldBeTested,
		IsTested:           rawPatient.IsTested,
		Height:             rawPatient.Height,
		AtRiskDate:         rawPatient.AtRiskDate,
		InfectedDate:       rawPatient.InfectedDate,
		HealingDate:        rawPatient.HealingDate,
		DeathDate:          rawPatient.DeathDate,
	}
}

func (ps *patientService) ConstantResultToHealthConstant(info repository.HealthConstantResult) model.HealthConstant {

	return model.HealthConstant{
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

func (ps *patientService) TestResultResultToTestResult(info repository.TestResultResult) model.TestResult {

	return model.TestResult{
		PatientID:            info.PatientID,
		TestCode:             info.TestCode,
		DateTime:             info.DateTime,
		IsInfected:           info.IsInfected,
		IsReinfection:        info.IsReinfection,
		HealthStatus:         info.HealthStatus,
	}
}

func (ps *patientService) augmentTestResultRequestForEs(request model.TestResultRequest) model.TestResultRequest {

	switch request.HealthStatus {
	case "SAFE":
		safe := request.PatientID + "_" + "SAFE"
		request.PatientSafe = &safe
	case "ACTIVE":
		confirmed := request.PatientID + "_" + "CONFIRMED"
		actif := request.PatientID + "_" + "ACTIVE"
		reinfected := request.PatientID + "_" + "REINFECTED"
		request.PatientInfected = &confirmed
		request.PatientActive = &actif
		if (request.IsReinfection) {
			request.PatientReinfected = &reinfected
		}
	case "HEALED":
		healed := request.PatientID + "_" + "HEALED"
		request.PatientHealed = &healed
	case "DEATH":
		death := request.PatientID + "_" + "DEATH"
		request.PatientDead = &death
	}

	return request
}

func (ps *patientService) PatientList() ([]model.ShortPatient, error) {
	result, err := ps.repository.PatientList()
	if err != nil {
		return nil, err
	}
	response := make([]model.ShortPatient, len(result))
	for i, p := range result{
		response[i] = model.ShortPatient{
			PhoneNumber: p.PhoneNumber,
			Location:    p.Location,
			Hits:        p.Hits,
			IsAtRisk:    p.IsAtRisk,
		}
	}

	return response, nil
}
