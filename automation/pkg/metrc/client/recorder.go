package client

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/thunkmetrc/automation/internal/runner/common"
	"github.com/thunkmetrc/automation/pkg/bruno"
)

func SaveResponse(cfg common.Config, req *bruno.Request, data []byte) {
	if len(data) == 0 {
		return
	}

	fileName := filepath.Base(req.FilePath)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

	if fileName == "." || fileName == "" {
		fileName = strings.ReplaceAll(req.Name, "/", "_")
		fileName = strings.ReplaceAll(fileName, ":", "")
	}

	safeName := strings.TrimSpace(fileName)

	version := "unknown"
	if matches := regexp.MustCompile(`[/-](v\d+)(?:[/?]|$)`).FindStringSubmatch(req.URL); len(matches) > 1 {
		version = matches[1]
	}

	group := req.Group
	if group == "" {
		group = "Unsorted"
	}

	dir := filepath.Join(cfg.ResponsesDir, group, version)
	os.MkdirAll(dir, 0755)

	path := filepath.Join(dir, safeName+".json")
	os.WriteFile(path, data, 0644)
}
