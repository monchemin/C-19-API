package handler

import (
	"log"
	"net/http"

	"github.com/monchemin/C-19-API/patient/model"

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
	c.JSON(http.StatusOK, Response{ID: id})
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
	c.JSON(http.StatusOK, Response{Data:patient})
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
	c.JSON(http.StatusOK, Response{ID: id})
}

func (h *handler) Connexion(c *gin.Context) {
	var request model.GetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connection, err := h.patientService.Connect(request.PhoneNumber)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Data:connection})
}

func (h *handler) NewTestResult(c *gin.Context) {
	var request model.TestResultRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.patientService.NewTestResult(request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{ID: id})
}

func (h *handler) ReadPatientTestResult(c *gin.Context) {
	var request model.GetTestResultByPatientRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient, err := h.patientService.PatientTestResult(request.PatientID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Data:patient})
}