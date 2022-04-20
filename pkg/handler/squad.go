package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getSquad(ctx *gin.Context) {
	squad, _ := h.service.GetSquad()
	ctx.JSON(http.StatusOK, squad)
}
