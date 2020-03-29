package main

import (
	"c19/connector/es"
	"c19/connector/pgsql"
	"c19/handler"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	os.Setenv("POSTGRES_HOSTNAME", "pg.co9rbmwsfcbm.us-west-2.rds.amazonaws.com")
	os.Setenv("POSTGRES_USERNAME", "C19")
	os.Setenv("POSTGRES_PASSWORD","OEBQLO5FO6oYidsupHiu")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "C19")
	os.Setenv("ES_URL", "https://search-c-19-i2xjjsrykhr2wb57yoaaejnkme.us-east-2.es.amazonaws.com")
}

func main() {
	//initlocal()
	pg, err := pgsql.Open()
	if err != nil {
		panic(err)
	}
	defer pg.Close()

	esClient, err := es.Open([]string{os.Getenv("ES_URL")}, "", "")
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router = handler.Setup(router, pg, esClient)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func initlocal() {
	os.Setenv("POSTGRES_HOSTNAME", "localhost")
	os.Setenv("POSTGRES_USERNAME", "c19")
	os.Setenv("POSTGRES_PASSWORD","c19")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "C19")
	os.Setenv("ES_URL", "http://localhost:9200")
}
