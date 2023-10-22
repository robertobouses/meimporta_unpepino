package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) GetCultivos(ctx *gin.Context) {
	cultivos, err := h.service.PrintCultivos()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, cultivos)

}
