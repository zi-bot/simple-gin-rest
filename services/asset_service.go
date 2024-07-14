package services

import (
	"context"

	"github.com/zi-bot/simple-gin-rest/models"
	"github.com/zi-bot/simple-gin-rest/repositories"
)

type AssetService interface {
	GetAllAssets(ctx context.Context, pagination *models.Pagination) ([]*models.Asset, error)
	GetAssetById(ctx context.Context, id uint64) (*models.Asset, error)
	CreateAsset(ctx context.Context, asset *models.Asset) error
	UpdateAsset(ctx context.Context, asset *models.Asset) error
	DeleteAsset(ctx context.Context, id uint64) error
}

type assetServiceImpl struct {
	repository repositories.AssetRepository
}

// CreateAsset implements AssetService.
func (s *assetServiceImpl) CreateAsset(ctx context.Context, asset *models.Asset) error {
	asset.ID = 0
	return s.repository.Save(ctx, asset)
}

// DeleteAsset implements AssetService.
func (s *assetServiceImpl) DeleteAsset(ctx context.Context, id uint64) (err error) {
	_, err = s.GetAssetById(ctx, id)
	if err != nil {
		return
	}
	return s.repository.Delete(ctx, id)
}

// GetAllAssets implements AssetService.
func (s *assetServiceImpl) GetAllAssets(ctx context.Context, pagination *models.Pagination) (assets []*models.Asset, err error) {

	// return s.repository.GetAllAsset(ctx)
	assets, pagination.Total, err = s.repository.GetAllAssetWithPagination(ctx, pagination.Page, pagination.Limit)
	return
}

// GetAssetById implements AssetService.
func (s *assetServiceImpl) GetAssetById(ctx context.Context, id uint64) (asset *models.Asset, err error) {
	asset, err = s.repository.GetAssetById(ctx, id)
	return
}

// UpdateAsset implements AssetService.
func (s *assetServiceImpl) UpdateAsset(ctx context.Context, asset *models.Asset) (err error) {
	_, err = s.GetAssetById(ctx, asset.ID)
	if err != nil {
		return
	}
	return s.repository.Save(ctx, asset)
}

func NewAssetService(repositories repositories.AssetRepository) AssetService {
	return &assetServiceImpl{
		repository: repositories,
	}
}
