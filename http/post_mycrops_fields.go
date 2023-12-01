package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
)

func (h *Http) PostMyCropsFields(ctx *gin.Context) {
	var field mycrop.MyField

	err := ctx.BindJSON(&field)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error BindJSON": err.Error()})
		return
	}
	log.Print(field)

	err = h.service.CreateFields(field)
	if err != nil {
		log.Println("Error al llamar desde HTTP al servicio:", err)
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	log.Println("Finca insertada correctamente")
	ctx.JSON(nethttp.StatusOK, gin.H{
		"field":   field,
		"message": "Finca insertada correctamente",
	})
}
