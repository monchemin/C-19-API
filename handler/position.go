package handler

import (
	"log"
	"net/http"

	"github.com/monchemin/C-19-API/position/model"

	"github.com/gin-gonic/gin"
)

func (h *handler) NewCountry(c *gin.Context) {
	var request model.CountryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.positionService.NewCountry(request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "")
}

func (h *handler) NewTown(c *gin.Context) {
	var request model.TownRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.positionService.NewTown(request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{ID: id})
}

func (h *handler) NewDistrict(c *gin.Context) {
	var request model.DistrictRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.positionService.NewDistrict(request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{ID: id})
}

func (h *handler) Countries(c *gin.Context) {
	countries, err := h.positionService.Countries()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Data: countries})
}

func (h *handler) Localizations(c *gin.Context) {
	localizations, err := h.positionService.Localizations()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Data: localizations})
}

func (h *handler) PatientList(c *gin.Context) {
	res, err := h.patientService.PatientList()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Data: res})
}
