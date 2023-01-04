package httprest

import (
	"context"
	"net/http"

	"github.com/andyj29/wannabet/internal/log"
	"github.com/andyj29/wannabet/internal/oauth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	s            http.Server
	router       *chi.Mux
	logger       log.Logger
	heathChecker http.HandlerFunc
	readyChecker http.HandlerFunc
}

type Config struct {
	Address string
	Logger  log.Logger
	Router  *chi.Mux
}

func NewServer(cfg Config) (*Server, error) {
	var server Server
	server.s.Addr = ":8080"
	if cfg.Address != "" {
		server.s.Addr = cfg.Address
	}
	server.logger = cfg.Logger
	server.router = cfg.Router
	server.setupMiddleware()
	return &server, nil
}

func (s *Server) RegisterHandler(httpMethod, pattern string, handler http.HandlerFunc) {
	s.router.Method(httpMethod, pattern, handler)
}

func (s *Server) ListenandServe() error {
	s.logger.Infof("Setting up health and ready checker")
	s.router.Post("/api/health-check", s.heathChecker)
	s.router.Post("/api/ready-check", s.readyChecker)

	s.logger.Infof("Starting HTTP server running on %s", s.s.Addr)
	return s.s.ListenAndServe()
}

func (s *Server) setupMiddleware() {
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		MaxAge:         300,
	}), oauth.ValidateToken())
}

func (s *Server) setHealthChecker(h http.HandlerFunc) {
	s.heathChecker = h
}

func (s *Server) setReadyChecker(h http.HandlerFunc) {
	s.readyChecker = h
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Infof("Gracefully shutting down HTTP server running on %s", s.s.Addr)
	return s.s.Shutdown(ctx)
}
