/*
Copyright © 2022 Cory M. MacDonald <corymacd@netsrv.co>
*/
package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/cloudtesting/internal/server"
)

func TestServerIntegration(t *testing.T) {
	srv := server.New("test-version")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start server
	go func() {
		if err := srv.Run(ctx); err != nil {
			t.Errorf("server error: %v", err)
		}
	}()

	// Wait for server to start
	time.Sleep(100 * time.Millisecond)

	// Test endpoints
	resp, err := http.Get("http://localhost:8080/healthz")
	if err != nil {
		t.Fatalf("could not make request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", resp.StatusCode, http.StatusOK)
	}
}
