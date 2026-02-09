package bruno

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

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
				return nil
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

	if braceBalance != 0 {
		return nil, fmt.Errorf("balanced braces mismatch: %d open blocks remaining", braceBalance)
	}

	req.PathParams = extractPathParams(req.URL)
	req.Permissions = extractPermissions(req.Docs)
	req.DocParams = extractDocParams(req.Docs)

	return req, nil
}

func extractPathParams(url string) []string {
	re := regexp.MustCompile(`\{+([a-zA-Z0-9_]+)\}+`)
	matches := re.FindAllStringSubmatch(url, -1)

	unique := make(map[string]bool)
	var params []string

	ignore := map[string]bool{
		"baseUrl":  true,
		"username": true,
		"password": true,
	}

	for _, m := range matches {
		if len(m) > 1 {
			p := m[1]
			if !ignore[p] && !unique[p] {
				unique[p] = true
				params = append(params, p)
			}
		}
	}
	return params
}

func extractPermissions(content string) []string {
	re := regexp.MustCompile(`Permissions Required:\s*([\s\S]*?)(?:Example|$)`)
	matches := re.FindStringSubmatch(content)
	if len(matches) < 2 {
		return nil
	}

	raw := matches[1]
	var perms []string
	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.TrimPrefix(line, "- ")
		line = strings.TrimPrefix(line, "* ")
		line = strings.TrimSpace(line)
		if line != "" {
			perms = append(perms, line)
		}
	}
	return perms
}

func extractDocParams(docs string) []DocParam {
	var params []DocParam
	if !strings.Contains(docs, "Parameters:") {
		return params
	}

	lines := strings.Split(docs, "\n")
	inParams := false

	re := regexp.MustCompile(`^-\s+([a-zA-Z0-9_]+)\s+\((Required|Optional)\):\s*(.*)$`)

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "Parameters:") {
			inParams = true
			continue
		}
		if inParams {
			if trimmed == "" {
				continue
			}
			// If we hit another section (like Example Response), stop
			if strings.HasSuffix(trimmed, ":") && !strings.HasPrefix(trimmed, "-") {
				break
			}

			matches := re.FindStringSubmatch(trimmed)
			if len(matches) == 4 {
				p := DocParam{
					Name:        matches[1],
					Optional:    matches[2] == "Optional",
					Description: matches[3],
				}
				params = append(params, p)
			}
		}
	}
	return params
}

func isHTTPMethod(m string) bool {
	switch m {
	case "get", "post", "put", "delete", "patch", "head", "options":
		return true
	}
	return false
}
