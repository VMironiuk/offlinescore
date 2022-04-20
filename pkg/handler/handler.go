package handler

import (
	"github.com/VMironiuk/offlinescore/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	squad := router.Group("/squad")
	{
		squad.GET("/", h.getSquad)
	}

	return router
}
