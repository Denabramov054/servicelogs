package controller

import (
	"project/servicelogs/pkg/service"

	"github.com/gin-gonic/gin"
)

type ControllerMain struct {
	services *service.Service
}

func NewController(services *service.Service) *ControllerMain {
	return &ControllerMain{
		services: services,
	}
}

func (h *ControllerMain) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/statistics")
	}

	return router
}
