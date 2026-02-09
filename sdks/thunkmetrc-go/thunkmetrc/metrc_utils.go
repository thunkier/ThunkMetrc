package thunkmetrc

import (
	"context"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/thunkmetrc/utils"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper"
)

// ExecuteOnAllActiveFacilities executes a specified Metrc operation parallel across all active facilities accessible by the wrapper.
// This utility abstracts the process of fetching facilities, filtering for valid licenses, and managing concurrency.
func ExecuteOnAllActiveFacilities[T any](
	ctx context.Context,
	w *wrapper.MetrcWrapper,
	op func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]T, error),
	concurrencyLimit int,
) (map[string][]T, error) {
	// 1. Fetch Facilities
	facilities, err := w.Facilities.GetFacilities(ctx)
	if err != nil {
		return nil, err
	}

	// 2. Build Targets
	var targets []utils.SynchronizationTarget
	for _, f := range facilities {
		// Ensure facility has a license number
		if f.License.Number != "" {
			targets = append(targets, utils.SynchronizationTarget{
				LicenseNumber: f.License.Number,
				Wrapper:       w,
			})
		}
	}

	if len(targets) == 0 {
		return make(map[string][]T), nil
	}

	// 3. Parallel Sync
	return utils.ParallelSync(ctx, targets, concurrencyLimit, op)
}

// Chunk splits a slice into smaller slices of the specified size.
// Useful for batching requests (e.g. creating packages) to respect API payload limits.
func Chunk[T any](items []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(items); i += chunkSize {
		end := i + chunkSize
		if end > len(items) {
			end = len(items)
		}
		chunks = append(chunks, items[i:end])
	}
	return chunks
}

