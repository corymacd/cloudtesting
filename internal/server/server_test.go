package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cloudtesting/internal/version"
)

func TestHealthzHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/healthz", nil)
	rr := httptest.NewRecorder()

	HealthzHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "ok"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %q want %q",
			rr.Body.String(), expected)
	}
}

func TestVersionHandler(t *testing.T) {
	originalVersion := version.Version
	version.Version = "1.0.0-test"
	defer func() { version.Version = originalVersion }()

	req := httptest.NewRequest("GET", "/version", nil)
	rr := httptest.NewRecorder()

	handler := makeVersionHandler()
	handler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Version: 1.0.0-test"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %q want %q",
			rr.Body.String(), expected)
	}
}

func TestRootHandler(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		wantStatus     int
		wantBodyPrefix string
	}{
		{
			name:           "root path",
			path:           "/",
			wantStatus:     http.StatusOK,
			wantBodyPrefix: "API Server",
		},
		{
			name:           "not found",
			path:           "/notfound",
			wantStatus:     http.StatusNotFound,
			wantBodyPrefix: "404",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			rr := httptest.NewRecorder()

			rootHandler(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}

			if !strings.HasPrefix(rr.Body.String(), tt.wantBodyPrefix) {
				t.Errorf("handler returned unexpected body: got %q want prefix %q",
					rr.Body.String(), tt.wantBodyPrefix)
			}
		})
	}
}
