package repository

import (
	"database/sql"

	_ "embed"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

type REPOSITORY interface {
	InsertCrop(crop entity.Crop) error
	ExtractCrops() ([]entity.Crop, error)
	DeleteCropsAll() error
	DeleteCropsId(id int) error
	CheckExistCropId(id int) (bool, error)
	ExtractCropName(name string) (entity.CalculateResponse, error)
	ExtractSearchCrop(request entity.SearchRequest) ([]entity.Crop, error)
}

type Repository struct {
	db                         *sql.DB
	insertCropStmt             *sql.Stmt
	insertCropInformation Stmt  *sql.Stmt
	insertCropRequirements Stmt *sql.Stmt
	insertCropDatesStmt        *sql.Stmt
	insertCropFruitStmt        *sql.Stmt
	insertCropSeedStmt         *sql.Stmt
	insertCropIssuesStmt    *sql.Stmt
	deleteCropsAllStmt         *sql.Stmt
	deleteCropsIdStmt          *sql.Stmt

	checkExistCropIdStmt *sql.Stmt

	extractCropsStmt      *sql.Stmt
	extractCropNameStmt   *sql.Stmt
	extractSearchCropStmt *sql.Stmt
}

//go:embed sql/insert_crop.sql
var InsertCropQuery string

//go:embed sql/insert_crop_information.sql
var InsertCropInformation Query string

//go:embed sql/insert_crop_requirements.sql
var InsertCropRequirements Query string

//go:embed sql/insert_crop_dates.sql
var InsertCropDatesQuery string

//go:embed sql/insert_crop_fruit.sql
var InsertCropFruitQuery string

//go:embed sql/insert_crop_seed.sql
var InsertCropSeedQuery string

//go:embed sql/insert_crop_problemas.sql
var InsertCropIssuesQuery string

//go:embed sql/extract_crops.sql
var ExtractCropsQuery string

//go:embed sql/extract_crop_name.sql
var ExtractCropNameQuery string

//go:embed sql/delete_crops_all.sql
var DeleteCropsAllQuery string

//go:embed sql/check_exist_crop_id.sql
var CheckExistCropIdQuery string

//go:embed sql/delete_crops_id.sql
var DeleteCropsIdQuery string

//go:embed sql/extract_crops_search.sql
var ExtractSearchCropsQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	insertCropStmt, err := db.Prepare(InsertCropQuery)
	if err != nil {
		return nil, err
	}

	insertCropInformation Stmt, err := db.Prepare(InsertCropInformation Query)
	if err != nil {
		return nil, err
	}

	insertCropRequirements Stmt, err := db.Prepare(InsertCropRequirements Query)
	if err != nil {
		return nil, err
	}

	insertCropDatesStmt, err := db.Prepare(InsertCropDatesQuery)
	if err != nil {
		return nil, err
	}

	insertCropFruitStmt, err := db.Prepare(InsertCropFruitQuery)
	if err != nil {
		return nil, err
	}

	insertCropSeedStmt, err := db.Prepare(InsertCropSeedQuery)
	if err != nil {
		return nil, err
	}

	insertCropIssuesStmt, err := db.Prepare(InsertCropIssuesQuery)
	if err != nil {
		return nil, err
	}

	extractCropsStmt, err := db.Prepare(ExtractCropsQuery)
	if err != nil {
		return nil, err
	}
	extractCropNameStmt, err := db.Prepare(ExtractCropNameQuery)
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
	checkExistCropIdStmt, err := db.Prepare(CheckExistCropIdQuery)
	if err != nil {
		return nil, err
	}
	extractSearchCropStmt, err := db.Prepare(ExtractSearchCropsQuery)
	if err != nil {
		return nil, err
	}
	return &Repository{
		db:                         db,
		insertCropStmt:             insertCropStmt,
		insertCropInformation Stmt:  insertCropInformation Stmt,
		insertCropRequirements Stmt: insertCropRequirements Stmt,
		insertCropDatesStmt:        insertCropDatesStmt,
		insertCropFruitStmt:        insertCropFruitStmt,
		insertCropSeedStmt:         insertCropSeedStmt,
		insertCropIssuesStmt:    insertCropIssuesStmt,
		extractCropsStmt:           extractCropsStmt,
		extractCropNameStmt:        extractCropNameStmt,
		deleteCropsAllStmt:         deleteCropsAllStmt,
		deleteCropsIdStmt:          deleteCropsIdStmt,
		checkExistCropIdStmt:       checkExistCropIdStmt,
		extractSearchCropStmt:      extractSearchCropStmt,
	}, nil
}
