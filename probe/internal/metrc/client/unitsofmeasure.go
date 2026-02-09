package client

import (
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// GetActiveUnitsOfMeasure (UnitsOfMeasure)
// GET {{baseUrl}}/unitsofmeasure/v2/active
func (f *Fetcher) GetActiveUnitsOfMeasure() ([]models.UnitOfMeasure, error) {
	url := "{{baseUrl}}/unitsofmeasure/v2/active"
	return fetchList[models.UnitOfMeasure](f, "UnitsOfMeasure", "GetActiveUnitsOfMeasure", "GET", url, nil)
}

// GetInactiveUnitsOfMeasure (UnitsOfMeasure)
// GET {{baseUrl}}/unitsofmeasure/v2/inactive
func (f *Fetcher) GetInactiveUnitsOfMeasure() ([]models.UnitOfMeasure, error) {
	url := "{{baseUrl}}/unitsofmeasure/v2/inactive"
	return fetchList[models.UnitOfMeasure](f, "UnitsOfMeasure", "GetInactiveUnitsOfMeasure", "GET", url, nil)
}
