package specgen

import (
	"io"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/thunkier/thunkmetrc/tools/internal/app"
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
	Required    bool
}

func ParseDocumentation(r io.Reader, cfg app.ScraperConfig) ([]ParsedOperation, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var operations []ParsedOperation

	doc.Find("h1, h2").Each(func(i int, s *goquery.Selection) {
		groupName := strings.TrimSpace(s.Text())

		s.NextUntil("h1, h2").Filter("h3").Each(func(j int, h3 *goquery.Selection) {
			opName := strings.TrimSpace(h3.Text())

			var method, urlStr string
			parts := strings.Fields(opName)
			if len(parts) >= 2 && isValidMethod(parts[0]) {
				method = parts[0]
				urlStr = parts[1]
			} else {
				if p := h3.Next(); goquery.NodeName(p) == "p" {
					parts := strings.Fields(p.Text())
					if len(parts) >= 2 && isValidMethod(parts[0]) {
						method = parts[0]
						urlStr = parts[1]
					}
				}
			}

			if method == "" || urlStr == "" {
				return
			}

			descr := ""
			h3.NextAll().EachWithBreak(func(_ int, sel *goquery.Selection) bool {
				if goquery.NodeName(sel) == "p" && !strings.Contains(sel.Text(), method) {
					descr = strings.TrimSpace(sel.Text())
					return false
				}
				if goquery.NodeName(sel) == "h3" || goquery.NodeName(sel) == "h2" {
					return false
				}
				return true
			})

			op := ParsedOperation{
				Group:       groupName,
				Name:        opName,
				Method:      method,
				URL:         urlStr,
				Description: descr,
			}

			// Parameters
			h3.NextAll().EachWithBreak(func(_ int, sel *goquery.Selection) bool {
				text := sel.Text()
				if strings.Contains(text, "Parameters") {
					parseParameters(sel, &op)
					return false
				}
				if goquery.NodeName(sel) == "h3" || goquery.NodeName(sel) == "h2" {
					return false
				}
				return true
			})

			h3.NextAll().EachWithBreak(func(_ int, sel *goquery.Selection) bool {
				text := sel.Text()
				if strings.Contains(text, "Permissions Required") {
					parsePermissions(sel, &op)
					return false
				}
				if goquery.NodeName(sel) == "h3" || goquery.NodeName(sel) == "h2" {
					return false
				}
				return true
			})

			h3.NextAll().EachWithBreak(func(_ int, sel *goquery.Selection) bool {
				text := sel.Text()
				if strings.Contains(text, "Request Body") || strings.Contains(text, "Example Request") {
					parseRequestBody(sel, &op)
					return false
				}
				if goquery.NodeName(sel) == "h3" || goquery.NodeName(sel) == "h2" {
					return false
				}
				return true
			})

			h3.NextAll().EachWithBreak(func(_ int, sel *goquery.Selection) bool {
				text := sel.Text()
				if strings.Contains(text, "Response Body") || strings.Contains(text, "Example Response") {
					parseResponseBody(sel, &op)
					return false
				}
				if goquery.NodeName(sel) == "h3" || goquery.NodeName(sel) == "h2" {
					return false
				}
				return true
			})

			operations = append(operations, op)
		})
	})

	return operations, nil
}

func parseParameters(s *goquery.Selection, op *ParsedOperation) {
	if table := s.Next(); goquery.NodeName(table) == "table" {
		table.Find("tr").Each(func(_ int, row *goquery.Selection) {
			cols := row.Find("td")
			if cols.Length() >= 2 {
				nameCol := cols.Eq(0)
				pName := ""

				if strong := nameCol.Find("strong"); strong.Length() > 0 {
					pName = strings.TrimSpace(strong.Text())
				} else {
					pName = strings.Fields(nameCol.Text())[0]
				}

				isRequired := true
				if strings.Contains(strings.ToLower(nameCol.Text()), "optional") {
					isRequired = false
				}

				pDesc := strings.TrimSpace(cols.Eq(1).Text())
				op.Params = append(op.Params, Param{
					Name:        pName,
					Description: pDesc,
					Required:    isRequired,
				})
			}
		})
	}
}

func parsePermissions(s *goquery.Selection, op *ParsedOperation) {
	next := s.Next()
	if goquery.NodeName(next) == "p" || goquery.NodeName(next) == "ul" {
		text := strings.TrimSpace(next.Text())
		if text != "" && text != "None" {
			parts := strings.Split(text, ",")
			for _, p := range parts {
				clean := strings.TrimSpace(p)
				if clean != "" {
					op.Permissions = append(op.Permissions, clean)
				}
			}
		}
	} else if goquery.NodeName(next) == "table" {
		next.Find("tr").Each(func(_ int, row *goquery.Selection) {
			text := strings.TrimSpace(row.Text())
			if text != "" && text != "None" {
				op.Permissions = append(op.Permissions, text)
			}
		})
	}
}

func parseRequestBody(s *goquery.Selection, op *ParsedOperation) {
	if pre := s.Next(); goquery.NodeName(pre) == "pre" {
		text := pre.Text()
		trimmed := strings.TrimSpace(text)
		if regexp.MustCompile(`(?i)^[A-Z]+\s+/`).MatchString(trimmed) {
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

func isValidMethod(m string) bool {
	m = strings.ToUpper(m)
	return m == "GET" || m == "POST" || m == "PUT" || m == "PATCH" || m == "DELETE"
}

var defaultSingularMap = map[string]string{}

var defaultCasingOverrides = map[string]string{
	"createpackages":           "Packages",
	"createplantbatchpackages": "PlantBatchPackages",
}

func GenerateName(group, method, url string, attempt int, cfg app.ScraperConfig) string {
	groupTokens := strings.Fields(strings.ToLower(strings.ReplaceAll(group, "-", " ")))

	parts := strings.Split(url, "/")
	var meaningfulParts []string
	var paramsFound []string

	isSingular := strings.Contains(url, "{id}")

	for _, p := range parts {
		if p == "" {
			continue
		}

		if regexp.MustCompile(`^v\d+$`).MatchString(p) {
			continue
		}

		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
			cleanP := strings.Trim(p, "{}")
			paramsFound = append(paramsFound, casesTitle(cleanP))
			continue
		}

		cleanP := strings.ReplaceAll(p, "-", "")
		lowerP := strings.ToLower(cleanP)

		if strings.EqualFold(group, "ProcessingJob") && lowerP == "processing" {
			continue
		}
		if lowerP == "all" {
			continue
		}

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

		singularVal := ""
		if v, ok := cfg.SingularMap[lowerP]; ok {
			singularVal = v
		} else if v, ok := defaultSingularMap[lowerP]; ok {
			singularVal = v
		}

		if singularVal != "" {
			if isSingular {
				finalPart = singularVal
			} else {
				finalPart = casesTitle(cleanP)
			}
		} else {
			casingVal := ""
			if v, ok := cfg.CasingOverrides[lowerP]; ok {
				casingVal = v
			} else if v, ok := defaultCasingOverrides[lowerP]; ok {
				casingVal = v
			}

			if casingVal != "" {
				if isSingular {
					finalPart = singularize(casingVal)
				} else {
					finalPart = casingVal
				}
			} else {
				// Use helper to singularize if needed
				if isSingular {
					finalPart = casesTitle(singularize(cleanP))
				} else {
					finalPart = casesTitle(cleanP)
				}
			}
		}

		meaningfulParts = append(meaningfulParts, finalPart)
	}

	var adjectives []string
	var nouns []string
	for _, p := range meaningfulParts {
		lower := strings.ToLower(p)
		if lower == "active" || lower == "inactive" || lower == "intransit" || lower == "onhold" ||
			lower == "vegetative" || lower == "flowering" || lower == "mother" || lower == "adjust" ||
			lower == "start" || lower == "finish" || lower == "unfinish" ||
			lower == "staged" || lower == "available" || lower == "incoming" || lower == "outgoing" || lower == "rejected" || lower == "externalincoming" {
			adjectives = append(adjectives, p)
		} else {
			nouns = append(nouns, p)
		}
	}
	if len(nouns) == 0 {
		groupName := casesTitle(strings.ReplaceAll(group, " ", ""))
		nouns = append(nouns, groupName)
	}

	meaningfulParts = append(adjectives, nouns...)
	suffix := strings.Join(meaningfulParts, "")

	var prefix string
	switch method {
	case "GET":
		prefix = "Get"
	case "POST":
		prefix = "Create"
	case "PUT":
		prefix = "Update"
	case "PATCH":
		prefix = "Patch"
	case "DELETE":
		prefix = "Delete"
	default:
		prefix = casesTitle(strings.ToLower(method))
	}

	name := ""
	lowerSuffix := strings.ToLower(suffix)
	if strings.HasPrefix(lowerSuffix, strings.ToLower(prefix)) {
		name = suffix
	} else if strings.HasPrefix(lowerSuffix, "start") || strings.HasPrefix(lowerSuffix, "finish") || strings.HasPrefix(lowerSuffix, "unfinish") {
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
