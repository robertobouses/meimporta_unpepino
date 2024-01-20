package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/app"
	provinceApi "github.com/robertobouses/meimporta_unpepino/app/province"
	"github.com/robertobouses/meimporta_unpepino/http"
	provinceHandler "github.com/robertobouses/meimporta_unpepino/http/province"
	"github.com/robertobouses/meimporta_unpepino/internal"
	"github.com/robertobouses/meimporta_unpepino/repository"
	provinceRepo "github.com/robertobouses/meimporta_unpepino/repository/province"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASS:", os.Getenv("DB_PASS"))
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_DATABASE:", os.Getenv("DB_DATABASE"))

	db, err := internal.NewPostgres(internal.DBConfig{
		User:     "postgres",
		Pass:     "mysecretpassword",
		Host:     "localhost",
		Port:     "5432",
		Database: "mi1p",
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

	var repoProvince provinceRepo.REPOSITORYProvince
	repoProvince, err = provinceRepo.NewRepositoryProvince(db)
	if err != nil {
		panic(err)
	}

	var appController app.APP
	appController = app.NewAPP(repo, repoProvince)

	var httpController http.HTTP
	httpController = http.NewHTTP(appController)

	var appProvinceController provinceApi.APPProvince
	appProvinceController = provinceApi.NewAPP(repoProvince)

	var httpProvinceController provinceHandler.HTTPProvince
	httpProvinceController = provinceHandler.NewHTTP(appProvinceController)

	server := gin.Default()

	server.Use(internal.CORSMiddleware())

	server.POST("/crops", func(ctx *gin.Context) {
		httpController.PostCrops(ctx)

	})

	server.GET("/crops", func(ctx *gin.Context) {
		httpController.GetCrops(ctx)
	})

	server.DELETE("/crops/all", func(ctx *gin.Context) {
		httpController.DeleteCropsAll(ctx)
	})

	server.DELETE("/crops/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		httpController.DeleteCropsId(ctx, id)
	})

	server.GET("/crops/issues", func(ctx *gin.Context) {
		httpController.GetCropsIssues(ctx)
	})

	server.GET("/crops/calculate", func(ctx *gin.Context) {
		httpController.GetCropsCalculate(ctx)
	})

	server.GET("/crops/search", func(ctx *gin.Context) {
		httpController.GetCropsSearch(ctx)
	})

	server.GET("/crops/calendary", func(ctx *gin.Context) { ///TODO MATHEUS NO FUNCIONA, CREO QUE ES POR LOS ARRAY, PERO LA CONSULTA NO OBTIENE CULTIVO
		httpController.GetCropsCalendary(ctx)
	})

	server.POST("/provinces", func(ctx *gin.Context) {
		httpProvinceController.PostProvinces(ctx)
	})

	server.POST("/mycrops/fields", func(ctx *gin.Context) {
		httpController.PostMyCropsFields(ctx)
	}) //ya integra provinces y climate automaticamente, hacer prueba

	server.POST("/mycrops/define", func(ctx *gin.Context) {
		httpController.PostMyCropsDefine(ctx)
	}) //no tiene mayor dificultad, solo inserta los datos basicos de tu cultivo

	server.GET("mycrops/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		httpController.GetMyCrops(ctx, name)
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

//definir constantes para datos introduce administrador
