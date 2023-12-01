package repository

import (
	"database/sql"

	_ "embed"

	"github.com/robertobouses/meimporta_unpepino/entity"
	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
)

type REPOSITORY interface {
	InsertCrops(crop entity.Crop) error
	ExtractCrops() ([]entity.Crop, error)
	DeleteCropsAll() error
	DeleteCropsId(id int) error
	CheckExistCropsId(id int) (bool, error)
	ExtractCropsName(name string) (entity.CalculateResponse, error)
	ExtractCropsSearch(request entity.SearchRequest) ([]entity.Crop, error)
	ExtractCropsCalendary(month, climate string) ([]entity.Crop, error)
	InsertFields(field mycrop.MyField) error
}

type Repository struct {
	db                          *sql.DB
	insertCropsStmt             *sql.Stmt
	insertCropsInformationStmt  *sql.Stmt
	insertCropsRequirementsStmt *sql.Stmt
	insertCropsDatesStmt        *sql.Stmt
	insertCropsFruitStmt        *sql.Stmt
	insertCropsSeedStmt         *sql.Stmt
	insertCropsIssuesStmt       *sql.Stmt
	deleteCropsAllStmt          *sql.Stmt
	deleteCropsIdStmt           *sql.Stmt

	checkExistCropsIdStmt *sql.Stmt

	extractCropsStmt          *sql.Stmt
	extractCropsNameStmt      *sql.Stmt
	extractCropsSearchStmt    *sql.Stmt
	extractCropsCalendaryStmt *sql.Stmt
	insertFieldsStmt          *sql.Stmt
}

//go:embed sql/insert_crops.sql
var InsertCropsQuery string

//go:embed sql/insert_crops_information.sql
var InsertCropsInformationQuery string

//go:embed sql/insert_crops_requirements.sql
var InsertCropsRequirementsQuery string

//go:embed sql/insert_crops_dates.sql
var InsertCropsDatesQuery string

//go:embed sql/insert_crops_fruit.sql
var InsertCropsFruitQuery string

//go:embed sql/insert_crops_seed.sql
var InsertCropsSeedQuery string

//go:embed sql/insert_crops_issues.sql
var InsertCropsIssuesQuery string

//go:embed sql/extract_crops.sql
var ExtractCropsQuery string

//go:embed sql/extract_crops_name.sql
var ExtractCropsNameQuery string

//go:embed sql/delete_crops_all.sql
var DeleteCropsAllQuery string

//go:embed sql/check_exist_crops_id.sql
var CheckExistCropsIdQuery string

//go:embed sql/delete_crops_id.sql
var DeleteCropsIdQuery string

//go:embed sql/extract_crops_search.sql
var ExtractCropsSearchQuery string

//go:embed sql/extract_crops_calendary.sql
var ExtractCropsCalendaryQuery string

//go:embed sql/insert_fields.sql
var insertFieldsQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	insertCropsStmt, err := db.Prepare(InsertCropsQuery)
	if err != nil {
		return nil, err
	}

	insertCropsInformationStmt, err := db.Prepare(InsertCropsInformationQuery)
	if err != nil {
		return nil, err
	}

	insertCropsRequirementsStmt, err := db.Prepare(InsertCropsRequirementsQuery)
	if err != nil {
		return nil, err
	}

	insertCropsDatesStmt, err := db.Prepare(InsertCropsDatesQuery)
	if err != nil {
		return nil, err
	}

	insertCropsFruitStmt, err := db.Prepare(InsertCropsFruitQuery)
	if err != nil {
		return nil, err
	}

	insertCropsSeedStmt, err := db.Prepare(InsertCropsSeedQuery)
	if err != nil {
		return nil, err
	}

	insertCropsIssuesStmt, err := db.Prepare(InsertCropsIssuesQuery)
	if err != nil {
		return nil, err
	}

	extractCropsStmt, err := db.Prepare(ExtractCropsQuery)
	if err != nil {
		return nil, err
	}
	extractCropsNameStmt, err := db.Prepare(ExtractCropsNameQuery)
	if err != nil {
		return nil, err
	}
	deleteCropsAllStmt, err := db.Prepare(DeleteCropsAllQuery)
	if err != nil {
		return nil, err
	}
	deleteCropsIdStmt, err := db.Prepare(DeleteCropsIdQuery)
	if err != nil {
		return nil, err
	}
	checkExistCropsIdStmt, err := db.Prepare(CheckExistCropsIdQuery)
	if err != nil {
		return nil, err
	}
	extractCropsSearchStmt, err := db.Prepare(ExtractCropsSearchQuery)
	if err != nil {
		return nil, err
	}
	extractCropsCalendaryStmt, err := db.Prepare(ExtractCropsCalendaryQuery)
	if err != nil {
		return nil, err
	}

	insertFieldsStmt, err := db.Prepare(insertFieldsQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                          db,
		insertCropsStmt:             insertCropsStmt,
		insertCropsInformationStmt:  insertCropsInformationStmt,
		insertCropsRequirementsStmt: insertCropsRequirementsStmt,
		insertCropsDatesStmt:        insertCropsDatesStmt,
		insertCropsFruitStmt:        insertCropsFruitStmt,
		insertCropsSeedStmt:         insertCropsSeedStmt,
		insertCropsIssuesStmt:       insertCropsIssuesStmt,
		extractCropsStmt:            extractCropsStmt,
		extractCropsNameStmt:        extractCropsNameStmt,
		deleteCropsAllStmt:          deleteCropsAllStmt,
		deleteCropsIdStmt:           deleteCropsIdStmt,
		checkExistCropsIdStmt:       checkExistCropsIdStmt,
		extractCropsSearchStmt:      extractCropsSearchStmt,
		extractCropsCalendaryStmt:   extractCropsCalendaryStmt,
		insertFieldsStmt:            insertFieldsStmt,
	}, nil
}
