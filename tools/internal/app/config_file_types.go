package app

type ScraperConfig struct {
	SingularMap           map[string]string `json:"SingularMap"`
	CasingOverrides       map[string]string `json:"CasingOverrides"`
	States                []string          `json:"States"`
	Concurrency           int               `json:"Concurrency"`
	RequestTimeoutSeconds int               `json:"RequestTimeoutSeconds"`
	MaxRetries            int               `json:"MaxRetries"`
	RetryBackoffMs        int               `json:"RetryBackoffMs"`
	UserAgent             string            `json:"UserAgent"`
}

type FileConfig struct {
	PermissionOverrides map[string]string `json:"PermissionOverrides"`
	Scraper             ScraperConfig     `json:"Scraper"`
}
