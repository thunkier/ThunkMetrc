package app

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadFileConfig(configFile string) (FileConfig, error) {
	var cfg FileConfig
	paths := []string{configFile, "../" + configFile, "../../" + configFile}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			data, err := os.ReadFile(p)
			if err == nil {
				if err := json.Unmarshal(data, &cfg); err != nil {
					return cfg, fmt.Errorf("error parsing %s: %w", p, err)
				}
				fmt.Printf("Loaded configuration from %s\n", p)
				return cfg, nil
			}
		}
	}
	return cfg, fmt.Errorf("config file not found in paths: %v", paths)
}
