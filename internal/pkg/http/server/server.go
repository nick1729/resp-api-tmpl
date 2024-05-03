package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nick1729/resp-api-tmpl/internal/app/handler"
	"github.com/nick1729/resp-api-tmpl/internal/pkg/config"
	"github.com/nick1729/resp-api-tmpl/internal/pkg/repository"
	"github.com/rs/zerolog/log"
)

// Server - structure for HTTP server.
type Server struct {
	app    *fiber.App // http server
	notify chan error // server errors
}

// Notify - outputs errors from the server error channel.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Run  - starts an HTTP server at address in a separate goroutine.
func (s *Server) Run(address string) {
	go func() {
		err := s.app.Listen(address)
		if err != nil {
			s.notify <- err
		}
	}()
}

// Stop  - stops the server.
func (s *Server) Stop() {
	err := s.app.Shutdown()
	if err != nil {
		s.notify <- err
	}
}

// New  - creates a new HTTP server.
func New(cfg config.Server, appName string, repo *repository.Repository) *Server {
	app := fiber.New(fiber.Config{
		AppName:               appName,
		BodyLimit:             cfg.BodyLimit,
		DisableStartupMessage: true,
	})

	app.Server().Logger = &log.Logger // TODO
	app.Use(cors.New())
	//app.Use(pprof.New())

	handler.RouteRegister(app, repo)

	server := &Server{
		app:    app,
		notify: make(chan error, 1),
	}

	return server
}
