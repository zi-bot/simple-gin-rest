package repositories

import (
	"context"

	"github.com/zi-bot/simple-gin-rest/models"
	"gorm.io/gorm"
)

type AssetRepository interface {
	GetAllAsset(ctx context.Context) ([]*models.Asset, error)
	GetAssetById(ctx context.Context, id uint64) (*models.Asset, error)
	Save(ctx context.Context, asset *models.Asset) error
	Delete(ctx context.Context, id uint64) error
	GetAllAssetWithPagination(ctx context.Context, page, limit int) ([]*models.Asset, int64, error)
}

type assetRepositoryImpl struct {
	db *gorm.DB
}

// GetAllAssetWithPagination implements AssetRepository.
func (r *assetRepositoryImpl) GetAllAssetWithPagination(ctx context.Context, page, limit int) ([]*models.Asset, int64, error) {
	var totalRecords int64
	var assets []*models.Asset
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	err := r.db.WithContext(ctx).Session(&gorm.Session{QueryFields: true}).Limit(limit).Offset(offset).Find(&assets).Count(&totalRecords).Error
	return assets, totalRecords, err
}

// Delete implements AssetRepository.
func (r *assetRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&models.Asset{}, id).Error
}

// GetAllAsset implements AssetRepository.
func (r *assetRepositoryImpl) GetAllAsset(ctx context.Context) (assets []*models.Asset, err error) {
	err = r.db.WithContext(ctx).Find(&assets).Error
	return
}

// GetAssetById implements AssetRepository.
func (r *assetRepositoryImpl) GetAssetById(ctx context.Context, id uint64) (asset *models.Asset, err error) {
	err = r.db.WithContext(ctx).Session(&gorm.Session{QueryFields: true}).First(&asset, "id = ?", id).Error
	return
}

// Save implements AssetRepository.
func (r *assetRepositoryImpl) Save(ctx context.Context, asset *models.Asset) error {
	return r.db.WithContext(ctx).Save(&asset).Error
}

func NewAssetRepository(db *gorm.DB) AssetRepository {
	return &assetRepositoryImpl{
		db: db,
	}
}
