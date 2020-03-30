package main

import (
	"os"

	"c19/connector/es"
	"c19/connector/pgsql"
	"c19/handler"

	"github.com/gin-gonic/gin"
)

func main() {
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
