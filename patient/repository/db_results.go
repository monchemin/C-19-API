package repository

import (
	"github.com/google/uuid"
	"time"
)

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
	CountryIsoCode     string    `db:"iso_code"`
	Gender             string    `db:"gender"`
	TownLongitude      float64   `db:"town_longitude"`
	TownLatitude       float64   `db:"town_latitude"`
	IsAtRisk           bool      `db:"is_at_risk"`
	ShouldBeTested     bool      `db:"should_be_tested"`
	IsTested           bool      `db:"is_tested"`
	Height             float64   `db:"height"`
	AtRiskDate         *string   `db:"at_risk_date"`
	InfectedDate       *string   `db:"infected_date"`
	HealingDate        *string   `db:"healing_date"`
	DeathDate          *string   `db:"death_date"`
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

type TestResultResult struct {
	ID            string    `db:"id"`
	PatientID     string    `db:"patient_id"`
	TestCode      string    `db:"test_code"`
	DateTime      time.Time `db:"date_time"`
	IsInfected    bool      `db:"is_infected"`
	IsReinfection bool      `db:"is_reinfection"`
	HealthStatus  string    `db:"health_status"`
}
type RiskStatus struct {
	ID     uuid.UUID
	Status bool
}

type PatientListResult struct {
	PhoneNumber string `db:"phone_number"`
	Location    string `db:"position"`
	Hits        int    `db:"hits"`
	IsAtRisk    bool   `db:"is_at_risk"`
}
