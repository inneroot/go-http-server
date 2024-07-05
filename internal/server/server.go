package server

import (
	"context"
	"log/slog"
	"net/http"
)

type HttpServer struct {
	log *slog.Logger
	srv *http.Server
}

func New(port string, log *slog.Logger) *HttpServer {
	return &HttpServer{
		log: log,
		srv: &http.Server{
			Addr: ":" + port,
		},
	}
}

func (s *HttpServer) Start() {
	s.log.Info("staring http server")
	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			s.log.Error("http server error", slog.Any("error", err))
			if err != http.ErrServerClosed {
				panic("http server error")
			}
		}
	}()
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
