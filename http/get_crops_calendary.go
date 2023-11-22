package http

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) GetCropsCalendary(ctx *gin.Context) {
	month := ctx.Query("month")
	city := ctx.Query("city")
	log.Printf("Mes: %s, Ciudad: %s\n", month, city)

	provinceName, err := GetProvince(city)
	if err != nil {
		log.Printf("Error al obtener la provincia: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}
	log.Printf("Provincia obtenida: %s\n", provinceName)
	crops, err := h.service.ProcessCropsCalendary(month, provinceName)
	ctx.JSON(nethttp.StatusOK, crops)
}
