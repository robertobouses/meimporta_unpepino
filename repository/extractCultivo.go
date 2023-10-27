package repository

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) ExtractCultivos() ([]entity.Cultivo, error) {
	rows, err := r.db.Query(ExtractCultivosQuery)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL: %v", err)
		return []entity.Cultivo{}, err
	}
	defer rows.Close()

	var cultivos []entity.Cultivo

	for rows.Next() {
		var cultivo entity.Cultivo
		var ph string
		var requisitos entity.RequisitosCultivo
		var clima string
		var colorBytes []byte
		var colorSlice []string
		var asociacionesCSV string

		if err := rows.Scan(
			&cultivo.IdCultivo,
			&cultivo.InformacionCultivo.Nombre,
			&colorBytes,
			&cultivo.InformacionCultivo.Familia,
			&cultivo.InformacionCultivo.DensidadPlantacion,
			&cultivo.InformacionCultivo.LitrosTierraMaceta,
			&asociacionesCSV,
			&cultivo.RequisitosCultivo.Agua,
			&cultivo.RequisitosCultivo.Tierra,
			&cultivo.RequisitosCultivo.Nutricion,
			&cultivo.RequisitosCultivo.Salinidad,
			&ph,
			&clima,
			&cultivo.RequisitosCultivo.Profundidad,
			&cultivo.FechasCultivo.Siembra,
			&cultivo.FechasCultivo.Transplante,
			&cultivo.FechasCultivo.Cosecha,
			&cultivo.FechasCultivo.Ciclo,
			&cultivo.FrutoCultivo.Produccion,
			&cultivo.FrutoCultivo.Nutrientes,
			&cultivo.SemillaCultivo.Semilla,
			&cultivo.SemillaCultivo.SemillasGramo,
			&cultivo.SemillaCultivo.VidaSemilla,
			&cultivo.ProblemasCultivo.Plagas,
			&cultivo.ProblemasCultivo.Dificultades,
			&cultivo.ProblemasCultivo.Cuidados,
			&cultivo.ProblemasCultivo.Miscelanea,
		); err != nil {
			log.Printf("Error al escanear filas: %v", err)
			return nil, err
		}

		ph = strings.Replace(ph, "{", "", -1)
		ph = strings.Replace(ph, "}", "", -1)
		phValues := strings.Split(ph, ",")
		phFloats := make([]float64, len(phValues))
		for i, phStr := range phValues {
			phFloat, err := strconv.ParseFloat(phStr, 64)
			if err != nil {
				log.Printf("Error al convertir 'ph' en float64: %v", err)
				return nil, err
			}
			phFloats[i] = phFloat
		}

		requisitos.Ph = phFloats

		requisitos.Clima = []string{clima}

		err = json.Unmarshal(colorBytes, &colorSlice)
		if err != nil {

		}

		cultivo.InformacionCultivo.Color = colorSlice
		cultivo.InformacionCultivo.Asociaciones = strings.Split(asociacionesCSV, ",")

		cultivo.RequisitosCultivo = requisitos

		cultivos = append(cultivos, cultivo)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al recorrer filas: %v", err)
		return nil, err
	}

	return cultivos, nil
}
