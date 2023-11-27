package tcp

import (
	"fmt"
	"io"
	"net"

	"github.com/thoohv5/proxy/internal/config"
	"github.com/thoohv5/proxy/internal/proxy/standard"
)

type tcp struct {
	cfg *config.Config
}

func New(c *config.Config) standard.IProxy {
	return &tcp{
		cfg: c,
	}
}

// Handle 处理
func (p *tcp) Handle() (err error) {
	// 监听本地端口
	listener, err := net.Listen(p.cfg.Listen.Network, p.cfg.Listen.Address)
	if err != nil {
		return
	}
	defer func() {
		cErr := listener.Close()
		if cErr != nil {
			if err != nil {
				err = fmt.Errorf("err :%w, close err: %w", err, cErr)
			} else {
				err = fmt.Errorf("close err: %w", cErr)
			}
		}
	}()

	for {
		// 等待客户端连接
		conn, aErr := listener.Accept()
		if aErr != nil {
			err = aErr
			return
		}

		// 启动新的 goroutine 处理连接
		go func() {
			if err = p.dealClient(conn); err != nil {
				return
			}
		}()
	}
}

// DealClient 处理client连接
func (p *tcp) dealClient(clientConn net.Conn) (err error) {
	// 连接到实际的后端服务器
	conn, err := net.Dial(p.cfg.Dial.Network, p.cfg.Dial.Address)
	if err != nil {
		return
	}
	defer func() {
		cErr := conn.Close()
		if cErr != nil {
			if err != nil {
				err = fmt.Errorf("err :%w, close err: %w", err, cErr)
			} else {
				err = fmt.Errorf("close err: %w", cErr)
			}
		}
	}()

	// 启动 goroutine 将客户端的数据转发到后端服务器
	go func() {
		if _, err = io.Copy(conn, clientConn); err != nil {
			return
		}
	}()

	// 将后端服务器的数据转发回客户端
	if _, err = io.Copy(clientConn, conn); err != nil {
		return
	}

	// 关闭连接
	if err = clientConn.Close(); err != nil {
		return
	}
	return
}
