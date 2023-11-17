package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/app"
	"github.com/robertobouses/meimporta_unpepino/http"
	"github.com/robertobouses/meimporta_unpepino/internal"
	"github.com/robertobouses/meimporta_unpepino/repository"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASS:", os.Getenv("DB_PASS"))
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	//fmt.Println("DB_DATABASE:", "mi1p")
	fmt.Println("DB_DATABASE:", os.Getenv("DB_DATABASE"))

	db, err := internal.NewPostgres(internal.DBConfig{
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
		//Database: os.Getenv("mi1p"),
	})

	if err != nil {
		panic(err)
	}
	log.Println("la conexión se abrió correctamente")
	defer db.Close()

	var repo repository.REPOSITORY
	repo, err = repository.NewRepository(db)
	if err != nil {
		panic(err)
	}
	var appController app.APP
	var httpController http.HTTP

	appController = app.NewAPP(repo)
	httpController = http.NewHTTP(appController)

	server := gin.Default()

	server.Use(internal.CORSMiddleware())

	server.POST("/createCultivo", func(ctx *gin.Context) {
		httpController.PostCultivo(ctx)

	})

	server.GET("/printCultivos", func(ctx *gin.Context) {
		httpController.GetCultivos(ctx)
	})

	server.DELETE("/deleteAllCultivos", func(ctx *gin.Context) {
		httpController.DeleteCultivosAll(ctx)
	})

	server.DELETE("/deleteCultivo/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		httpController.DeleteCultivosId(ctx, id)
	})

	server.POST("/problemsCultivo", func(ctx *gin.Context) {
		httpController.PostProblemsCultivo(ctx)
	})

	server.POST("/calculateCultivo", func(ctx *gin.Context) {
		httpController.PostCalculateCultivo(ctx)
	})

	server.POST("/searchCultivo", func(ctx *gin.Context) {
		httpController.PostSearchCultivo(ctx)
	})
	//calendario cultivo segun mes y espacio tierra
	//delete drop table tal
	//delete all table
	//una vez tengas el print de lista de cultivos puedes comprobar si existe al eliminar por id por ejemplo
	port := ":8080"
	log.Printf("Escuchando en el puerto%s\n", port)
	if err := server.Run(port); err != nil {
		log.Fatal(err)
	}
}
