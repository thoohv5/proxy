package proxy

import (
	"github.com/thoohv5/proxy/internal/config"
	"github.com/thoohv5/proxy/internal/proxy/http"
	"github.com/thoohv5/proxy/internal/proxy/standard"
	"github.com/thoohv5/proxy/internal/proxy/tcp"
	"github.com/thoohv5/proxy/internal/proxy/udp"
)

func Adapter(cfg *config.Config) (p standard.IProxy) {
	switch cfg.Type {
	case "http", "https":
		p = http.New(cfg)
	case "tcp":
		p = tcp.New(cfg)
	default:
		p = udp.New(cfg)
	}
	return
}
