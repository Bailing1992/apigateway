package server

import "net"

type gracefulConn struct {
	net.Conn
	ln *CommonListener
}

func (c *gracefulConn) Close() error {
	defer c.ln.closeConn()

	err := c.Conn.Close()

	if err != nil {
		return err
	}
	return nil
}
