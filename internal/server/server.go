package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	srv *http.Server
}

func New() *Server {
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/healthz", HealthzHandler)
	mux.HandleFunc("/version", VersionHandler)
	mux.HandleFunc("/", rootHandler)

	return &Server{
		srv: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
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
	w.Write([]byte("ok"))
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	version := "dev" // or however you get your version

	switch r.Header.Get("Accept") {
	case "application/json":
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"version":"%s"}`, version)
	case "application/xml":
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintf(w, `<VersionInfo><version>%s</version></VersionInfo>`, version)
	default:
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Version: %s", version)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("API Server"))
}
