package repository

import (
	"fmt"

	"github.com/lib/pq"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) InsertCultivo(cultivo entity.Cultivo) error {
	fmt.Println("InsertCultivo - IdCultivo:", cultivo.IdCultivo)
	fmt.Printf("Valor que se pasa a la sentencia SQL: %d\n", cultivo.IdCultivo)

	_, err := r.insertCultivoStmt.Exec(cultivo.IdCultivo)
	if err != nil {
		fmt.Println("Error en insertCultivoStmt.Exec:", err)
		return err
	}

	_, err = r.insertInformacionCultivoStmt.Exec(
		cultivo.InformacionCultivo.Nombre,
		pq.Array(cultivo.InformacionCultivo.Color),
		cultivo.InformacionCultivo.Familia,
		cultivo.InformacionCultivo.DensidadPlantacion,
		cultivo.InformacionCultivo.LitrosTierraMaceta,
		pq.Array(cultivo.InformacionCultivo.Asociaciones),
	)
	if err != nil {
		return err
	}
	_, err = r.insertRequisitosCultivoStmt.Exec(
		cultivo.RequisitosCultivo.Agua,
		cultivo.RequisitosCultivo.Tierra,
		cultivo.RequisitosCultivo.Nutricion,
		cultivo.RequisitosCultivo.Salinidad,
		pq.Array(cultivo.RequisitosCultivo.Ph),
		pq.Array(cultivo.RequisitosCultivo.Clima),
		cultivo.RequisitosCultivo.Profundidad,
	)
	if err != nil {
		return err
	}
	_, err = r.insertFechasCultivoStmt.Exec(
		cultivo.FechasCultivo.Siembra,
		cultivo.FechasCultivo.Transplante,
		cultivo.FechasCultivo.Cosecha,
		cultivo.FechasCultivo.Ciclo,
	)
	if err != nil {
		return err
	}
	_, err = r.insertFrutoCultivoStmt.Exec(
		cultivo.FrutoCultivo.Produccion,
		cultivo.FrutoCultivo.Nutrientes,
	)
	if err != nil {
		return err
	}
	_, err = r.insertSemillaCultivoStmt.Exec(
		cultivo.SemillaCultivo.Semilla,
		cultivo.SemillaCultivo.SemillasGramo,
		cultivo.SemillaCultivo.VidaSemilla,
	)
	if err != nil {
		return err
	}
	_, err = r.insertProblemasCultivoStmt.Exec(
		cultivo.ProblemasCultivo.Plagas,
		cultivo.ProblemasCultivo.Dificultades,
		cultivo.ProblemasCultivo.Cuidados,
		cultivo.ProblemasCultivo.Miscelanea,
	)
	return err
}
