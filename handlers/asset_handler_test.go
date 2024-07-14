package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/zi-bot/simple-gin-rest/config"
	"github.com/zi-bot/simple-gin-rest/models"
	"github.com/zi-bot/simple-gin-rest/routes"
	"gorm.io/datatypes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var db = config.ConnectDatabaseTest()

func createTestApp() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(gin.Recovery())
	routes.SetupRoutes(r, db)
	return r
}

var app = createTestApp()

var asset = &models.Asset{
	Name:  "Test Asset",
	Type:  "Type asset",
	Value: 1000.5,
}

type ListResponse struct {
	Data       []*models.Asset    `json:"data"`
	Pagination *models.Pagination `json:"pagination"`
}

type Response struct {
	Data *models.Asset `json:"data"`
}

func insertOne() {
	jsonValue, _ := json.Marshal(asset)
	req := httptest.NewRequest(http.MethodPost, "/assets", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
}

func getOne() *models.Asset {
	req := httptest.NewRequest(http.MethodGet, "/assets/1", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	response := Response{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	return response.Data
}

func TestCreateUser(t *testing.T) {

	jsonValue, _ := json.Marshal(asset)
	req := httptest.NewRequest(http.MethodPost, "/assets", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllAssets(t *testing.T) {

	insertOne()

	req := httptest.NewRequest(http.MethodGet, "/assets?limit=10&?page=1", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	response := ListResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	for _, v := range response.Data {
		assert.Equal(t, v.Name, asset.Name)
	}

	assert.GreaterOrEqual(t, 10, len(response.Data))

}

func TestGetAssetById(t *testing.T) {
	insertOne()

	req := httptest.NewRequest(http.MethodGet, "/assets/1", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUser(t *testing.T) {
	insertOne()

	acquisitionDate := datatypes.Date(time.Now())
	asset.AcquisitionDate = &acquisitionDate

	jsonValue, _ := json.Marshal(asset)
	req := httptest.NewRequest(http.MethodPut, "/assets/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assetUpdated := getOne()
	assert.NotNil(t, assetUpdated)
}

func TestDeleteUser(t *testing.T) {
	insertOne()

	req := httptest.NewRequest(http.MethodDelete, "/assets/1", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	checkAsset := getOne()
	assert.Nil(t, checkAsset)
}
