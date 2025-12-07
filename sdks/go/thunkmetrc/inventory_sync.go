package thunkmetrc

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/thunkmetrc/sdks/go/wrapper"
)

type ThunkMetrc struct {
	Wrapper *wrapper.MetrcWrapper
}

func New(w *wrapper.MetrcWrapper) *ThunkMetrc {
	return &ThunkMetrc{Wrapper: w}
}

// ActivePackagesInventorySync fetches all active packages (paginated) modified within the time window.
// If lastKnownSync is provided, start = lastKnownSync - buffer.
// If lastKnownSync is nil, start = 24 hours ago (default safety).
func (tm *ThunkMetrc) ActivePackagesInventorySync(ctx context.Context, licenseNumber string, lastKnownSync *time.Time, bufferMinutes int) ([]map[string]interface{}, error) {
	endTime := time.Now().UTC()
	var startTime time.Time

	if lastKnownSync != nil {
		startTime = lastKnownSync.Add(-time.Duration(bufferMinutes) * time.Minute).UTC()
	} else {
		// Default to 24 hours ago if no sync state
		startTime = endTime.Add(-24 * time.Hour).UTC()
	}

	// Format as ISO 8601 with timezone (required by Metrc)
	// Go's RFC3339 is ISO 8601 compatible: "2006-01-02T15:04:05Z07:00"
	timeLayout := "2006-01-02T15:04:05-07:00" // Use explicit offset format if needed, or RFC3339
	// Metrc strictly needs the offset for proper parsing if '+' is involved (now handled by url encoding fix)
	startStr := startTime.Format(timeLayout)
	endStr := endTime.Format(timeLayout)

	var allPackages []map[string]interface{}
	page := 1
	pageSize := 20 // Max page size usually

	for {
		// Call Wrapper
		// Signature assumed: (lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
		// Based on sorted keys: lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
		resp, err := tm.Wrapper.PackagesGetActiveV2(
			endStr,
			startStr,
			licenseNumber,
			strconv.Itoa(page),
			strconv.Itoa(pageSize),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get active packages (page %d): %w", page, err)
		}

		// Handle Response (interface{} -> map)
		// The generic return is map[string]interface{} for JSON object
		respMap, ok := resp.(map[string]interface{})
		if !ok {
			// Maybe it returned nil or something else?
			// If we used json.Decoder, it should be map[string]interface{} for object
			// Verify if resp is already map
			// If "body: json" in result, it might be.
			// Re-marshal to be safe? No, client returns map.
			// Let's debug if type assertion fails.
			// For now assume ok.
			return nil, fmt.Errorf("unexpected response type: %T", resp)
		}

		// Pagination Metadata
		// "TotalPages": 1
		totalPagesVal, _ := respMap["TotalPages"].(float64)
		totalPages := int(totalPagesVal)

		// Data
		data, _ := respMap["Data"].([]interface{})
		for _, item := range data {
			if pkgMap, ok := item.(map[string]interface{}); ok {
				allPackages = append(allPackages, pkgMap)
			}
		}

		if page >= totalPages {
			break
		}
		page++
	}

	return allPackages, nil
}
