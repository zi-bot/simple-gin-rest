package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mocksRepo "github.com/zi-bot/simple-gin-rest/mocks/repositories"
	"github.com/zi-bot/simple-gin-rest/models"
)

var repo = mocksRepo.AssetRepository{Mock: mock.Mock{}}
var service = NewAssetService(&repo)

var asset = &models.Asset{Name: "Chair", Type: "Furniture", Value: 1000.5}

func TestCreateAsset(t *testing.T) {
	ctx := context.Background()

	repo.Mock.On("Save", ctx, asset).Return(nil)
	err := service.CreateAsset(ctx, asset)

	assert.NoError(t, err)

}

func TestGetListAsset(t *testing.T) {
	ctx := context.Background()
	pagination := &models.Pagination{}

	repo.Mock.On("GetAllAssetWithPagination", ctx, pagination.Page, pagination.Limit).Return([]*models.Asset{asset}, int64(10), nil)

	listAssets, err := service.GetAllAssets(ctx, pagination)
	assert.NoError(t, err)

	assert.GreaterOrEqual(t, 1, len(listAssets))
}
