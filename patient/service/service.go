package service

import (
	"c19/connector/es"
	"c19/patient/model"
	"c19/patient/repository"
)

type PatientService interface {
	//add new patient
	Add(request model.PatientRequest) (string, error)
	// add constant
	AddHealthConstant(request model.HealthConstantRequest) (string, error)
	//retrieve patient info
	Patient(predicate string) (model.Patient, error)
	//retrieve patient info and constant
	PatientHealthConstants(predicate string) ([]model.HealthConstant, error)
	// connect exist patient
	Connect(phoneNumber string)(model.Login, error)

	IndexConstants()
	//add new result
	NewTestResult(request model.TestResultRequest) (string, error)
	//retrieve patient info and test result
	PatientTestResult(predicate string) ([]model.TestResult, error)

	IndexPatients()

}

type patientService struct {
	repository repository.PatientRepository
	esClient   es.ElasticSearchClient
}

func NewPatientService(repo repository.PatientRepository, esClient es.ElasticSearchClient) PatientService {
	return &patientService{repository: repo, esClient: esClient}
}
