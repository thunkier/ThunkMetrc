package client

import (
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// GetFacilities (Facilities)
// GET {{baseUrl}}/facilities/v2
func (f *Fetcher) GetFacilities() ([]models.Facility, error) {
	url := "{{baseUrl}}/facilities/v2"
	return fetchList[models.Facility](f, "Facilities", "GetFacilities", "GET", url, nil)
}
