package repository

import (
	"c19/patient/model"
	"errors"
	"log"
)

func (r repository) NewPatient(patient model.PatientRequest) (string, error) {
	if !patient.IsValid() {
		return "", errors.New("invalid data")
	}
	row, err := r.db.NamedQuery(insertNewPatient, patient)
	if err != nil {
		return "", err
	}
	var insertID string
	if row.Next() {
		row.Scan(&insertID)
	}
	return insertID, err
}

func (r repository) AddHealthConstant(constant model.HealthConstantRequest) (string, error) {
	if !constant.IsValid() {
		return "", errors.New("invalid data")
	}
	row, err := r.db.NamedQuery(insertNewConstant, constant)
	if err != nil {
		return "", err
	}
	var insertID string
	if row.Next() {
		row.Scan(&insertID)
	}
	return insertID, err
}

func (r repository) Patient(predicate string) ([]PatientResult, error) {
	if len(predicate) == 0 {
		return nil, errors.New("invalid predicate")
	}
	var result []PatientResult
	err := r.db.Select(&result, getPatient, predicate)
	return  result, err
}

func (r repository) HealthConstant(predicate string) ([]HealthConstantResult, error) {
	log.Println(predicate)
	if len(predicate) == 0 {
		return nil, errors.New("invalid predicate")
	}
	var result []HealthConstantResult
	err := r.db.Select(&result, getPatientHealthConstants, predicate)
	return  result, err
}
