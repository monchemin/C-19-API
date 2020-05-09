package main

import (
	"os"

	"github.com/monchemin/C-19-API/connector/es"
	"github.com/monchemin/C-19-API/connector/pgsql"
	"github.com/monchemin/C-19-API/patient/repository"
	"github.com/monchemin/C-19-API/patient/service"
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