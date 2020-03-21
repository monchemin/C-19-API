package repository

import "time"

type PatientResult struct {
	ID string
	PhoneNumber        string  `json:"phone_number"`
	Age                int     `json:"age"`
	Weight             float64 `json:"weight"`
}

type HealthConstantResult struct {
	ID                                 string    `db:"id"`
	DateTime                           time.Time `db:"patient_id"`
	PatientID                          string    `db:"date_time"`
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
