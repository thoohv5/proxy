package http

import (
	"fmt"
	nh "net/http"
	"net/http/httputil"
	"net/url"

	"github.com/thoohv5/proxy/internal/config"
	"github.com/thoohv5/proxy/internal/proxy/standard"
)

type http struct {
	cfg *config.Config
}

func New(c *config.Config) standard.IProxy {
	return &http{
		cfg: c,
	}
}

// Handle 处理
func (p *http) Handle() (err error) {
	// 设置要转发的地址
	target, err := url.Parse(fmt.Sprintf("%s://%s", p.cfg.Type, p.cfg.Dial.Address))
	if err != nil {
		panic(err)
	}

	// 实例化 ReverseProxy 包
	rp := httputil.NewSingleHostReverseProxy(target)

	// 设置代理服务器
	server := nh.Server{
		Addr:    p.cfg.Listen.Address, // 代理服务器监听的地址和端口
		Handler: rp,
	}

	// 启动服务
	if err = server.ListenAndServe(); err != nil {
		return
	}
	return
}
