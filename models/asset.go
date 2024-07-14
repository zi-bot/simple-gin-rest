package models

import (
	"encoding/json"
	"time"

	"gorm.io/datatypes"
)

type Asset struct {
	ID              uint64          `gorm:"primaryKey" json:"id"`
	Name            string          `gorm:"size:50" json:"name"`
	Type            string          `gorm:"size:50" json:"type"`
	Value           float32         `json:"value"`
	AcquisitionDate *datatypes.Date `gorm:"type:date" json:"acquisition_date" time_format:"2006-01-02"`
}

func (a *Asset) UnmarshalJSON(data []byte) error {
	type Alias Asset
	aux := &struct {
		AcquisitionDate string `json:"acquisition_date"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.AcquisitionDate == "" {
		return nil
	}

	date, err := time.Parse("2006-01-02", aux.AcquisitionDate)
	if err != nil {
		return err
	}

	dt := datatypes.Date(date)

	a.AcquisitionDate = &dt
	return nil
}

type DetailRequest struct {
	ID uint64 `json:"id" param:"id"`
}

type Pagination struct {
	Limit int   `json:"limit" form:"limit"`
	Page  int   `json:"page" form:"page"`
	Total int64 `json:"total"`
}
