package repository

import (
	"github.com/lib/pq"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) InsertCultivo(cultivo entity.Cultivo) error {
	_, err := r.insertCultivoStmt.Exec(
		cultivo.IdCultivo,
		cultivo.InformacionCultivo.Nombre,
		pq.Array(cultivo.InformacionCultivo.Color),
		cultivo.InformacionCultivo.Familia,
		cultivo.InformacionCultivo.DensidadPlantacion,
		cultivo.InformacionCultivo.LitrosTierraMaceta,
		pq.Array(cultivo.InformacionCultivo.Asociaciones),
		cultivo.RequisitosCultivo.Agua,
		cultivo.RequisitosCultivo.Tierra,
		cultivo.RequisitosCultivo.Nutricion,
		cultivo.RequisitosCultivo.Salinidad,
		pq.Array(cultivo.RequisitosCultivo.Ph),
		pq.Array(cultivo.RequisitosCultivo.Clima),
		cultivo.RequisitosCultivo.Profundidad,
		cultivo.FechasCultivo.Siembra,
		cultivo.FechasCultivo.Transplante,
		cultivo.FechasCultivo.Cosecha,
		cultivo.FechasCultivo.Ciclo,
		cultivo.FrutoCultivo.Produccion,
		cultivo.FrutoCultivo.Nutrientes,
		cultivo.SemillaCultivo.Semilla,
		cultivo.SemillaCultivo.SemillasGramo,
		cultivo.SemillaCultivo.VidaSemilla,
		cultivo.ProblemasCultivo.Plagas,
		cultivo.ProblemasCultivo.Dificultades,
		cultivo.ProblemasCultivo.Cuidados,
		cultivo.ProblemasCultivo.Miscelanea,
	)
	return err
}
