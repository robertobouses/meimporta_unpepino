package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) GetCrops(ctx *gin.Context) {
	crops, err := h.service.PrintCrops()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, crops)

}
