package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-bot/simple-gin-rest/handlers"
	"github.com/zi-bot/simple-gin-rest/repositories"
	"github.com/zi-bot/simple-gin-rest/services"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	assetRepo := repositories.NewAssetRepository(db)
	assetService := services.NewAssetService(assetRepo)
	assetHandler := handlers.NewAssetHandler(assetService)

	asset := r.Group("assets")
	{
		asset.GET("", assetHandler.ListAssets)
		asset.POST("", assetHandler.CreateAsset)
		asset.GET("/:id", assetHandler.DetailAsset)
		asset.PUT("/:id", assetHandler.UpdateAsset)
		asset.DELETE("/:id", assetHandler.DeleteAsset)
	}

}
