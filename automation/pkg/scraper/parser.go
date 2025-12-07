package scraper

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ParsedOperation struct {
	Group       string
	Name        string
	Method      string
	URL         string
	Description string
	Params      []Param
	Body        string
	Response    string
	Permissions []string
}

type Param struct {
	Name        string
	Description string
	Value       string
}

// ScraperConfig holds configuration for the scraper.
type ScraperConfig struct {
	SingularMap     map[string]string `json:"SingularMap"`
	CasingOverrides map[string]string `json:"CasingOverrides"`
}

// ParseDocumentation parses the HTML content and returns a list of operations.
func ParseDocumentation(r io.Reader, config ScraperConfig) ([]ParsedOperation, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var operations []ParsedOperation
	currentTag := "Default"

	doc.Find("h1, h2, h3").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		tagName := goquery.NodeName(s)

		if (tagName == "h1" || tagName == "h2") && text != "" && !strings.Contains(text, "Server:") {
			currentTag = text
			return
		}

		if tagName == "h3" {
			re := regexp.MustCompile(`(?i)^(GET|POST|PUT|DELETE)\s+(/.*)`)
			match := re.FindStringSubmatch(text)
			if len(match) < 3 {
				return
			}

			rawMethod := match[1]
			rawURL := strings.TrimSpace(match[2])
			if parts := strings.Fields(rawURL); len(parts) > 0 {
				rawURL = parts[0]
			}
			rawURL = strings.TrimRight(rawURL, "/")

			op := ParsedOperation{
				Group:  currentTag,
				Method: strings.ToUpper(rawMethod),
				URL:    rawURL,
			}

			if next := s.Next(); goquery.NodeName(next) == "p" {
				op.Description = strings.TrimSpace(next.Text())
			}

			s.NextAll().EachWithBreak(func(_ int, sibling *goquery.Selection) bool {
				name := goquery.NodeName(sibling)
				if name == "h1" || name == "h2" || name == "h3" {
					return false
				}

				if name == "p" && op.Description == "" {
					op.Description = strings.TrimSpace(sibling.Text())
				}

				if name == "h4" {
					hText := sibling.Text()
					if strings.Contains(hText, "Permissions Required") {
						parsePermissions(sibling, &op)
					} else if strings.Contains(hText, "Parameters") {
						parseParameters(sibling, &op)
					} else if strings.Contains(hText, "Example Request") {
						parseRequestBody(sibling, &op)
					} else if strings.Contains(hText, "Example Response") {
						parseResponseBody(sibling, &op)
					}
				}
				return true
			})

			if !isDuplicate(operations, op) {
				op.Name = generateUniqueName(operations, op, config)
				operations = append(operations, op)
			}
		}
	})

	return operations, nil
}

func parsePermissions(s *goquery.Selection, op *ParsedOperation) {
	if table := s.Next(); goquery.NodeName(table) == "table" {
		table.Find("tr td:first-child").Each(func(_ int, cell *goquery.Selection) {
			op.Permissions = append(op.Permissions, strings.TrimSpace(cell.Text()))
		})
	}
}

func parseParameters(s *goquery.Selection, op *ParsedOperation) {
	if table := s.Next(); goquery.NodeName(table) == "table" {
		table.Find("tr").Each(func(_ int, row *goquery.Selection) {
			cols := row.Find("td")
			if cols.Length() >= 2 {
				pName := strings.Fields(cols.Eq(0).Text())[0]
				pDesc := strings.TrimSpace(cols.Eq(1).Text())
				op.Params = append(op.Params, Param{Name: pName, Description: pDesc})
			}
		})
	}
}

func parseRequestBody(s *goquery.Selection, op *ParsedOperation) {
	if pre := s.Next(); goquery.NodeName(pre) == "pre" {
		text := pre.Text()
		trimmed := strings.TrimSpace(text)
		if regexp.MustCompile(`(?i)^[A-Z]+\s+/`).MatchString(trimmed) {
			// Try to find start of JSON content
			idxObj := strings.Index(text, "{")
			idxArr := strings.Index(text, "[")
			start := -1
			if idxObj > -1 && (idxArr == -1 || idxObj < idxArr) {
				start = idxObj
			} else if idxArr > -1 {
				start = idxArr
			}

			if start > -1 {
				op.Body = strings.TrimSpace(text[start:])
			} else if lines := strings.Split(text, "\n"); len(lines) > 1 {
				op.Body = strings.TrimSpace(strings.Join(lines[1:], "\n"))
			}
		} else {
			op.Body = trimmed
		}
	}
}

func parseResponseBody(s *goquery.Selection, op *ParsedOperation) {
	if pre := s.Next(); goquery.NodeName(pre) == "pre" {
		op.Response = strings.TrimSpace(pre.Text())
	}
}

func isDuplicate(ops []ParsedOperation, op ParsedOperation) bool {
	for _, existing := range ops {
		if existing.Group == op.Group && existing.Method == op.Method && existing.URL == op.URL {
			return true
		}
	}
	return false
}

func generateUniqueName(ops []ParsedOperation, op ParsedOperation, config ScraperConfig) string {
	attempt := 0
	for {
		baseName := GenerateName(op.Group, op.Method, op.URL, attempt, config)
		exists := false
		for _, existing := range ops {
			if existing.Group == op.Group && existing.Name == baseName {
				exists = true
				break
			}
		}
		if !exists {
			return baseName
		}
		attempt++
		if attempt > 5 {
			return fmt.Sprintf("%s %d", baseName, attempt)
		}
	}
}

var defaultSingularMap = map[string]string{}

var defaultCasingOverrides = map[string]string{}

func GenerateName(group, method, url string, attempt int, config ScraperConfig) string {
	groupTokens := strings.Fields(strings.ToLower(strings.ReplaceAll(group, "-", " ")))

	parts := strings.Split(url, "/")
	var meaningfulParts []string
	var version string
	var paramsFound []string

	isSingular := strings.Contains(url, "{id}") || method == "POST" || method == "PUT" || method == "DELETE"

	for _, p := range parts {
		if p == "" {
			continue
		}

		if regexp.MustCompile(`^v\d+$`).MatchString(p) {
			version = p
			continue
		}

		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
			cleanP := strings.Trim(p, "{}")
			if strings.ToLower(cleanP) != "id" {
				paramsFound = append(paramsFound, casesTitle(cleanP))
			}
			continue
		}

		cleanP := strings.ReplaceAll(p, "-", "")
		lowerP := strings.ToLower(cleanP)

		isGroupToken := false
		for _, t := range groupTokens {
			if lowerP == t || lowerP == strings.ToLower(strings.ReplaceAll(group, " ", "")) {
				isGroupToken = true
				break
			}
			if attempt == 0 {
				if strings.TrimSuffix(lowerP, "s") == t || strings.TrimSuffix(t, "s") == lowerP {
					isGroupToken = true
					break
				}
			}
		}
		if isGroupToken {
			continue
		}

		finalPart := ""

		// Check Config/Defaults for Singular Map
		singularVal := ""
		if v, ok := config.SingularMap[lowerP]; ok {
			singularVal = v
		} else if v, ok := defaultSingularMap[lowerP]; ok {
			singularVal = v
		}

		if singularVal != "" {
			if isSingular {
				finalPart = singularVal
			} else {
				if singularVal == "UOM" {
					finalPart = "UOMs"
				} else if singularVal == "PatientRegistration" {
					finalPart = "PatientRegistrations"
				} else {
					finalPart = casesTitle(cleanP)
				}
			}
		} else {
			// Check Config/Defaults for Casing Overrides
			casingVal := ""
			if v, ok := config.CasingOverrides[lowerP]; ok {
				casingVal = v
			} else if v, ok := defaultCasingOverrides[lowerP]; ok {
				casingVal = v
			}

			if casingVal != "" {
				finalPart = casingVal
			} else {
				finalPart = casesTitle(cleanP)
			}
		}

		meaningfulParts = append(meaningfulParts, finalPart)
	}

	suffix := strings.Join(meaningfulParts, "")

	var prefix string
	switch method {
	case "GET":
		prefix = "Get"
		if len(meaningfulParts) == 0 && !strings.Contains(url, "{id}") && len(paramsFound) == 0 {
			prefix = "GetAll"
		}
	case "POST":
		prefix = "Create"
	case "PUT":
		prefix = "Update"
	case "DELETE":
		prefix = "Delete"
	default:
		prefix = casesTitle(strings.ToLower(method))
	}

	name := ""
	if strings.HasPrefix(strings.ToLower(suffix), strings.ToLower(prefix)) {
		name = suffix
	} else {
		name = prefix + suffix
	}

	if len(paramsFound) > 0 {
		name += "By" + strings.Join(paramsFound, "And")
	}

	if name == "" {
		name = method
	}

	if version != "" {
		name += " " + casesTitle(version)
	}

	return name
}

func casesTitle(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return strings.ToUpper(s)
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}
