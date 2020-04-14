package service

import (
	"encoding/json"
	"log"
	"time"

	"c19/connector/es"
	"c19/patient/repository"

	"github.com/google/uuid"
)

const (
	temperature       = 8
	Tired             = 3
	DryCough          = 6
	shortnessofbreath = 5
	Headache          = 1
	nasalcongestion   = 1
	musclepain        = 2
	diarrhea          = 1
	returnfromtravel  = 2
	diabetic          = 1
	hypertensive      = 1
	Asthmatic         = 3
	smoke             = 2
	CardioIschemic    = 3
	lungdisease       = 3
	kidneydisease     = 1
	threshold         = 20
)

func (ps *patientService) IndexPatients() {

	result, err := ps.repository.NewToIndex()
	if err != nil {
		log.Println(err)
		return
	}
	if len(result) == 0 {
		return
	}

	patients := make(map[string]repository.PatientResult)
	for _, row := range result {
		patients[row.ID] = row
	}
	endDate := time.Now().UTC()
	startDate := endDate.Add(time.Duration(-indexationDelta) * time.Hour)

	constants, err := ps.repository.NewConstantToIndex(startDate, endDate)
	if err != nil {
		log.Println(err)
		return
	}
	constantsMap := make(map[string][]repository.HealthConstantResult)
	for _, hc := range constants {
		constantsMap[hc.PatientID] = append(constantsMap[hc.PatientID], hc)
	}
	patientStatus := make([]repository.RiskStatus, len(patients))
	patientIds := make([]string, len(patients))
	index := 0
	for id, hcs := range constantsMap {
		isAtRisk := ps.riskStatus(patients[id], hcs)
		patientStatus[index] = repository.RiskStatus{ID: uuid.MustParse(id), Status: isAtRisk}
		patientIds[index] = id
		index++
	}

	err = ps.repository.UpdatePatientStatus(patientStatus...)
	if err != nil {
		log.Println(err)
		return
	}
	updatedPatients, err := ps.repository.InPatient(patientIds...)
	for _, rawPatient := range updatedPatients {
		patient := ps.patientResultToPatient(rawPatient)
		patientJson, err := json.Marshal(patient)
		if err != nil {
			log.Println(err)
			_ = ps.repository.IndexedConstant(false, err.Error())
			return
		}
		pData := string(patientJson)

		doc := es.Document{
			ID:    patient.ID,
			Index: "patients",
			Json:  pData,
		}
		err = ps.esClient.Insert(doc, true)
		if err != nil {
			log.Println(err)
			_ = ps.repository.IndexedConstant(false, err.Error())
			return
		}
	}

	err = ps.repository.IndexedConstant(true, "")
	log.Println(err)

}

func (ps *patientService) riskStatus(p repository.PatientResult, hcs []repository.HealthConstantResult) bool {
	score := 0
	for _, hc := range hcs {

		if hc.Temperature > 38 {
			score += temperature
		}
		if hc.IsTired {
			score += Tired
		}
		if hc.HasDryCough {
			score += DryCough
		}
		if hc.HasShortnessOfBreath {
			score += shortnessofbreath
		}
		if hc.HasHeadache {
			score += Headache
		}
		if hc.HasNasalCongestion {
			score += nasalcongestion
		}
		if hc.HasMusclePain {
			score += musclepain
		}
		if hc.HasDiarrhea {
			score += diarrhea
		}
	}
	score = score / len(hcs)

	if p.IsReturnFromTravel {
		score += returnfromtravel
	}
	if p.IsDiabetic {
		score += diabetic
	}
	if p.IsHypertensive {
		score += hypertensive
	}
	if p.IsAsthmatic {
		score += Asthmatic
	}
	if p.IsSmoker {
		score += smoke
	}
	if p.IsCardioIschemic {
		score += CardioIschemic
	}
	if p.HasLungDisease {
		score += lungdisease
	}
	if p.HasKidneyDisease {
		score += kidneydisease
	}
	return score > threshold
}

func (ps *patientService) IndexConstants() {

	result, err := ps.repository.NotIndexedConstants()
	if err != nil {
		log.Println(err)
		return
	}

	constantMap := make(map[string][]repository.HealthConstantResult)
	for _, hc := range result {
		constantMap[hc.PatientID] = append(constantMap[hc.PatientID], hc)
	}

	keys := make([]string, len(constantMap))
	for key, _ := range constantMap {
		keys = append(keys, key)
	}
	patients, err := ps.repository.InPatient(keys...)
	if err != nil {
		log.Println(err)
		_ = ps.repository.IndexedConstant(false, err.Error())
		return
	}

	for _, rawPatient := range patients {
		patient := ps.patientResultToPatient(rawPatient)
		values := constantMap[rawPatient.ID]
		patientJson, err := json.Marshal(patient)
		if err != nil {
			log.Println(err)
			_ = ps.repository.IndexedConstant(false, err.Error())
			return
		}
		pData := string(patientJson)
		for _, c := range values {
			constantJson, err := json.Marshal(ps.ConstantResultToHealthConstant(c))
			if err != nil {
				log.Println(err)
				_ = ps.repository.IndexedConstant(false, err.Error())
				return
			}
			cData := string(constantJson)
			doc := es.Document{
				ID:    c.ID,
				Index: "patientconstants",
				Json:  cData[:len(cData)-1] + "," + pData[1:],
			}
			err = ps.esClient.Insert(doc, true)
			if err != nil {
				log.Println(err)
				_ = ps.repository.IndexedConstant(false, err.Error())
				return
			}
		}
	}
	err = ps.repository.IndexedConstant(true, "")
	log.Println(err)

}
