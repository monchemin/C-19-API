package main

import (
	"github.com/gin-contrib/cors"
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
