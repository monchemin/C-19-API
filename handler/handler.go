package handler

import (
	"c19/connector/es"
	"c19/connector/pgsql"
	"c19/patient/repository"
	"c19/patient/service"
	repository2 "c19/position/repository"
	service2 "c19/position/service"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	NewPatient(c *gin.Context)
	ReadPatient(c *gin.Context)
	NewHealthConstant(c *gin.Context)
	NewCountry(c *gin.Context)
	NewTown(c *gin.Context)
	NewDistrict(c *gin.Context)
	Countries(c *gin.Context)
	Localizations(c *gin.Context)
}
type handler struct {
	patientService  service.PatientService
	positionService service2.PositionService
}

func Setup(router *gin.Engine, pg *pgsql.DB, esClient es.ElasticSearchClient) *gin.Engine {
	patientRepository := repository.NewPatientRepository(pg)
	patientService := service.NewPatientService(patientRepository, esClient)
	positionRepository := repository2.NewPositionRepository(pg)
	positionService := service2.NewPositionService(positionRepository)
	handler := handler{
		patientService:  patientService,
		positionService: positionService,
	}
	patientRouter := router.Group("/patient")
	{
		patientRouter.POST("/add", handler.NewPatient)
		patientRouter.POST("/read", handler.ReadPatient)
		patientRouter.POST("/connect", handler.Connexion)
		patientRouter.POST("/constant/add", handler.NewHealthConstant)
	}

	routerPosition := router.Group("/position")
	{
		routerPosition.POST("/country/add", handler.NewCountry)
		routerPosition.POST("/town/add", handler.NewTown)
		routerPosition.POST("/district/add", handler.NewDistrict)
		routerPosition.GET("/countries", handler.Countries)
		routerPosition.GET("/localizations", handler.Localizations)
	}
	return router
}
