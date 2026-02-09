package versions

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config represents the versions.yaml configuration.
type Config struct {
	SDKVersion   string                       `yaml:"sdk_version"`
	Dependencies map[string]map[string]string `yaml:"dependencies"`
}

// Load reads and parses versions.yaml.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read versions.yaml: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse versions.yaml: %w", err)
	}

	return &cfg, nil
}

// Save writes the config back to disk.
func (c *Config) Save(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Add header comment
	header := `# SDK and Dependency Version Configuration
# This file is the single source of truth for all versions across SDKs.
# Run ` + "`task gen:sdks`" + ` after modifying to regenerate all package manifests.

`
	if err := os.WriteFile(path, []byte(header+string(data)), 0644); err != nil {
		return fmt.Errorf("failed to write versions.yaml: %w", err)
	}

	return nil
}

// BumpSDKVersion increments the SDK version.
func (c *Config) BumpSDKVersion(part string) (string, error) {
	parts := strings.Split(c.SDKVersion, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid version format: %s", c.SDKVersion)
	}

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	switch part {
	case "major":
		major++
		minor = 0
		patch = 0
	case "minor":
		minor++
		patch = 0
	case "patch":
		patch++
	default:
		return "", fmt.Errorf("unknown version part: %s (use major, minor, or patch)", part)
	}

	c.SDKVersion = fmt.Sprintf("%d.%d.%d", major, minor, patch)
	return c.SDKVersion, nil
}

// GetDep returns a dependency version.
func (c *Config) GetDep(lang, dep string) string {
	if langDeps, ok := c.Dependencies[lang]; ok {
		if ver, ok := langDeps[dep]; ok {
			return ver
		}
	}
	return ""
}

// SetDep sets a dependency version.
func (c *Config) SetDep(lang, dep, version string) {
	if c.Dependencies == nil {
		c.Dependencies = make(map[string]map[string]string)
	}
	if c.Dependencies[lang] == nil {
		c.Dependencies[lang] = make(map[string]string)
	}
	c.Dependencies[lang][dep] = version
}
