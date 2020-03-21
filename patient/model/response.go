package model

import "time"

type Patient struct {
	ID              string           `json:"ID"`
	PhoneNumber     string           `json:"phone_number"`
	HealthConstants []HealthConstant `json:"health_constants"`
}


type HealthConstant struct {
	PatientID            string  `json:"patient_id"`
	DateTime             time.Time  `json:"date_time"`
	Temperature          float64 `json:"temperature"`
	IsTired              bool    `json:"is_tired"`
	HasDryCough          bool    `json:"has_dry_cough"`
	HasShortnessOfBreath bool    `json:"has_shortness_of_breath"`
	HasHeadache          bool    `json:"has_headache"`
	HasRunnyNose         bool    `json:"has_runny_nose"`
	HasNasalCongestion   bool    `json:"has_nasal_congestion"`
	HasSoreThroat        bool    `json:"has_sore_throat"`
	HasMusclePain        bool    `json:"has_muscle_pain"`
	HasDiarrhea          bool    `json:"has_diarrhea"`
}

type CreationResponse struct {
	ID string
}
