package handler

import (
	"c19/patient/model"
	"c19/patient/repository"
	"c19/patient/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler interface {
	NewPatient(c *gin.Context)
	ReadPatient(c *gin.Context)
	NewHealthConstant(c *gin.Context)
}
type handler struct {
	service service.PatientService
}

func Setup(router *gin.Engine, repository repository.PatientRepository) *gin.Engine {
	patientService := service.NewPatientService(repository)
	handler := handler{service: patientService}
	router.POST("/patient/add", handler.NewPatient)
	router.POST("/patient/read", handler.ReadPatient)
	router.POST("/constant/add", handler.NewHealthConstant)
	return router
}

func (h *handler) NewPatient(c *gin.Context) {
	var request model.PatientRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.Add(request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.CreationResponse{ID:id})
}

func (h *handler) ReadPatient(c *gin.Context) {
	var request model.GetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient, err := h.service.Patient(request.PhoneNumber)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func (h *handler) NewHealthConstant(c *gin.Context) {
	var request model.HealthConstantRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.AddHealthConstant(request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.CreationResponse{ID:id})
}

