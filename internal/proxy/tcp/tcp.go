package tcp

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/thoohv5/proxy/internal/config"
	"github.com/thoohv5/proxy/internal/proxy/standard"
	"github.com/thoohv5/proxy/pkg/errgroup"
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

	errCh := make(chan error)
	for {
		select {
		case err = <-errCh:
			return err
		default:
			// 等待客户端连接
			conn, aErr := listener.Accept()
			if aErr != nil {
				err = aErr
				return
			}
			// 启动新的 goroutine 处理连接
			go func() {
				var gErr error
				defer func() {
					// 关闭连接
					cErr := conn.Close()
					if cErr != nil {
						if gErr != nil {
							gErr = fmt.Errorf("err :%w, close err: %w", gErr, cErr)
						} else {
							gErr = fmt.Errorf("close err: %w", cErr)
						}
					}
					if gErr != nil {
						errCh <- gErr
					}
				}()
				if gErr = p.dealClient(context.Background(), conn); gErr != nil {
					return
				}
			}()
		}
	}
}

// DealClient 处理client连接
func (p *tcp) dealClient(ctx context.Context, cConn net.Conn) (err error) {
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

	eg := errgroup.WithContext(ctx)

	// 启动 goroutine 将客户端的数据转发到后端服务器
	eg.Go(func(ctx context.Context) error {
		if _, cErr := io.Copy(conn, cConn); err != nil {
			return cErr
		}
		return nil
	})

	// 将后端服务器的数据转发回客户端
	eg.Go(func(ctx context.Context) error {
		if _, cErr := io.Copy(cConn, conn); err != nil {
			return cErr
		}
		return nil
	})

	if err = eg.Wait(); err != nil {
		return
	}

	return
}
