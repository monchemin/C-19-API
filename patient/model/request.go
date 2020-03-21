package model

type PatientRequest struct {
	PhoneNumber        string  `json:"phone_number"`
	Age                int     `json:"age"`
	Weight             float64 `json:"weight"`
	IsDiabetic         bool    `json:"is_diabetic"`
	IsHypertensive     bool    `json:"is_hypertensive"`
	IsAsthmatic        bool    `json:"is_asthmatic"`
	IsCardioIschemic   bool    `json:"is_cardio_ischemic"`
	HasLungDisease     bool    `json:"has_lung_disease"`
	HasKidneyDisease   bool    `json:"has_kidney_disease"`
	IsSmoker           bool    `json:"is_smoker"`
	IsReturnFromTravel bool    `json:"is_return_from_travel"`
	Longitude          float64 `json:"longitude"`
	Latitude           float64 `json:"latitude"`
}

type HealthConstantRequest struct {
	PatientID                          string  `json:"patient_id"`
	Temperature                        float64 `json:"temperature"`
	IsTired                            bool    `json:"is_tired"`
	HasDryCough                        bool    `json:"has_dry_cough"`
	HasShortnessOfBreath               bool    `json:"has_shortness_of_breath"`
	HasBeenInContactWithInfectedPerson bool    `json:"has_been_in_contact_with_infected_person"`
	HasHeadache                        bool    `json:"has_headache"`
	HasRunnyNose                       bool    `json:"has_runny_nose"`
	HasNasalCongestion                 bool    `json:"has_nasal_congestion"`
	HasSoreThroat                      bool    `json:"has_sore_throat"`
	HasMusclePain                      bool    `json:"has_muscle_pain"`
	HasDiarrhea                        bool    `json:"has_diarrhea"`
}

func (pr *PatientRequest) IsValid() bool {
	return len(pr.PhoneNumber) > 0 && pr.Age > 0
}

func (h *HealthConstantRequest) IsValid() bool {
	return h.Temperature > 0 && len(h.PatientID) > 0
}

type GetRequest struct {
	PhoneNumber string `json:"phone_number"`
}