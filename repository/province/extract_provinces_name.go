package province

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity/province"
)

func (r *RepositoryProvince) ExtractProvincesName(name string) (province.Province, error) {
	var provinceResult province.Province

	err := r.db.QueryRow(ExtractProvincesNameQuery, name).
		Scan(
			&provinceResult.NameProvince,
			&provinceResult.ClimateProvince,
		)
	fmt.Println("densidad plantacion repo", provinceResult.ClimateProvince)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No se encontraron filas para la provincia con nombre: %s", name)
			return province.Province{}, nil
		}
		log.Printf("Error al ejecutar la consulta SQL: %v", err)
		return province.Province{}, err
	}
	log.Println("crop en Repo", provinceResult)
	return provinceResult, nil
}
