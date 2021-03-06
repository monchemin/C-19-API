package repository

import (
	"time"

	"github.com/monchemin/C-19-API/connector/pgsql"
	"github.com/monchemin/C-19-API/patient/model"
)

type PatientRepository interface {
	//Add new patient return insert ID (uuid) as string or error
	NewPatient(patient model.PatientRequest) (string, error)
	//Add new constant return insert ID (uuid) as string or error
	AddHealthConstant(constant model.HealthConstantRequest) (string, string, error)
	//retrieve patient info base on ID or phone number both as string
	Patient(predicate string) ([]PatientResult, error)
	//retrieve patient's health constant info base on ID or phone number both as string
	HealthConstant(predicate string) ([]HealthConstantResult, error)

	Connect(predicate string) ([]PatientResult, error)

	//constant that have not indexed in es service
	NotIndexedConstants() ([]HealthConstantResult, error)
	//get patient list from id list
	InPatient(patientIds ...string) ([]PatientResult, error)

	IndexedConstant(state bool, message string) error

	NewTestResult(testResult model.TestResultRequest) (string, string, error)

	TestResult(predicate string) ([]TestResultResult, error)

	NewToIndex() ([]PatientResult, error)

	NewConstantToIndex(startDate, endDate time.Time) ([]HealthConstantResult, error)

	UpdatePatientStatus(statuses ...RiskStatus) error

	PatientList() ([]PatientListResult, error)
}

type repository struct {
	db *pgsql.DB
}


func NewPatientRepository(db *pgsql.DB) PatientRepository {
	return repository{db: db}
}
