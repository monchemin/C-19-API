package handler

import (
	"context"
	"log"
	"net/http"
	"strings"

	appContext "github.com/monchemin/C-19-API/context"
	"github.com/monchemin/C-19-API/security/model"

	"github.com/gin-gonic/gin"
)

func (h *handler) Login(c *gin.Context) {
	var request model.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connection, err := h.securityService.Login(request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Data:connection})
}

func (h *handler) CreateUser(c *gin.Context) {
	var request model.UserCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := c.GetHeader("Authorization")
	ctx := context.WithValue(context.Background(), appContext.TokenKey, strings.TrimPrefix(token, "Bearer "))
	ctx = context.WithValue(ctx, appContext.ResourceIDKey, request.ResourceID)
	connection, err := h.securityService.CreateUser(ctx, request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Data:connection})
}

