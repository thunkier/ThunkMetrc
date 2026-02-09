package app

type ScraperConfig struct {
	SingularMap     map[string]string `json:"SingularMap"`
	CasingOverrides map[string]string `json:"CasingOverrides"`
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
	Automation          AutomationConfig  `json:"Automation"`
	PermissionOverrides map[string]string `json:"PermissionOverrides"`
	Scraper             ScraperConfig     `json:"Scraper"`
}
