package service

import (
	"c19/patient/model"
	"c19/patient/repository"
)

type PatientService interface {
	//add new patient
  Add(request model.PatientRequest) (string, error)
  // add constant
  AddHealthConstant(request model.HealthConstantRequest) (string, error)
  //retrieve patient info and constant since date
  Patient(predicate string) (model.Patient, error)
}

type patientService struct {
	repository repository.PatientRepository
}

func NewPatientService (repo repository.PatientRepository) PatientService {
	return &patientService{repository:repo}
}
