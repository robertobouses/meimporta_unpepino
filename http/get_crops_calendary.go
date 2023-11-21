package http

import (
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) GetCropsCalendary(ctx *gin.Context) {
	month := ctx.Query("month")
	city := ctx.Query("city")

	province, err := GetProvince(city)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	crops, err := h.service.ProcessCropsCalendary(month, province)
	ctx.JSON(nethttp.StatusOK, crops)
}
