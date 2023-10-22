package repository

import (
	"database/sql"

	_ "embed"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

type REPOSITORY interface {
	InsertCultivo(cultivo entity.Cultivo) error
	ExtractCultivos() ([]entity.Cultivo, error)
}

type Repository struct {
	db                  *sql.DB
	insertCultivoStmt   *sql.Stmt
	extractCultivosStmt *sql.Stmt
}

//go:embed sql/insert_cultivo.sql
var InsertCultivoQuery string

//go:embed sql/extract_cultivos.sql
var ExtractCultivosQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	insertCultivoStmt, err := db.Prepare(InsertCultivoQuery)
	if err != nil {
		return nil, err
	}
	extractCultivosStmt, err := db.Prepare(ExtractCultivosQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                  db,
		insertCultivoStmt:   insertCultivoStmt,
		extractCultivosStmt: extractCultivosStmt,
	}, nil
}
