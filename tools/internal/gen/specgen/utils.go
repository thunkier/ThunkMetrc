package specgen

import "strings"

func cleanFileName(s string) string {
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, ":", "")
	return strings.TrimSpace(s)
}

func isV1Endpoint(url string) bool {
	return strings.Contains(url, "/v1/") || strings.HasSuffix(url, "/v1")
}

func singularize(word string) string {
	if strings.HasSuffix(word, "ies") {
		return strings.TrimSuffix(word, "ies") + "y"
	}
	if strings.HasSuffix(word, "ches") || strings.HasSuffix(word, "shes") || strings.HasSuffix(word, "xes") {
		return strings.TrimSuffix(word, "es")
	}
	if strings.HasSuffix(word, "s") && !strings.HasSuffix(word, "ss") {
		return strings.TrimSuffix(word, "s")
	}
	return word
}
