/*
Copyright © 2022 Cory M. MacDonald <corymacd@netsrv.co>
*/
package cmd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cloudtesting/internal/server"
)

func TestHealthzHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/healthz", nil)
	rr := httptest.NewRecorder()

	server.HealthzHandler(rr, req)

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
	tests := []struct {
		name        string
		acceptType  string
		wantStatus  int
		wantContent string
	}{
		{
			name:        "plain text",
			acceptType:  "text/plain",
			wantStatus:  http.StatusOK,
			wantContent: "Version:",
		},
		{
			name:        "json",
			acceptType:  "application/json",
			wantStatus:  http.StatusOK,
			wantContent: `{"version":"dev"`,
		},
		{
			name:        "xml",
			acceptType:  "application/xml",
			wantStatus:  http.StatusOK,
			wantContent: `<VersionInfo><version>dev</version>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/version", nil)
			req.Header.Set("Accept", tt.acceptType)
			rr := httptest.NewRecorder()

			server.VersionHandler(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}

			if !strings.Contains(rr.Body.String(), tt.wantContent) {
				t.Errorf("handler returned unexpected body: got %q, want it to contain %q",
					rr.Body.String(), tt.wantContent)
			}
		})
	}
}
