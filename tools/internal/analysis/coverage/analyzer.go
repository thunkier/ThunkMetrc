package coverage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type CoverageSource string

const (
	SourceMissing  CoverageSource = "Missing"
	SourceCaptured CoverageSource = "Captured"
	SourceExample  CoverageSource = "Example"
	SourceInferred CoverageSource = "Inferred"
)

type EndpointStats struct {
	Group        string
	Name         string
	Source       CoverageSource
	ResponseType string
	IsPaginated  bool
	IsWrite      bool
}

type GroupStats struct {
	Name     string
	Total    int
	Covered  int
	Captured int
	Example  int
	Inferred int
}

type ManifestEntry struct {
	ResponseType string `json:"responseType"`
	SchemaFile   string `json:"schemaFile"`
	IsPaginated  bool   `json:"isPaginated"`
	IsWrite      bool   `json:"isWrite"`
}

func Run(manifestPath, responsesDir, outPath string) error {
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return fmt.Errorf("failed to read manifest: %w", err)
	}

	var manifest map[string]ManifestEntry
	if err := json.Unmarshal(data, &manifest); err != nil {
		return fmt.Errorf("failed to parse manifest: %w", err)
	}

	captured := scanResponses(responsesDir)

	endpoints := make([]EndpointStats, 0, len(manifest))
	groupStats := make(map[string]*GroupStats)
	modelGroups := make(map[string][]int)

	for key, entry := range manifest {
		parts := strings.Split(key, "/")
		if len(parts) != 2 {
			continue
		}
		group := parts[0]
		name := parts[1]

		source := SourceMissing
		if captured[key] {
			source = SourceCaptured
		} else if entry.SchemaFile != "" {
			source = SourceExample
		}

		idx := len(endpoints)
		endpoints = append(endpoints, EndpointStats{
			Group:        group,
			Name:         name,
			Source:       source,
			ResponseType: entry.ResponseType,
			IsPaginated:  entry.IsPaginated,
			IsWrite:      entry.IsWrite,
		})

		modelKey := fmt.Sprintf("%s/%s", group, entry.ResponseType)
		if entry.ResponseType != "" && entry.ResponseType != "WriteResponse" {
			modelGroups[modelKey] = append(modelGroups[modelKey], idx)
		}
	}

	for _, indices := range modelGroups {
		hasSource := false
		for _, idx := range indices {
			if endpoints[idx].Source == SourceCaptured || endpoints[idx].Source == SourceExample {
				hasSource = true
				break
			}
		}
		if hasSource {
			for _, idx := range indices {
				if endpoints[idx].Source == SourceMissing {
					endpoints[idx].Source = SourceInferred
				}
			}
		}
	}

	for _, ep := range endpoints {
		gStats, ok := groupStats[ep.Group]
		if !ok {
			gStats = &GroupStats{Name: ep.Group}
			groupStats[ep.Group] = gStats
		}
		gStats.Total++
		switch ep.Source {
		case SourceCaptured:
			gStats.Covered++
			gStats.Captured++
		case SourceExample:
			gStats.Covered++
			gStats.Example++
		case SourceInferred:
			gStats.Covered++
			gStats.Inferred++
		}
	}

	return generateReport(outPath, endpoints, groupStats)
}

func scanResponses(rootDir string) map[string]bool {
	captured := make(map[string]bool)
	filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(d.Name(), ".json") {
			return nil
		}
		rel, _ := filepath.Rel(rootDir, path)
		parts := strings.Split(filepath.ToSlash(rel), "/")
		if len(parts) >= 2 {
			group := parts[len(parts)-2]
			name := strings.TrimSuffix(parts[len(parts)-1], ".json")
			if group == "Tags" && strings.HasPrefix(name, "GetAvailableTag") {
				name = strings.Replace(name, "GetAvailableTag", "GetAvailable", 1)
			}
			captured[fmt.Sprintf("%s/%s", group, name)] = true
		}
		return nil
	})
	return captured
}

func generateReport(outPath string, endpoints []EndpointStats, groupMap map[string]*GroupStats) error {
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	totalEps := len(endpoints)
	var totalCovered, totalCaptured, totalExample, totalInferred int
	for _, g := range groupMap {
		totalCovered += g.Covered
		totalCaptured += g.Captured
		totalExample += g.Example
		totalInferred += g.Inferred
	}

	pct := 0.0
	if totalEps > 0 {
		pct = float64(totalCovered) / float64(totalEps) * 100
	}

	barLen := 20
	filledLen := int(pct / 5)
	bar := strings.Repeat("█", filledLen) + strings.Repeat("░", barLen-filledLen)

	fmt.Fprintf(f, "# 📊 Public SDK Coverage Report\n\n")
	fmt.Fprintf(f, "> **Coverage Sources:**\n")
	fmt.Fprintf(f, "> - 🟢 **Captured**: Verified against real API response\n")
	fmt.Fprintf(f, "> - 🔵 **Example**: Derived from Bruno documentation\n")
	fmt.Fprintf(f, "> - 🟣 **Inferred**: Auto-derived from sibling endpoints\n\n")

	fmt.Fprintf(f, "## 📈 Overall Coverage\n\n")
	fmt.Fprintf(f, "| Metric | Count | %%\n")
	fmt.Fprintf(f, "|--------|-------|---\n")
	fmt.Fprintf(f, "| **Total Endpoints** | %d | 100%%\n", totalEps)
	fmt.Fprintf(f, "| **Covered** | %d | %.1f%%\n", totalCovered, pct)
	fmt.Fprintf(f, "| 🟢 Captured | %d | %.1f%%\n", totalCaptured, float64(totalCaptured)/float64(totalEps)*100)
	fmt.Fprintf(f, "| 🔵 Example | %d | %.1f%%\n", totalExample, float64(totalExample)/float64(totalEps)*100)
	fmt.Fprintf(f, "| 🟣 Inferred | %d | %.1f%%\n", totalInferred, float64(totalInferred)/float64(totalEps)*100)

	fmt.Fprintf(f, "\n`%s` **%.1f%%**\n\n", bar, pct)

	fmt.Fprintf(f, "## 📋 Coverage by Group\n\n")
	fmt.Fprintf(f, "| Group | Total | Covered | Gap | Status |\n")
	fmt.Fprintf(f, "|-------|-------|---------|-----|--------|\n")

	var groups []string
	for k := range groupMap {
		groups = append(groups, k)
	}
	sort.Strings(groups)

	for _, g := range groups {
		stats := groupMap[g]
		gap := stats.Total - stats.Covered
		indicator := "🟥"
		if stats.Covered == stats.Total && stats.Total > 0 {
			indicator = "🟩"
		} else if stats.Covered > stats.Total/2 {
			indicator = "🟨"
		} else if stats.Covered > 0 {
			indicator = "🟧"
		}
		fmt.Fprintf(f, "| **%s** | %d | %d | %d | %s |\n", g, stats.Total, stats.Covered, gap, indicator)
	}

	fmt.Fprintf(f, "\n## 🕵️ Detailed Breakdown\n\n")

	for _, g := range groups {
		fmt.Fprintf(f, "### %s\n\n", g)
		fmt.Fprintf(f, "<details>\n<summary>Endpoints (%d/%d covered)</summary>\n\n", groupMap[g].Covered, groupMap[g].Total)
		fmt.Fprintf(f, "| Operation | Source | Model |\n")
		fmt.Fprintf(f, "|-----------|--------|-------|\n")

		var glist []EndpointStats
		for _, ep := range endpoints {
			if ep.Group == g {
				glist = append(glist, ep)
			}
		}
		sort.Slice(glist, func(i, j int) bool { return glist[i].Name < glist[j].Name })

		for _, ep := range glist {
			icon := "❌"
			srcLabel := "Missing"
			switch ep.Source {
			case SourceCaptured:
				icon = "🟢"
				srcLabel = "Captured"
			case SourceExample:
				icon = "🔵"
				srcLabel = "Example"
			case SourceInferred:
				icon = "🟣"
				srcLabel = "Inferred"
			}

			model := ep.ResponseType
			if model == "" {
				model = "-"
			}
			fmt.Fprintf(f, "| `%s` | %s %s | `%s` |\n", ep.Name, icon, srcLabel, model)
		}
		fmt.Fprintf(f, "\n</details>\n\n")
	}

	fmt.Fprintf(f, "---\n*Generated from manifest.json + response scanning*\n")
	return nil
}
