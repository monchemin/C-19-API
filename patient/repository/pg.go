package repository

import (
	"c19/patient/model"
	"errors"
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

func (r repository) Patient(predicate string) (PatientResult, error) {
	panic("implement me")
}

func (r repository) HealthConstant(predicate string) (HealthConstantResult, error) {
	panic("implement me")
}
