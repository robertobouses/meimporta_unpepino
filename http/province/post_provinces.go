package province

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity/province"
)

func (h *Http) PostProvinces(ctx *gin.Context) {
	var newProvince province.Province
	err := ctx.BindJSON(&newProvince)

	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al hacer BindJSON": err.Error()})
		return
	}

	err = h.serviceProvince.CreateProvinces(newProvince)
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	response := gin.H{
		"province": newProvince,
		"mensaje":  "Province insertada correctamente",
	}

	ctx.JSON(nethttp.StatusOK, response)
}
