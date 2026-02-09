package utils

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper"
	"golang.org/x/sync/errgroup"
)

// TimeWindow helps calculate the correct LastModifiedStart and LastModifiedEnd parameters for Metrc delta syncs.
type TimeWindow struct {
	StartString string
	EndString   string
	StartTime   time.Time
	EndTime     time.Time
}

// GetTimeWindow calculates a TimeWindow based on the last known sync time.
func GetTimeWindow(lastKnownSync *time.Time, bufferMinutes int) TimeWindow {
	endTime := time.Now().UTC()
	startTime := endTime.Add(-24 * time.Hour)

	if lastKnownSync != nil {
		if bufferMinutes < 0 {
			bufferMinutes = 0
		}
		startTime = lastKnownSync.Add(-time.Duration(bufferMinutes) * time.Minute)
	}

	const metrcTimeFormat = "2006-01-02T15:04:05Z07:00"

	return TimeWindow{
		StartTime:   startTime,
		EndTime:     endTime,
		StartString: startTime.Format(metrcTimeFormat),
		EndString:   endTime.Format(metrcTimeFormat),
	}
}

// Collect serves as a generic helper for the Sync functions.
func Collect[T any](ctx context.Context, iterator func(handler func(T) error) error) ([]T, error) {
	var results []T
	err := iterator(func(item T) error {
		results = append(results, item)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}

// SynchronizationTarget binds a license number to a specific configured wrapper.
type SynchronizationTarget struct {
	LicenseNumber string
	Wrapper       *wrapper.MetrcWrapper
}

// ParallelSync executes a Metrc operation across multiple targets concurrently.
func ParallelSync[T any](
	ctx context.Context,
	targets []SynchronizationTarget,
	concurrencyLimit int,
	op func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]T, error),
) (map[string][]T, error) {

	if concurrencyLimit <= 0 {
		concurrencyLimit = 20
	}

	results := make(map[string][]T)
	var resultsMu sync.Mutex

	sem := make(chan struct{}, concurrencyLimit)
	g, ctx := errgroup.WithContext(ctx)

	for _, t := range targets {
		t := t // capture variable
		sem <- struct{}{}
		g.Go(func() error {
			defer func() { <-sem }()

			items, err := op(ctx, t.Wrapper, t.LicenseNumber)
			if err != nil {
				return fmt.Errorf("sync failed for %s: %w", t.LicenseNumber, err)
			}

			resultsMu.Lock()
			results[t.LicenseNumber] = items
			resultsMu.Unlock()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}

