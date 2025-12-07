package common

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/thunkmetrc/automation/pkg/metrc/models"
	"github.com/thunkmetrc/automation/pkg/scraper"
)

type Config struct {
	BaseURL               string
	APIKey                string
	SpecsDir              string
	ResponsesDir          string
	DryRun                bool
	ManualStep            bool
	FacilityLicenseNumber string
	Facility              *models.Facility
	Logger                Logger
	UsedRequests          *sync.Map
	DependencyCache       *sync.Map
	ExternalConfig        AutomationConfig
	EmployeePermissions   []string
	PermissionOverrides   map[string]string
	Scraper               scraper.ScraperConfig
}

type SanitizeRule struct {
	Pattern string   `json:"Pattern"`
	Fields  []string `json:"Fields"`
}

type AutomationConfig struct {
	CaregiverLicenseNumber string         `json:"CaregiverLicenseNumber"`
	SpecialEmployeeUserKey string         `json:"SpecialEmployeeUserKey"`
	SanitizeRules          []SanitizeRule `json:"SanitizeRules"`
	BlockedEndpoints       []string       `json:"BlockedEndpoints"`
}

type FileConfig struct {
	Automation          AutomationConfig      `json:"Automation"`
	PermissionOverrides map[string]string     `json:"PermissionOverrides"`
	Scraper             scraper.ScraperConfig `json:"Scraper"`
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Notice: Could not load .env from current directory: %v\n", err)
		if err := godotenv.Load("../.env"); err != nil {
			fmt.Printf("Notice: Could not load .env from parent directory: %v\n", err)
		} else {
			fmt.Println("Loaded .env from parent directory")
		}
	} else {
		fmt.Println("Loaded .env from current directory")
	}

	dryKey := os.Getenv("DRY_RUN")
	dryRun := dryKey != "false"

	baseUrl := os.Getenv("METRC_BASE_URL")
	if baseUrl == "" && !dryRun {
		fmt.Println("WARNING: METRC_BASE_URL is not set. Requests will fail with unsupported protocol scheme.")
	}

	var fileCfg FileConfig
	configFile := "config.json"

	if _, err := os.Stat(configFile); err == nil {
		data, err := os.ReadFile(configFile)
		if err == nil {
			_ = json.Unmarshal(data, &fileCfg)
			fmt.Printf("Loaded configuration from %s\n", configFile)
		} else {
			fmt.Printf("Notice: Could not read %s: %v\n", configFile, err)
		}
	} else if _, err := os.Stat("../" + configFile); err == nil {
		data, err := os.ReadFile("../" + configFile)
		if err == nil {
			_ = json.Unmarshal(data, &fileCfg)
			fmt.Printf("Loaded configuration from ../%s\n", configFile)
		}
	} else {
		fmt.Printf("Notice: %s not found. Using defaults.\n", configFile)
	}

	if val := os.Getenv("METRC_SPECIAL_EMPLOYEE_USER_KEY"); val != "" {
		fileCfg.Automation.SpecialEmployeeUserKey = val
	}

	if fileCfg.PermissionOverrides == nil {
		fileCfg.PermissionOverrides = make(map[string]string)
	}

	return Config{
		BaseURL:             baseUrl,
		APIKey:              getEnv("METRC_VENDOR_API_KEY", ""),
		SpecsDir:            "../specs/metrc-bruno",
		ResponsesDir:        getEnv("METRC_RESPONSES_DIR", "../specs/metrc-responses"),
		DryRun:              dryRun,
		UsedRequests:        &sync.Map{},
		DependencyCache:     &sync.Map{},
		ExternalConfig:      fileCfg.Automation,
		EmployeePermissions: []string{},
		PermissionOverrides: fileCfg.PermissionOverrides,
		Scraper:             fileCfg.Scraper,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
