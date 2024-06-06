package network

import (
	"context"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

type TCPTransport struct {
	listener manet.Listener
}

// Listen 监听指定的multiaddr地址
func (t *TCPTransport) Listen(addr ma.Multiaddr) error {
	ln, err := manet.Listen(addr)
	if err != nil {
		return err
	}
	t.listener = ln
}

// Accept 接受连接
func (t *TCPTransport) Accept() (Conn, error) {
	conn, err := t.listener.Accept()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Dail 连接到multiaddr地址
func (t *TCPTransport) Dail(ctx context.Context, addr ma.Multiaddr) (Conn, error) {
	var d manet.Dialer
	conn, err := d.DialContext(ctx, addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
