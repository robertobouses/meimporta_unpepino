package repository

import (
	"database/sql"

	_ "embed"

	"github.com/robertobouses/meimporta_unpepino/entity/province"
)

type REPOSITORYProvince interface {
	InsertProvinces(province province.Province) error
}

type RepositoryProvince struct {
	db              *sql.DB
	insertCropsStmt *sql.Stmt
}

//go:embed sql/insert_provinces.sql
var InsertProvincesQuery string

func NewRepository(db *sql.DB) (*RepositoryProvince, error) {
	insertProvincesStmt, err := db.Prepare(InsertProvincesQuery)
	if err != nil {
		return nil, err
	}

	return &RepositoryProvince{
		db:                  db,
		insertProvincesStmt: insertProvincesStmt,
	}, nil
}
