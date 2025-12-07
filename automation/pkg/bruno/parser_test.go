package bruno

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseRequest(t *testing.T) {
	// Create a temp file
	// Note: Our strict parser expects blocks to be well-formed.
	content := `meta {
  name: Test Request
  type: http
  seq: 1
}

get {
  url: https://api.example.com/v1/test
  body: none
  auth: none
}

headers {
  Content-Type: application/json
}

body:json {
  {
    "foo": "bar"
  }
}
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.bru")
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	req, err := ParseRequest(tmpFile)
	if err != nil {
		t.Fatalf("ParseRequest failed: %v", err)
	}

	if req.Name != "Test Request" {
		t.Errorf("expected name 'Test Request', got '%s'", req.Name)
	}
	if req.Method != "GET" {
		t.Errorf("expected method 'GET', got '%s'", req.Method)
	}
	if req.URL != "https://api.example.com/v1/test" {
		t.Errorf("expected URL 'https://api.example.com/v1/test', got '%s'", req.URL)
	}
	if req.Seq != 1 {
		t.Errorf("expected Seq 1, got %d", req.Seq)
	}
	// Check body (simple check)
	if len(req.Body) == 0 {
		t.Error("expected body to be populated")
	}
}

func TestParseCollection(t *testing.T) {
	tmpDir := t.TempDir()
	// Create two files
	f1 := filepath.Join(tmpDir, "req1.bru")
	f2 := filepath.Join(tmpDir, "req2.bru")

	_ = os.WriteFile(f1, []byte("meta { \n name: Req1 \n } \n get { \n url: http://foo \n }"), 0644)
	_ = os.WriteFile(f2, []byte("meta { \n name: Req2 \n } \n post { \n url: http://bar \n }"), 0644)

	col, err := ParseCollection(tmpDir)
	if err != nil {
		t.Fatalf("ParseCollection failed: %v", err)
	}

	if len(col.Requests) != 2 {
		t.Errorf("expected 2 requests, got %d", len(col.Requests))
	}
}
