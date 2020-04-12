package repository

import (
	"c19/patient/model"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
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

func (r repository) AddHealthConstant(constant model.HealthConstantRequest) (string, string, error) {
	if !constant.IsValid() {
		return "", "", errors.New("invalid data")
	}
	row, err := r.db.NamedQuery(insertNewConstant, constant)
	if err != nil {
		return "", "", err
	}
	var insertID, dt string
	if row.Next() {
		row.Scan(&insertID, &dt)
	}
	return insertID, dt, err
}

func (r repository) Patient(predicate string) ([]PatientResult, error) {
	if len(predicate) == 0 {
		return nil, errors.New("invalid predicate")
	}
	var result []PatientResult

	err := r.db.Select(&result, getPatient, predicate)
	return result, err
}

func (r repository) Connect(predicate string) ([]PatientResult, error) {
	if len(predicate) == 0 {
		return nil, errors.New("invalid predicate")
	}
	var result []PatientResult

	err := r.db.Select(&result, getPatientConnection, predicate)
	return result, err
}

func (r repository) HealthConstant(predicate string) ([]HealthConstantResult, error) {
	if len(predicate) == 0 {
		return nil, errors.New("invalid predicate")
	}
	var result []HealthConstantResult
	err := r.db.Select(&result, getPatientHealthConstants, predicate)
	return result, err
}

func (r repository) NotIndexedConstants() ([]HealthConstantResult, error) {
	var result []HealthConstantResult
	err := r.db.Select(&result, notIndexedConstants)
	return result, err
}

func (r repository) InPatient(patientIds ...string) ([]PatientResult, error) {
	if len(patientIds) == 0 {
		return nil, errors.New("invalid list")
	}
	var result []PatientResult
	query, args, err := sqlx.In(InPatient, patientIds)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	err = r.db.Select(&result, query, args...)
	return result, err
}

func (r repository) IndexedConstant(state bool, message string) error {
	_, err := r.db.Query(InsertIndexedDate, state, message)
	return err
}

func (r repository) NewTestResult(testResult model.TestResultRequest) (string, string, error) {
	if !testResult.IsValid() {
		return "", "", errors.New("invalid data")
	}
	row, err := r.db.NamedQuery(insertNewTestResult, testResult)
	if err != nil {
		return "", "", err
	}
	var insertID, dt string
	if row.Next() {
		row.Scan(&insertID, &dt)
	}
	return insertID, dt, err
}

func (r repository) TestResult(predicate string) ([]TestResultResult, error) {
	if len(predicate) == 0 {
		return nil, errors.New("invalid predicate")
	}
	var result []TestResultResult

	err := r.db.Select(&result, getPatientTestResult, predicate)
	return result, err
}

func (r repository) NewToIndex() ([]string, error) {
	var result []HealthConstantResult
	err := r.db.Select(&result, newToIndex)
	response := make([]string, len(result))
	for i, row := range result {
		response[i] = row.PatientID
	}
	return response, err
}

func (r repository) NewConstantToIndex(startDate, endDate time.Time) ([]HealthConstantResult, error) {
	var result []HealthConstantResult
	log.Println(startDate, endDate)
	err := r.db.Select(&result, newConstantsToIndex, startDate, endDate)
	return result, err
}

func (r repository) UpdatePatientStatus(statuses ...RiskStatus)  error {
	if len(statuses) == 0 {
		return errors.New("invalid list")
	}
	for _, status := range statuses {
		_, err := r.db.NamedQuery(patientUpdate, status)
		if err != nil {
			return err
		}
	}
	return  nil
}
