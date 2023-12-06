package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) GetMyCrops(ctx *gin.Context, name string) {
	mycrops, err := h.service.ProcessMyCrops(name)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, mycrops)

}
