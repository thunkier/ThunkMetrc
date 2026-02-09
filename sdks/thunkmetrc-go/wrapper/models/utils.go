package models

import (
	"context"
	"encoding/json"
	"fmt"
)

// DeserializePage attempts to convert an interface{} (from client) into a PaginatedResponse.
func DeserializePage[T any](input interface{}) (*PaginatedResponse[T], error) {
	// 1. Marshal back to JSON
	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input: %w", err)
	}

	// 2. Unmarshal into PaginatedResponse
	var resp PaginatedResponse[T]
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal to PaginatedResponse: %w", err)
	}

	// 3. Fallback: If Data is empty, check if input itself was an array (raw list response)
	if len(resp.Data) == 0 {
		var list []T
		if err := json.Unmarshal(bytes, &list); err == nil && len(list) > 0 {
			resp.Data = list
			resp.TotalPages = 1 // Assume single page
		}
	}

	return &resp, nil
}

// Ptr returns a pointer to the value passed in.
// This is a helper function to easily create pointers for optional fields in request structs.
func Ptr[T any](v T) *T {
	return &v
}

// IteratePages returns a function that yields items one by one from a paginated API.
// It handles fetching pages, buffering items, and checking for end-of-data conditions.
func IteratePages[T any](ctx context.Context, fetchPage func(page int) (*PaginatedResponse[T], error)) func() (T, error) {
	currentPage := 1
	var buffer []T
	var bufferIndex int
	done := false

	return func() (T, error) {
		var zero T
		
		// If we have items in the buffer, return next one
		if bufferIndex < len(buffer) {
			item := buffer[bufferIndex]
			bufferIndex++
			return item, nil
		}

		if done {
			return zero, nil // End of iteration
		}

		// Fetch next page
		// Loop until we get data or finish
		for {
			if ctx.Err() != nil {
				return zero, ctx.Err()
			}

			response, err := fetchPage(currentPage)
			if err != nil {
				return zero, err
			}

			if response == nil || len(response.Data) == 0 {
				done = true
				return zero, nil
			}

			// got data
			buffer = response.Data
			bufferIndex = 0
			currentPage++

			// confirm we actually have data to return
			if len(buffer) > 0 {
				item := buffer[bufferIndex]
				bufferIndex++
				
				// check total pages to see if we should stop after this buffer
				if response.TotalPages > 0 && currentPage > response.TotalPages {
					done = true
				}
				// safety check for single-page responses
				if response.TotalPages == 0 {
					done = true
				}

				return item, nil
			}
			
			// response had data structure but empty list? try next page or stop?
			// if totalpages says we are done, stop.
			if response.TotalPages > 0 && currentPage > response.TotalPages {
				done = true
				return zero, nil
			}
			
			// otherwise loop to try next page
		}
	}
}
