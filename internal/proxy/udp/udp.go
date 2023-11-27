package udp

import (
	"net"

	"github.com/thoohv5/proxy/internal/config"
	"github.com/thoohv5/proxy/internal/proxy/standard"
)

type udp struct {
	cfg *config.Config
}

func New(c *config.Config) standard.IProxy {
	return &udp{
		cfg: c,
	}
}

// Handle 处理
func (p *udp) Handle() (err error) {
	// 监听本地端口
	listener, err := net.ListenPacket(p.cfg.Listen.Network, p.cfg.Listen.Address)
	if err != nil {
		return
	}
	defer listener.Close()

	// 远程 UDP 地址
	remoteAddr, err := net.ResolveUDPAddr(p.cfg.Dial.Network, p.cfg.Dial.Address)
	if err != nil {
		return
	}

	// 转发接收到的数据包到远程服务器
	buffer := make([]byte, 1024)
	for {
		n, _, rErr := listener.ReadFrom(buffer)
		if rErr != nil {
			err = rErr
			return
		}

		// 发送接收到的数据包到远程服务器
		_, err = listener.WriteTo(buffer[:n], remoteAddr)
		if err != nil {
			return
		}
	}
}
