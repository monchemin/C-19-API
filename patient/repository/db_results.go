package repository

import "time"

type PatientResult struct {
	ID                 string    `db:"id"`
	PhoneNumber        string    `db:"phone_number"`
	Age                int       `db:"age"`
	Weight             float64   `db:"weight"`
	IsDiabetic         bool      `db:"is_diabetic"`
	IsHypertensive     bool      `db:"is_hypertensive"`
	IsAsthmatic        bool      `db:"is_asthmatic"`
	IsCardioIschemic   bool      `db:"is_cardio_ischemic"`
	HasLungDisease     bool      `db:"has_lung_disease"`
	HasKidneyDisease   bool      `db:"has_kidney_disease"`
	IsSmoker           bool      `db:"is_smoker"`
	IsReturnFromTravel bool      `db:"is_return_from_travel"`
	Longitude          float64   `db:"longitude"`
	Latitude           float64   `db:"latitude"`
	CreatedAt          time.Time `db:"created_at"`
	DistrictID         string    `db:"district_id"`
	DistrictName       string    `db:"district_name"`
	TownID             string    `db:"town_id"`
	TownName           string    `db:"town_name"`
	CountryCode        string    `db:"country_code"`
	CountryName        string    `db:"country_name"`
	Gender             string    `db:"gender"`
}

type HealthConstantResult struct {
	ID                                 string    `db:"id"`
	DateTime                           time.Time `db:"date_time"`
	PatientID                          string    `db:"patient_id"`
	Temperature                        float64   `db:"temperature"`
	IsTired                            bool      `db:"is_tired"`
	HasDryCough                        bool      `db:"has_dry_cough"`
	HasShortnessOfBreath               bool      `db:"has_shortness_of_breath"`
	HasBeenInContactWithInfectedPerson bool      `db:"has_been_in_contact_with_infected_person"`
	HasHeadache                        bool      `db:"has_headache"`
	HasRunnyNose                       bool      `db:"has_runny_nose"`
	HasNasalCongestion                 bool      `db:"has_nasal_congestion"`
	HasSoreThroat                      bool      `db:"has_sore_throat"`
	HasMusclePain                      bool      `db:"has_muscle_pain"`
	HasDiarrhea                        bool      `db:"has_diarrhea"`
}
