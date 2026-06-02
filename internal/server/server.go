package server

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
)

type Server struct {
	httpServer *http.Server
	logger     *zerolog.Logger
}

func (s *Server) Run() {
	log.Info().Msgf("server started at: %v", s.httpServer.Addr)

	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Failed run HTTP server")
	}

	return
}

func (s *Server) Shutdown(ctx context.Context) {
	<-ctx.Done()
	s.logger.Info().Msg("shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		s.logger.Fatal().Err(err).Msg("Server forced to shutdown")
	}
	s.logger.Info().Msg("server closed")

	return
}

func NewServer(host, port string, handler http.Handler, logger *zerolog.Logger) *Server {
	httpServer := &http.Server{
		Addr:           host + ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return &Server{httpServer: httpServer, logger: logger}
}
