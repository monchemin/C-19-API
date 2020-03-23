package main

import (
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

	router := gin.Default()
	router = handler.Setup(router, pg)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

