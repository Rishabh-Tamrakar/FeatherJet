package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"net/http/httputil"
    "net/url"

	"github.com/featherjet/featherjet/internal/config"
	"github.com/featherjet/featherjet/internal/middleware"
)

// Server represents the FeatherJet HTTP server
type Server struct {
	config     *config.Config
	httpServer *http.Server
	mux        *http.ServeMux
}

// New creates a new FeatherJet server instance
func New(cfg *config.Config) *Server {
	if err := cfg.Validate(); err != nil {
		panic(fmt.Sprintf("Invalid configuration: %v", err))
	}

	mux := http.NewServeMux()
	
	server := &Server{
		config: cfg,
		mux:    mux,
		httpServer: &http.Server{
			Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
			IdleTimeout:  cfg.Server.IdleTimeout,
		},
	}

	server.setupRoutes()
	server.setupMiddleware()

	return server
}

func (s *Server) handleTasksProxy(w http.ResponseWriter, r *http.Request) {
    target, _ := url.Parse("http://localhost:8080") // Replace PORT with VelocityTasks port
    proxy := httputil.NewSingleHostReverseProxy(target)
    proxy.ServeHTTP(w, r)
}

// setupRoutes configures the server routes
func (s *Server) setupRoutes() {
	// API routes
	s.mux.HandleFunc("/api/hello", s.handleHello)
	s.mux.HandleFunc("/api/status", s.handleStatus)
	s.mux.HandleFunc("/api/info", s.handleInfo)
	s.mux.HandleFunc("/api/tasks/", s.handleTasksProxy)
	s.mux.HandleFunc("/api/tasks", s.handleTasksProxy)  // Proxy to VelocityTasks
	
	// Static file handler
	staticHandler := s.createStaticFileHandler()
	s.mux.Handle("/", staticHandler)
}

// setupMiddleware configures middleware chain
func (s *Server) setupMiddleware() {
	var handler http.Handler = s.mux

	// Add security headers
	handler = middleware.Security(handler)

	// Add request logging
	handler = middleware.RequestLogger(handler)

	// Add CORS if enabled
	if s.config.Middleware.EnableCORS {
		handler = middleware.CORS(handler)
	}

	s.httpServer.Handler = handler
}

// createStaticFileHandler creates a handler for serving static files
func (s *Server) createStaticFileHandler() http.Handler {
	staticDir := s.config.Static.Directory
	
	// Ensure static directory exists
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		fmt.Printf("Warning: Static directory %s does not exist\n", staticDir)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Static directory not found", http.StatusNotFound)
		})
	}

	fileServer := http.FileServer(http.Dir(staticDir))
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set cache headers for static files
		if s.config.Static.CacheMaxAge != "" {
			w.Header().Set("Cache-Control", "max-age="+s.config.Static.CacheMaxAge)
		}

		// Check if it's an API route
		if len(r.URL.Path) >= 4 && r.URL.Path[:4] == "/api" {
			http.NotFound(w, r)
			return
		}

		// Serve the file or directory listing
		fileServer.ServeHTTP(w, r)
	})
}

// API Handlers

// handleHello responds to /api/hello
func (s *Server) handleHello(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message":   "Hello from FeatherJet!",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"method":    r.Method,
		"path":      r.URL.Path,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// handleStatus responds to /api/status
func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "healthy",
		"server":    "FeatherJet",
		"version":   "1.0.0",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"uptime":    time.Since(time.Now()).String(), // This would be calculated from server start time in production
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// handleInfo responds to /api/info
func (s *Server) handleInfo(w http.ResponseWriter, r *http.Request) {
	// Get static directory info
	staticInfo := map[string]interface{}{
		"directory": s.config.Static.Directory,
		"exists":    false,
	}

	if stat, err := os.Stat(s.config.Static.Directory); err == nil {
		staticInfo["exists"] = true
		staticInfo["is_directory"] = stat.IsDir()
	}

	response := map[string]interface{}{
		"server": map[string]interface{}{
			"name":    "FeatherJet",
			"version": "1.0.0",
			"host":    s.config.Server.Host,
			"port":    s.config.Server.Port,
		},
		"static": staticInfo,
		"middleware": map[string]interface{}{
			"cors_enabled":        s.config.Middleware.EnableCORS,
			"compression_enabled": s.config.Middleware.EnableCompression,
			"request_logging":     s.config.Logging.EnableRequestLogging,
		},
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	fmt.Printf("FeatherJet server listening on %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
