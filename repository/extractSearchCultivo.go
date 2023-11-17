package repository

import (
	"log"

	"github.com/lib/pq"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) ExtractSearchCultivo(request entity.SearchRequest) ([]entity.Cultivo, error) {
	var cultivos []entity.Cultivo

	rows, err := r.db.Query(ExtractSearchCultivosQuery,
		request.Nombre,
		pq.Array(request.Color),
		request.DensidadPlantacion,
		request.Agua,
		request.Tierra,
		request.Nutricion,
		request.Salinidad,
		request.Ciclo,
	)

	log.Printf("rows tras la Query en extractSearchCultivo del repo", rows)
	log.Printf("Consulta SQL completa: %s", ExtractSearchCultivosQuery)
	log.Printf("Valores de la consulta: nombre=%s, color=%v, agua=%v, ...", request.Nombre, request.Color, request.Agua)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cultivo entity.Cultivo
		var colorBytes []byte
		var asociacionesCSV string
		var ph, clima float64

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

		log.Println("Cultivo en Repo:", cultivo)
		cultivos = append(cultivos, cultivo)
	}

	return cultivos, nil
}
