package main

import (
	"c19/connector/es"
	"c19/connector/pgsql"
	"c19/patient/repository"
	"c19/patient/service"
	"os"
)

func main() {
	pg, err := pgsql.Open()
	if err != nil {
		panic(err)
	}
	defer pg.Close()

	esClient, err := es.Open([]string{os.Getenv("ES_URL")}, os.Getenv("ES_USER"), os.Getenv("ES_PASSWORD"))
	if err != nil {
		panic(err)
	}
	patientRepository := repository.NewPatientRepository(pg)
	patientService := service.NewPatientService(patientRepository, esClient)
	patientService.IndexPatients()

}