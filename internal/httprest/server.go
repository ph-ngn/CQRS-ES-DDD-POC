package httprest

import (
	"context"
	"net/http"

	"github.com/andyj29/wannabet/internal/application/common"
	"github.com/andyj29/wannabet/internal/infrastructure/oauth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type server struct {
	s            http.Server
	router       *chi.Mux
	logger       common.Logger
	heathChecker http.HandlerFunc
	readyChecker http.HandlerFunc
}

type Config struct {
	Address string
	Logger  common.Logger
	Router  *chi.Mux
}

func NewHTTPServer(cfg Config) (*server, error) {
	var server server
	server.s.Addr = ":8080"
	if cfg.Address != "" {
		server.s.Addr = cfg.Address
	}
	server.logger = cfg.Logger
	server.router = cfg.Router
	server.setupMiddleware()
	return &server, nil
}

func (s *server) RegisterHandler(httpMethod, pattern string, handler http.HandlerFunc) {
	s.router.Method(httpMethod, pattern, handler)
}

func (s *server) ListenandServe() error {
	s.logger.Infof("Starting HTTP server running on %s", s.s.Addr)
	return s.s.ListenAndServe()
}

func (s *server) setupMiddleware() {
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		MaxAge:         300,
	}), oauth.ValidateToken())
}

func (s *server) setHealthChecker(h http.HandlerFunc) {
	s.heathChecker = h
}

func (s *server) setReadyChecker(h http.HandlerFunc) {
	s.readyChecker = h
}

func (s *server) Shutdown(ctx context.Context) error {
	s.logger.Infof("Gracefully shutting down HTTP server running on %s", s.s.Addr)
	return s.s.Shutdown(ctx)
}
