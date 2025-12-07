package bruno

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ParseCollection parses all .bru files in the root directory.
func ParseCollection(rootDir string) (*Collection, error) {
	var requests []Request
	var errs []error

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".bru") {
			req, err := ParseRequest(path)
			if err != nil {
				errs = append(errs, fmt.Errorf("failed to parse %s: %w", path, err))
				return nil // Continue walking
			}
			requests = append(requests, *req)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(errs) > 0 {
		return &Collection{Requests: requests}, errors.Join(errs...)
	}

	return &Collection{Requests: requests}, nil
}

// ParseRequest parses a single .bru file.
func ParseRequest(path string) (*Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	req := &Request{
		FilePath: path,
		Group:    filepath.Base(filepath.Dir(path)),
		Params:   make(map[string]string),
	}

	scanner := bufio.NewScanner(file)
	var currentBlock string
	var braceBalance int

	for scanner.Scan() {
		rawLine := scanner.Text()
		line := strings.TrimSpace(rawLine)

		if currentBlock == "" && line == "" {
			continue
		}

		if currentBlock == "" {
			if strings.HasSuffix(line, "{") {
				blockName := strings.TrimSpace(strings.TrimSuffix(line, "{"))
				if blockName != "" {
					currentBlock = blockName
					braceBalance = 1

					if isHTTPMethod(strings.ToLower(currentBlock)) {
						req.Method = strings.ToUpper(currentBlock)
					}
					continue
				}
			}
			continue
		}

		open := strings.Count(rawLine, "{")
		close := strings.Count(rawLine, "}")
		braceBalance += (open - close)

		if braceBalance <= 0 {
			currentBlock = ""
			braceBalance = 0
			continue
		}

		switch {
		case currentBlock == "meta":
			if strings.HasPrefix(line, "name:") {
				req.Name = strings.TrimSpace(strings.TrimPrefix(line, "name:"))
			} else if strings.HasPrefix(line, "type:") {
				req.Type = strings.TrimSpace(strings.TrimPrefix(line, "type:"))
			} else if strings.HasPrefix(line, "seq:") {
				if i, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "seq:"))); err == nil {
					req.Seq = i
				}
			}

		case isHTTPMethod(strings.ToLower(currentBlock)):
			if strings.HasPrefix(line, "url:") {
				req.URL = strings.TrimSpace(strings.TrimPrefix(line, "url:"))
			} else if strings.HasPrefix(line, "body:") {
				req.BodyType = strings.TrimSpace(strings.TrimPrefix(line, "body:"))
			} else if strings.HasPrefix(line, "auth:") {
				req.AuthType = strings.TrimSpace(strings.TrimPrefix(line, "auth:"))
			}

		case currentBlock == "params:query":
			if idx := strings.Index(line, ":"); idx != -1 {
				key := strings.TrimSpace(line[:idx])
				val := strings.TrimSpace(line[idx+1:])
				req.Params[key] = val
			}

		case currentBlock == "body:json":
			req.Body += rawLine + "\n"

		case currentBlock == "docs":
			req.Docs += rawLine + "\n"
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return req, nil
}

func isHTTPMethod(m string) bool {
	switch m {
	case "get", "post", "put", "delete", "patch", "head", "options":
		return true
	}
	return false
}
