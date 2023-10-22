package repository

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) ExtractCultivos() ([]entity.Cultivo, error) {

	rows, err := r.db.Query(ExtractCultivosQuery)

	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL", err)
		return []entity.Cultivo{}, err
	}
	defer rows.Close()

	var cultivos []entity.Cultivo

	for rows.Next() {
		var cultivo entity.Cultivo
		if err := rows.Scan(
			&cultivo.IdCultivo,
			&cultivo.InformacionCultivo.Nombre,
			&cultivo.InformacionCultivo.Color,
			&cultivo.InformacionCultivo.Familia,
			&cultivo.InformacionCultivo.DensidadPlantacion,
			&cultivo.InformacionCultivo.LitrosTierraMaceta,
			&cultivo.InformacionCultivo.Asociaciones,
			&cultivo.RequisitosCultivo.Agua,
			&cultivo.RequisitosCultivo.Tierra,
			&cultivo.RequisitosCultivo.Nutricion,
			&cultivo.RequisitosCultivo.Salinidad,
			&cultivo.RequisitosCultivo.Ph,
			&cultivo.RequisitosCultivo.Clima,
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
			log.Printf("Error al escanear filas", err)
			return nil, err
		}
		cultivos = append(cultivos, cultivo)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al recorrer filas", err)
		return nil, err
	}

	return cultivos, nil
}
