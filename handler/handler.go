package handler

import (
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
}
type handler struct {
	patientService  service.PatientService
	positionService service2.PositionService
}

func Setup(router *gin.Engine, pg *pgsql.DB) *gin.Engine {
	patientRepository := repository.NewPatientRepository(pg)
	patientService := service.NewPatientService(patientRepository)
	positionRepository := repository2.NewPositionRepository(pg)
	positionService := service2.NewPositionService(positionRepository)
	handler := handler{
		patientService:  patientService,
		positionService: positionService,
	}

	router.POST("/patient/add", handler.NewPatient)
	router.POST("/patient/read", handler.ReadPatient)
	router.POST("/constant/add", handler.NewHealthConstant)
	routerPosition := router.Group("/position")
	{
		routerPosition.POST("/country/add", handler.NewCountry)
		routerPosition.POST("/town/add", handler.NewTown)
		routerPosition.POST("/district/add", handler.NewDistrict)
	}
	return router
}
