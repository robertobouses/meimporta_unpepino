package repository

import (
	"database/sql"

	_ "embed"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

type REPOSITORY interface {
	InsertCultivo(cultivo entity.Cultivo) error
	ExtractCultivos() ([]entity.Cultivo, error)
	DeleteCultivosAll() error
	DeleteCultivosId(id int) error
	CheckExistCultivoId(id int) (bool, error)
	ExtractCultivoName(name string) (entity.CalculateResponse, error)
}

type Repository struct {
	db                           *sql.DB
	insertCultivoStmt            *sql.Stmt
	insertInformacionCultivoStmt *sql.Stmt
	insertRequisitosCultivoStmt  *sql.Stmt
	insertFechasCultivoStmt      *sql.Stmt
	insertFrutoCultivoStmt       *sql.Stmt
	insertSemillaCultivoStmt     *sql.Stmt
	insertProblemasCultivoStmt   *sql.Stmt
	deleteCultivosAllStmt        *sql.Stmt
	deleteCultivosIdStmt         *sql.Stmt

	checkExistCultivoIdStmt *sql.Stmt

	extractCultivosStmt    *sql.Stmt
	extractCultivoNameStmt *sql.Stmt
}

//go:embed sql/insert_cultivo.sql
var InsertCultivoQuery string

//go:embed sql/insert_informacion_cultivo.sql
var InsertInformacionCultivoQuery string

//go:embed sql/insert_requisitos_cultivo.sql
var InsertRequisitosCultivoQuery string

//go:embed sql/insert_fechas_cultivo.sql
var InsertFechasCultivoQuery string

//go:embed sql/insert_fruto_cultivo.sql
var InsertFrutoCultivoQuery string

//go:embed sql/insert_semilla_cultivo.sql
var InsertSemillaCultivoQuery string

//go:embed sql/insert_problemas_cultivo.sql
var InsertProblemasCultivoQuery string

//go:embed sql/extract_cultivos.sql
var ExtractCultivosQuery string

//go:embed sql/extract_cultivo_name.sql
var ExtractCultivoNameQuery string

//go:embed sql/delete_cultivos_all.sql
var DeleteCultivosAllQuery string

//go:embed sql/delete_cultivos_id.sql
var DeleteCultivosIdQuery string

//go:embed sql/check_exist_cultivo_id.sql
var CheckExistCultivoIdQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	insertCultivoStmt, err := db.Prepare(InsertCultivoQuery)
	if err != nil {
		return nil, err
	}

	insertInformacionCultivoStmt, err := db.Prepare(InsertInformacionCultivoQuery)
	if err != nil {
		return nil, err
	}

	insertRequisitosCultivoStmt, err := db.Prepare(InsertRequisitosCultivoQuery)
	if err != nil {
		return nil, err
	}

	insertFechasCultivoStmt, err := db.Prepare(InsertFechasCultivoQuery)
	if err != nil {
		return nil, err
	}

	insertFrutoCultivoStmt, err := db.Prepare(InsertFrutoCultivoQuery)
	if err != nil {
		return nil, err
	}

	insertSemillaCultivoStmt, err := db.Prepare(InsertSemillaCultivoQuery)
	if err != nil {
		return nil, err
	}

	insertProblemasCultivoStmt, err := db.Prepare(InsertProblemasCultivoQuery)
	if err != nil {
		return nil, err
	}

	extractCultivosStmt, err := db.Prepare(ExtractCultivosQuery)
	if err != nil {
		return nil, err
	}
	extractCultivoNameStmt, err := db.Prepare(ExtractCultivoNameQuery)
	if err != nil {
		return nil, err
	}
	deleteCultivosAllStmt, err := db.Prepare(DeleteCultivosAllQuery)
	if err != nil {
		return nil, err
	}
	deleteCultivosIdStmt, err := db.Prepare(DeleteCultivosIdQuery)
	if err != nil {
		return nil, err
	}
	checkExistCultivoIdStmt, err := db.Prepare(CheckExistCultivoIdQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                           db,
		insertCultivoStmt:            insertCultivoStmt,
		insertInformacionCultivoStmt: insertInformacionCultivoStmt,
		insertRequisitosCultivoStmt:  insertRequisitosCultivoStmt,
		insertFechasCultivoStmt:      insertFechasCultivoStmt,
		insertFrutoCultivoStmt:       insertFrutoCultivoStmt,
		insertSemillaCultivoStmt:     insertSemillaCultivoStmt,
		insertProblemasCultivoStmt:   insertProblemasCultivoStmt,
		extractCultivosStmt:          extractCultivosStmt,
		extractCultivoNameStmt:       extractCultivoNameStmt,
		deleteCultivosAllStmt:        deleteCultivosAllStmt,
		deleteCultivosIdStmt:         deleteCultivosIdStmt,
		checkExistCultivoIdStmt:      checkExistCultivoIdStmt,
	}, nil
}
