package network

import (
	"context"

	ma "github.com/multiformats/go-multiaddr"
)

// Conn 是一个通用的连接接口
type Conn interface {
	Read([]byte) (int, error)  // 读取数据
	Write([]byte) (int, error) // 写入数据
	Close() error              // 关闭连接
}

// Transport 定义了传输层通用接口
type Transport interface {
	Listen(addr ma.Multiaddr) error                            // 监听指定的multiaddr地址
	Accept() (Conn, error)                                     // 接受连接
	Dail(ctx context.Context, addr ma.Multiaddr) (Conn, error) // 连接到multiaddr地址
	Close() error                                              // 关闭传输层
}
