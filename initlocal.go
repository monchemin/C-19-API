package main

import (
	"c19/connector/pgsql"
	"c19/handler"
	"c19/patient/model"
	"c19/patient/repository"
	"c19/patient/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func initE() {

	pg, err := pgsql.OpenWithConfig(pgsql.Config{
		HostName:     "localhost",
		UserName:     "c19",
		Password:     "c19",
		Port:         5432,
		DataBaseName: "C19",
	})
	if err != nil {
		panic(err)
	}
	defer pg.Close()

	repo := repository.NewPatientRepository(pg)
	router := gin.Default()

	router = handler.Setup(router, repo)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

	pservice := service.NewPatientService(repo)
	id, err := pservice.Add(model.PatientRequest{
		PhoneNumber:        "4357645",
		Age:                20,
		Weight:             45,
		IsDiabetic:         false,
		IsHypertensive:     false,
		IsAsthmatic:        false,
		IsCardioIschemic:   false,
		HasLungDisease:     false,
		HasKidneyDisease:   false,
		IsSmoker:           false,
		IsReturnFromTravel: false,
		Longitude:          0,
		Latitude:           0,
	})
	fmt.Printf("patient id: %s\n error: %v", id, err)
	id, err = pservice.AddHealthConstant(model.HealthConstantRequest{
		PatientID:                          id,
		Temperature:                        36.5,
		IsTired:                            false,
		HasDryCough:                        false,
		HasShortnessOfBreath:               false,
		HasBeenInContactWithInfectedPerson: false,
		HasHeadache:                        false,
		HasRunnyNose:                       false,
		HasNasalCongestion:                 false,
		HasSoreThroat:                      false,
		HasMusclePain:                      false,
		HasDiarrhea:                        false,
	})
	fmt.Printf("constantid: %s\n error: %v", id, err)
}

func initEnv2() {
	os.Setenv("POSTGRES_HOSTNAME", "localhost")
	os.Setenv("POSTGRES_USERNAME", "c19")
	os.Setenv("POSTGRES_PASSWORD","c19")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "C19")
}