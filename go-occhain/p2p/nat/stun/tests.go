package stun

import (
	"net"
)

func (c *Client) test1(conn net.PacketConn, addr net.Addr) (*response, error) {
	return c.sendBindingReq(conn, addr, false, false)
}

func (c *Client) test2(conn net.PacketConn, addr net.Addr) (*response, error) {
	return c.sendBindingReq(conn, addr, true, true)
}

func (c *Client) test3(conn net.PacketConn, addr net.Addr) (*response, error) {
	return c.sendBindingReq(conn, addr, false, true)
}
