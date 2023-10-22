package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) GetCultivo(ctx *gin.Context) {
	cultivo, err := h.service.PrintCultivo()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, cultivo)

}
