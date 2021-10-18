package server

import (
	"fmt"
	"github.com/Bailing1992/apigateway/core/config"
	"net"
	"sync/atomic"
	"time"
)

type CommonListener struct {
	// inner listener
	ln net.Listener

	connCount uint64

	// maximum wait time for graceful shutdown
	maxWaitTime time.Duration

	stop chan error

	// becomes non-zero when graceful shutdown starts
	shutdown uint64
}

func NewTcpListener(config config.ServerConfig) (net.Listener, error) {
	addr := fmt.Sprintf("%s:%d", config.GetIP(), config.GetPort())
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("new server listening on %v, error: %v\n", addr, err)
		return nil, err
	}
	return &CommonListener{
		ln:          ln,
		maxWaitTime: config.GetMaxWaitTime(),
		stop:        make(chan error),
	}, nil
}

func (ln *CommonListener) Accept() (net.Conn, error) {
	c, err := ln.ln.Accept()

	if err != nil {
		return nil, err
	}
	atomic.AddUint64(&ln.connCount, 1)
	return &gracefulConn{
		Conn: c,
		ln:   ln,
	}, nil
}

func (ln *CommonListener) Close() error {
	err := ln.ln.Close()
	if err != nil {
		return nil
	}

	return ln.waitForZeroConn()
}

func (ln *CommonListener) Addr() net.Addr {
	return ln.ln.Addr()
}

// ^uint64(0) = -1
func (ln *CommonListener) closeConn() {
	// 相当于减1
	connCount := atomic.AddUint64(&ln.connCount, ^uint64(0))
	if atomic.LoadUint64(&ln.shutdown) != 0 && connCount == 0 {
		close(ln.stop)
	}
}

func (ln *CommonListener) waitForZeroConn() error {
	atomic.AddUint64(&ln.shutdown, 1)

	if atomic.LoadUint64(&ln.connCount) == 0 {
		close(ln.stop)
		return nil
	}

	select {
	case <-ln.stop:
		return nil
	case <-time.After(ln.maxWaitTime):
		return fmt.Errorf("cannot complete graceful shutdown in %s", ln.maxWaitTime)
	}

	return nil
}
