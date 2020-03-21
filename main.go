package main

import (
	"c19/connector/pgsql"
	"c19/handler"
	"c19/patient/repository"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	initEnv()
	pg, err := pgsql.Open()
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
}

func initEnv() {
	os.Setenv("POSTGRES_HOSTNAME", "localhost")
	os.Setenv("POSTGRES_USERNAME", "c19")
	os.Setenv("POSTGRES_PASSWORD","c19")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "C19")
}
