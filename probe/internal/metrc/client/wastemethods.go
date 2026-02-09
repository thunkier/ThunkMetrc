package client

import (
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// GetWasteMethodsWasteMethods (WasteMethods)
// GET {{baseUrl}}/wastemethods/v2
func (f *Fetcher) GetWasteMethodsWasteMethods() ([]models.WasteMethod, error) {
	url := "{{baseUrl}}/wastemethods/v2"
	return fetchList[models.WasteMethod](f, "WasteMethods", "GetWasteMethodsWasteMethods", "GET", url, nil)
}
