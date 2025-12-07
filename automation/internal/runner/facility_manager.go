package runner

import (
	"fmt"

	"github.com/thunkmetrc/automation/internal/runner/common"
	"github.com/thunkmetrc/automation/pkg/metrc/client"
	"github.com/thunkmetrc/automation/pkg/metrc/models"
)

func FetchFacilities(cfg common.Config) ([]models.Facility, error) {
	f := client.NewFetcher(cfg)

	_, _ = f.Facilities_GetAllV2()

	facilities, err := f.Facilities_GetAllV1()
	if err != nil {
		return nil, err
	}

	if len(facilities) == 0 && !cfg.DryRun {
		return nil, fmt.Errorf("no facilities found")
	}

	return facilities, nil
}
