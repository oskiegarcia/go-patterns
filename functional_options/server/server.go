package server

import (
	"fmt"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}
type Option func(*Server)

func New(options ...Option) *Server {
	svr := &Server{}
	for _, f := range options {
		f(svr)
	}
	return svr
}
func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) Option {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

func (s *Server) Start() error {
	fmt.Printf("%v", s)
	return nil
}
