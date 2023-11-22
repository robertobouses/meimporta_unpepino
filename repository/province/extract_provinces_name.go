package province

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/robertobouses/meimporta_unpepino/entity/province"
)

func (r *RepositoryProvince) ExtractProvincesName(name string) (province.Province, error) {
	var provinceResult province.Province
	normalizedName := normalizeName(name)
	fmt.Println("Nombre de la provincia:", normalizedName)
	err := r.db.QueryRow(ExtractProvincesNameQuery, normalizedName).
		Scan(
			&provinceResult.NameProvince,
			&provinceResult.ClimateProvince,
		)
	fmt.Println("nombre province en el repo:", provinceResult.NameProvince)

	fmt.Println("climate province en el repo:", provinceResult.ClimateProvince)
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
func normalizeName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}
