package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	squad := router.Group("/squad")
	{
		squad.GET("/", h.getSquad)
	}

	return router
}
