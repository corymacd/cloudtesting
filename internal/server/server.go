package server

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	ver "github.com/cloudtesting/internal/version"
)

type Server struct {
	srv     *http.Server
	version string
}

func New(version string) *Server {
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/healthz", HealthzHandler)
	mux.HandleFunc("/version", makeVersionHandler())
	mux.HandleFunc("/", rootHandler)

	return &Server{
		srv: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
		version: version,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig
		shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		if err := s.srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}
	}()

	log.Printf("Server starting on %s", s.srv.Addr)
	if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("server error: %v", err)
	}

	return nil
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("ok")); err != nil {
		log.Printf("Error writing healthz response: %v", err)
	}
}

func makeVersionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info := ver.GetInfo()

		switch r.Header.Get("Accept") {
		case "application/json":
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(info); err != nil {
				http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
				return
			}
		case "application/xml":
			w.Header().Set("Content-Type", "application/xml")
			if err := xml.NewEncoder(w).Encode(info); err != nil {
				http.Error(w, "Failed to encode XML response", http.StatusInternalServerError)
				return
			}
		default:
			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintf(w, "Version: %s\nGitCommit: %s\nBuildTime: %s\nBuildUser: %s\nGoVersion: %s\n",
				info.Version, info.GitCommit, info.BuildTime, info.BuildUser, info.GoVersion)
		}
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if _, err := w.Write([]byte("API Server")); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
