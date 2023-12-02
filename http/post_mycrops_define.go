package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
)

func (h *Http) PostMyCropsDefine(ctx *gin.Context) {
	var mycrop mycrop.MyCrop
	err := ctx.BindJSON(&mycrop)

	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error BindJSON": err.Error()})
		return
	}

	log.Print("my crop", mycrop)

	err = h.service.CreateMyCrops(mycrop)
	if err != nil {
		log.Println("La finca que le quieres asignar al cultivo no está creada:", err)
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	log.Println("Cultivo personal insertado correctamente")
	ctx.JSON(nethttp.StatusOK, gin.H{
		"cultivo personal": mycrop,
		"message":          "Cultivo personal insertado correctamente",
	})
}
