package service

import (
	"encoding/json"
	"errors"
	"log"

	"c19/connector/es"
	"c19/patient/model"
	"c19/patient/repository"
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
	cData := string(testResultJson)
	doc := es.Document{
		ID:    ID,
		Index: "patientresults",
		Json:  cData[:len(cData)-1] + "," + pData[1:],
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

func (ps *patientService) IndexConstants() {

	result, err := ps.repository.NotIndexedConstants()
	if err != nil {
		log.Println(err)
		return
	}

	constantMap := make(map[string][]repository.HealthConstantResult)
	for _, hc := range result {
		constantMap[hc.PatientID] = append(constantMap[hc.PatientID], hc)
	}

	keys := make([]string, len(constantMap))
	for key, _ := range constantMap {
		keys = append(keys, key)
	}
	patients, err := ps.repository.InPatient(keys...)
	if err != nil {
		log.Println(err)
		_ = ps.repository.IndexedConstant(false, err.Error())
		return
	}

	for _, rawPatient := range patients {
		patient := ps.patientResultToPatient(rawPatient)
		values := constantMap[rawPatient.ID]
		patientJson, err := json.Marshal(patient)
		if err != nil {
			log.Println(err)
			_ = ps.repository.IndexedConstant(false, err.Error())
			return
		}
		pData := string(patientJson)
		for _, c := range values {
			constantJson, err := json.Marshal(ps.ConstantResultToHealthConstant(c))
			if err != nil {
				log.Println(err)
				_ = ps.repository.IndexedConstant(false, err.Error())
				return
			}
			cData := string(constantJson)
			doc := es.Document{
				ID:    c.ID,
				Index: "patientconstants",
				Json:  cData[:len(cData)-1] + "," + pData[1:],
			}
			err = ps.esClient.Insert(doc, true)
			if err != nil {
				log.Println(err)
				_ = ps.repository.IndexedConstant(false, err.Error())
				return
			}
		}
	}
	err = ps.repository.IndexedConstant(true, "")
	log.Println(err)

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
	case "ACTIVE":
		request.PatientInfected = request.PatientID + "_" + "CONFIRMED"
		request.PatientActive = request.PatientID + "_" + "ACTIVE"
		if (request.IsReinfection) {
			request.PatientReinfected = request.PatientID + "_" + "REINFECTED"
		}
	case "HEALED":
		request.PatientHealed = request.PatientID + "_" + "HEALED"
	case "DEATH":
		request.PatientDead = request.PatientID + "_" + "DEATH"
	}

	return request
}
