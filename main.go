package main

import (
	"context"
	"flag"

	"github.com/thoohv5/proxy/internal/config"
	"github.com/thoohv5/proxy/internal/proxy"
	pconfig "github.com/thoohv5/proxy/pkg/config"
	"github.com/thoohv5/proxy/pkg/errgroup"
	"github.com/thoohv5/proxy/pkg/file"
)

var (
	cfgFile = flag.String("config", file.AbPath("./config/config.yaml"), "--config=.")
)

type Cfg struct {
	Proxy []*config.Config
}

func main() {
	flag.Parse()

	cfg := &Cfg{}

	if err := pconfig.Parse(*cfgFile, cfg); err != nil {
		panic(err)
	}

	eg := errgroup.WithContext(context.Background())
	for _, c := range cfg.Proxy {
		cc := c
		eg.Go(func(ctx context.Context) error {
			return proxy.Adapter(cc).Handle()
		})
	}
	panic(eg.Wait())
}
