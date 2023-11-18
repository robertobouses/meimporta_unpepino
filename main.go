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
		User:     "postgres",
		Pass:     "mysecretpassword",
		Host:     "localhost",
		Port:     "5432",
		Database: "mi1p",
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

	server.POST("/createCrop", func(ctx *gin.Context) {
		httpController.PostCrop(ctx)

	})

	server.GET("/printCrops", func(ctx *gin.Context) {
		httpController.GetCrops(ctx)
	})

	server.DELETE("/deleteAllCrops", func(ctx *gin.Context) {
		httpController.DeleteCropsAll(ctx)
	})

	server.DELETE("/deleteCrop/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		httpController.DeleteCropsId(ctx, id)
	})

	server.POST("/problemsCrop", func(ctx *gin.Context) {
		httpController.PostProblemsCrop(ctx)
	})

	server.POST("/calculateCrop", func(ctx *gin.Context) {
		httpController.PostCalculateCrop(ctx)
	})

	server.POST("/searchCrop", func(ctx *gin.Context) {
		httpController.PostSearchCrop(ctx)
	})
	//calendario crop segun mes y espacio soil
	//delete drop table tal
	//delete all table
	//una vez tengas el print de lista de crops puedes comprobar si existe al eliminar por id por ejemplo
	port := ":8080"
	log.Printf("Escuchando en el puerto%s\n", port)
	if err := server.Run(port); err != nil {
		log.Fatal(err)
	}
}
