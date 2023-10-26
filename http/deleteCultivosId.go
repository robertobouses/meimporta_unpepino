package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) DeleteCultivosId(ctx *gin.Context, id string) {
	err := h.service.DeleteCultivosId(id)
	log.Print("el valor del id en la función de la capa http es", id)
	if err != nil {
		log.Print("Error al llamar http a app. Función delete", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(nethttp.StatusOK, "se ha borrado correctamente")
	}
}
