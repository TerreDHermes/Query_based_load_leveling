package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	defaultPort = 8080
	defaultHost = "localhost"
)

type Server struct {
	host       string
	port       int
	timeout    time.Duration
	maxConn    int
	handler    http.Handler
	httpServer *http.Server
}

func New(handler http.Handler, options ...func(*Server)) *Server {
	svr := &Server{
		port:    defaultPort,
		host:    defaultHost,
		handler: handler,
	}
	for _, o := range options {
		o(svr)
	}
	return svr
}

func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%d", s.host, s.port)
	srv := &http.Server{
		Addr:         address,
		Handler:      s.handler,
		ReadTimeout:  s.timeout,
		WriteTimeout: s.timeout,
	}
	s.httpServer = srv
	return srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	logrus.Info("Закрываем сервер")
	return s.httpServer.Shutdown(ctx)
}

func WithHost(host string) func(*Server) {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) func(*Server) {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

func WithHandler(handler http.Handler) func(*Server) {
	return func(s *Server) {
		s.handler = handler
	}
}
