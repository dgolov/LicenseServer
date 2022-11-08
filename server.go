package LicenseServer

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s Server) Run(port string, handler http.Handler) error {
	// Run server
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	logrus.Printf("Server start on port %s", port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	// Stop server
	return s.httpServer.Shutdown(ctx)
}
