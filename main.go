package main

import (
	"os"

	"github.com/monchemin/C-19-API/connector/es"
	"github.com/monchemin/C-19-API/connector/pgsql"
	"github.com/monchemin/C-19-API/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	os.Setenv("POSTGRES_HOSTNAME", "3.217.233.250")
	os.Setenv("POSTGRES_USERNAME", "C19")
	os.Setenv("POSTGRES_PASSWORD","c19PG@0520")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "C19")
	os.Setenv("ES_URL", "http://3.217.233.250:9200")
	os.Setenv("ES_USER", "user_api_go")
	os.Setenv("ES_PASSWORD", "c19apigo@042020")
	os.Setenv("SU", "root")
}

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

	router := gin.Default()
	router.Use(cors.Default())
	router = handler.Setup(router, pg, esClient)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
