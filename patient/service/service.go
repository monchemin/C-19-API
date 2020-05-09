package service

import (
	"github.com/monchemin/C-19-API/connector/es"
	"github.com/monchemin/C-19-API/patient/model"
	"github.com/monchemin/C-19-API/patient/repository"
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

	PatientList() ([]model.ShortPatient, error)

}

type patientService struct {
	repository repository.PatientRepository
	esClient   es.ElasticSearchClient
}

func NewPatientService(repo repository.PatientRepository, esClient es.ElasticSearchClient) PatientService {
	return &patientService{repository: repo, esClient: esClient}
}
