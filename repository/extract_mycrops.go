package repository

// import (
// 	"log"

// 	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
// )

// func (r *Repository) ExtractMyCrops(name string) (mycrop.MyCropResponse, error) {
// 	//var crop mycrop.CropResponse

// 	log.Printf("Consulta SQL: %s", ExtractMyCropsNameQuery)
// 	log.Printf("name del cultivo: %d", name)

// 	rows, err := r.db.Query(ExtractMyCropsNameQuery, name)
// 	if err != nil {
// 		log.Printf("Error al ejecutar la consulta: %v", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var myCropResponse mycrop.MyCropResponse

// 		if err := rows.Scan(
// 			&myCropResponse.IdMyCrop,
// 			&myCropResponse.Name,
// 			&myCropResponse.Planting,
// 			&myCropResponse.SquareMeters,
// 			&myCropResponse.Soil,
// 			&myCropResponse.Climate,
// 		); err != nil {
// 			log.Printf("Error al escanear las filas: %v", err)
// 			return nil, err
// 		}

// 		if err != nil {
// 			log.Printf("Error al ejecutar la consulta: %v", err)
// 			return nil, err
// 		}

// 		return &myCropResponse, nil
// 	}
// }
