package model

import "time"

type Patient struct {
	ID                 string           `json:"ID"`
	PhoneNumber        string           `json:"phone_number"`
	Age                int              `json:"age"`
	Weight             float64          `json:"weight"`
	IsDiabetic         bool             `json:"is_diabetic"`
	IsHypertensive     bool             `json:"is_hypertensive"`
	IsAsthmatic        bool             `json:"is_asthmatic"`
	IsCardioIschemic   bool             `json:"is_cardio_ischemic"`
	HasLungDisease     bool             `json:"has_lung_disease"`
	HasKidneyDisease   bool             `json:"has_kidney_disease"`
	IsSmoker           bool             `json:"is_smoker"`
	IsReturnFromTravel bool             `json:"is_return_from_travel"`
	Longitude          float64          `json:"longitude"`
	Latitude           float64          `json:"latitude"`
	Localization       GeoPoint         `json:"lat_lon"`
	CreatedAt          time.Time        `json:"created_at"`
	DistrictID         string           `json:"district_id"`
	DistrictName       string           `json:"district_name"`
	TownID             string           `json:"town_id"`
	TownName           string           `json:"town_name"`
	CountryCode        string           `json:"country_code"`
	CountryName        string           `json:"country_name"`
	Gender             string           `json:"gender"`
	HealthConstants    []HealthConstant `json:"health_constants"`
}

type HealthConstant struct {
	PatientID            string    `json:"patient_id"`
	DateTime             time.Time `json:"date_time"`
	Temperature          float64   `json:"temperature"`
	IsTired              bool      `json:"is_tired"`
	HasDryCough          bool      `json:"has_dry_cough"`
	HasShortnessOfBreath bool      `json:"has_shortness_of_breath"`
	HasHeadache          bool      `json:"has_headache"`
	HasRunnyNose         bool      `json:"has_runny_nose"`
	HasNasalCongestion   bool      `json:"has_nasal_congestion"`
	HasSoreThroat        bool      `json:"has_sore_throat"`
	HasMusclePain        bool      `json:"has_muscle_pain"`
	HasDiarrhea          bool      `json:"has_diarrhea"`
}

type CreationResponse struct {
	ID string
}

type Login struct {
	ID               string           `json:"id"`
	PhoneNumber      string           `json:"phone_number"`
	DailyInformation []HealthConstant `json:"daily_information"`
}

type GeoPoint struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
