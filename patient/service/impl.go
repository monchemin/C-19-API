package service

import (
	"c19/patient/model"
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
	return ps.repository.AddHealthConstant(request)
}

func (ps *patientService) Patient(predicate string) (model.Patient, error) {
	if len(predicate) == 0 {
		return model.Patient{}, errors.New("invalid data")
	}

	result, err := ps.repository.Patient(predicate)
	if err != nil {
		return model.Patient{}, err
	}
	patientInfo := result[0]
	constantInfos, err := ps.repository.HealthConstant(patientInfo.ID)

	if err != nil {
		return model.Patient{}, err
	}
	patient := model.Patient{
		ID:          patientInfo.ID,
		PhoneNumber: patientInfo.PhoneNumber,
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
