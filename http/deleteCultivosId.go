package http

import (
	"errors"
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/app"
)

func (h *Http) DeleteCropsId(ctx *gin.Context, id string) {
	err := h.service.DeleteCropsId(id)
	log.Print("el valor del id en la función de la capa http es", id)
	if err != nil {
		if errors.Is(err, app.ErrCropNoExiste) {
			log.Print("Error al llamar http a app. Función delete", err)
			ctx.JSON(nethttp.StatusNotFound, gin.H{"error": "El crop no se encontró"})
		} else {
			log.Print("Error al llamar http a app. Función delete", err)
			ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Error al eliminar el crop"})
		}
	} else {
		ctx.JSON(nethttp.StatusOK, "El crop se ha eliminado correctamente")
	}
}
