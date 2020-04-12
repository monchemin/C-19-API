package repository

import (
	"c19/connector/pgsql"
	"c19/patient/model"
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
}

type repository struct {
	db *pgsql.DB
}

func NewPatientRepository(db *pgsql.DB) PatientRepository {
	return repository{db: db}
}
