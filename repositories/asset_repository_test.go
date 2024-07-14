package repositories

import (
	"context"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zi-bot/simple-gin-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

var DB, mock = NewMockDB()

var assetRepo = NewAssetRepository(DB)

func TestCreateAsset(t *testing.T) {

	t.Skip()

	asset := &models.Asset{Name: "Chair", Type: "Furniture", Value: 1000.5}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "assets" ("name","type","value","acquisition_date") VALUES ($1,$2,$3,$4) RETURNING "id"`)).
		WithArgs(asset.Name, asset.Type, asset.Value, nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := assetRepo.Save(context.Background(), asset)

	assert.NoError(t, err)

}
