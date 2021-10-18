package fasthttp

import (
	"fmt"
	"github.com/Bailing1992/apigateway/core/config"
	"github.com/Bailing1992/apigateway/core/server"
	"github.com/valyala/fasthttp"
	"net"
)

type Server struct {
	server.Server
	config     config.ServerConfig
	httpServer *fasthttp.Server
	listener   net.Listener
}

func (s *Server) getDefaultIP() string {
	return "0.0.0.0"
}

func NewServer(config config.ServerConfig, requestHandler fasthttp.RequestHandler) *Server {
	httpServer := &fasthttp.Server{
		Name:    config.GetName(),
		Handler: requestHandler,
	}
	return &Server{
		config:     config,
		httpServer: httpServer,
	}
}

func (s *Server) StartServer() error {
	var (
		ln  net.Listener
		err error
	)
	ln, err = server.NewTcpListener(s.config)
	if err != nil {
		return err
	}
	s.listener = ln
	err = s.httpServer.Serve(ln)
	if err != nil {
		fmt.Printf("http server bind error: %v\n", err)
		return err
	}

	return nil
}
