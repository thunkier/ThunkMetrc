package app

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/thunkier/thunkmetrc/probe/internal/metrc/client"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

func (c Config) ToClientConfig() client.Config {
	var canUseLabel bool
	if c.Facility != nil {
		canUseLabel = c.Facility.FacilityType.CanUseComplianceLabel
	}

	return client.Config{
		BaseURL:               c.BaseURL,
		APIKey:                c.APIKey,
		UserKey:               c.UserKey,
		FacilityLicenseNumber: c.FacilityLicenseNumber,
		DryRun:                c.DryRun,
		Logger:                c.Logger,
		ErrorReporter:         c.ErrorAggregator,
		UsedRequests:          c.UsedRequests,
		CurrentWorkflow:       c.CurrentWorkflow,
		BlockedEndpoints:      c.ExternalConfig.BlockedEndpoints,
		CanUseComplianceLabel: canUseLabel,
		ResponsesDir:          c.ResponsesDir,
		SuppressConsoleLog:    c.Dashboard != nil,
	}
}

type Config struct {
	BaseURL               string
	APIKey                string
	UserKey               string
	SpecsDir              string
	ResponsesDir          string
	DryRun                bool
	ManualStep            bool
	UseMockPermissions    bool
	FacilityLicenseNumber string
	Facility              *models.Facility
	Logger                Logger
	UsedRequests          *sync.Map
	DependencyCache       *sync.Map
	ExternalConfig        AutomationConfig
	EmployeePermissions   []string
	PermissionOverrides   map[string]string
	Scraper               ScraperConfig
	Cleanup               *CleanupManager
	ErrorAggregator       *ErrorAggregator
	CurrentWorkflow       string
	Dashboard             *Dashboard
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err == nil {
		fmt.Println("Loaded .env from current directory")
	} else if err := godotenv.Load("../.env"); err == nil {
		fmt.Println("Loaded .env from parent directory")
	} else {
		fmt.Printf("Notice: Could not load .env (checked current and parent directories): %v\n", err)
	}

	dryKey := os.Getenv("DRY_RUN")
	dryRun := dryKey != "false"

	baseUrl := os.Getenv("METRC_BASE_URL")
	if dryRun {
		baseUrl = "http://localhost:8080"
		fmt.Println("DryRun active: Using Mock Server at http://localhost:8080")
	} else if baseUrl == "" {
		fmt.Println("WARNING: METRC_BASE_URL is not set. Requests will fail with unsupported protocol scheme.")
	}

	fileCfg, err := LoadFileConfig("config.json")
	if err != nil {
		fmt.Printf("Notice: Failed to load config: %v. Using defaults.\n", err)
	}

	if val := os.Getenv("METRC_EMPLOYEE_LICENSE_NUMBER"); val != "" {
		fileCfg.Automation.SpecialEmployeeUserKey = val
	}

	if fileCfg.PermissionOverrides == nil {
		fileCfg.PermissionOverrides = make(map[string]string)
	}

	return Config{
		BaseURL:             baseUrl,
		APIKey:              getEnv("METRC_VENDOR_API_KEY", ""),
		UserKey:             getEnv("METRC_USER_API_KEY", ""),
		SpecsDir:            "../specs/generated/bruno",
		ResponsesDir:        getEnv("METRC_RESPONSES_DIR", "../specs/source/captured/responses"),
		DryRun:              dryRun,
		UsedRequests:        &sync.Map{},
		DependencyCache:     &sync.Map{},
		ExternalConfig:      fileCfg.Automation,
		EmployeePermissions: []string{},
		PermissionOverrides: fileCfg.PermissionOverrides,
		Scraper:             fileCfg.Scraper,
		Cleanup:             &CleanupManager{},
		ErrorAggregator:     NewErrorAggregator(),
	}
}

type CleanupManager struct {
	mu    sync.Mutex
	Tasks []func()
}

func (m *CleanupManager) Add(task func()) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Tasks = append(m.Tasks, task)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
