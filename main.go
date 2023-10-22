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
	fmt.Println("DB_DATABASE:", os.Getenv("DB_DATABASE"))

	db, err := internal.NewPostgres(internal.DBConfig{
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
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

	server.POST("/cultivo", func(ctx *gin.Context) {
		httpController.PostCultivo(ctx)
	})

	server.GET("/cultivo", func(ctx *gin.Context) {
		httpController.GetCultivo(ctx)
	})

	port := ":8080"
	log.Printf("Escuchando en el puerto%s\n", port)

	if err := server.Run(port); err != nil {
		log.Fatal(err)
	}
}
