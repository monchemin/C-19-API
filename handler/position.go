package handler

import (
	"log"
	"net/http"

	m "c19/patient/model"
	"c19/position/model"

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
	c.JSON(http.StatusOK, m.CreationResponse{ID: id})
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
	c.JSON(http.StatusOK, m.CreationResponse{ID: id})
}
