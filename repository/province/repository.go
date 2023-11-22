package province

import (
	"database/sql"

	_ "embed"

	"github.com/robertobouses/meimporta_unpepino/entity/province"
)

type REPOSITORYProvince interface {
	InsertProvinces(province province.Province) error
	ExtractProvincesName(name string) (province.Province, error)
}

type RepositoryProvince struct {
	db                       *sql.DB
	insertProvincesStmt      *sql.Stmt
	extractProvincesNameStmt *sql.Stmt
}

//go:embed sql/insert_provinces.sql
var InsertProvincesQuery string

//go:embed sql/extract_provinces_name.sql
var ExtractProvincesNameQuery string

func NewRepositoryProvince(db *sql.DB) (*RepositoryProvince, error) {
	insertProvincesStmt, err := db.Prepare(InsertProvincesQuery)
	if err != nil {
		return nil, err
	}
	extractProvincesNameStmt, err := db.Prepare(ExtractProvincesNameQuery)
	if err != nil {
		return nil, err
	}
	return &RepositoryProvince{
		db:                       db,
		insertProvincesStmt:      insertProvincesStmt,
		extractProvincesNameStmt: extractProvincesNameStmt,
	}, nil
}
