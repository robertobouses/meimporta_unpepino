package http

import (
	"fmt"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (h *Http) PostCrop(ctx *gin.Context) {
	var newCrop entity.Crop
	err := ctx.BindJSON(&newCrop)
	fmt.Print("crop en PostCrop", newCrop)
	fmt.Print("Id crop en PostCrop", newCrop.IdCrop)

	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al hacer BindJSON": err.Error()})
		return
	}

	fmt.Printf("Valores de newCrop: %+v\n", newCrop)

	err = h.service.CreateCrop(newCrop)
	fmt.Print("crop en PostCrop2", newCrop)
	fmt.Print("Id crop en PostCrop2", newCrop.IdCrop)
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	//cropID := newCrop.IdCrop
	ctx.JSON(nethttp.StatusOK, gin.H{"idcrop": newCrop.IdCrop})
	ctx.JSON(nethttp.StatusOK, newCrop)
	fmt.Printf("Crop insertado con ID: %d\n", newCrop.IdCrop)
	ctx.JSON(nethttp.StatusOK, gin.H{"mensaje": "Crop insertado correctamente"})
}
