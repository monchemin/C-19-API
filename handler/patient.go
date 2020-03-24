package handler

import (
	"log"
	"net/http"

	"c19/patient/model"

	"github.com/gin-gonic/gin"
)

func (h *handler) NewPatient(c *gin.Context) {
	var request model.PatientRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.patientService.Add(request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.CreationResponse{ID: id})
}

func (h *handler) ReadPatient(c *gin.Context) {
	var request model.GetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient, err := h.patientService.Patient(request.PhoneNumber)
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

	id, err := h.patientService.AddHealthConstant(request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.CreationResponse{ID: id})
}