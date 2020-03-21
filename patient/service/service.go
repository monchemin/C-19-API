package service

import (
	"c19/patient/model"
	"c19/patient/repository"
	"errors"
)

type PatientService interface {
  Add(request model.PatientRequest) (string, error)
  AddHealthConstant(request model.HealthConstantRequest) (string, error)

}

type patientService struct {
	repository repository.PatientRepository
}

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

func NewPatientService (repo repository.PatientRepository) PatientService {
	return &patientService{repository:repo}
}
