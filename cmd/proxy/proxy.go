package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/thoohv5/proxy/internal/config"
	"github.com/thoohv5/proxy/internal/proxy"
	pconfig "github.com/thoohv5/proxy/pkg/config"
	"github.com/thoohv5/proxy/pkg/errgroup"
	"github.com/thoohv5/proxy/pkg/file"
)

var (
	cfgData = flag.String("data", "", "{\"proxy\":[{\"type\":\"http\",\"listen\":{\"network\":\"\",\"address\":\"0.0.0.0:8081\"},\"dial\":{\"network\":\"\",\"address\":\"127.0.0.1:8080\"}}]},配置数据，type支持http/tcp/udp")
	cfgFile = flag.String("config", file.AbPath("../../config/config.yaml"), "--config=.")
)

type Cfg struct {
	Proxy []*config.Config
}

func main() {
	flag.Parse()

	cfg := &Cfg{}

	if len(*cfgData) > 0 {
		if err := pconfig.ParseData(*cfgData, cfg); err != nil {
			panic(err)
		}
	} else if len(*cfgFile) > 0 {
		if err := pconfig.ParseFile(*cfgFile, cfg); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("请配置数据")
		return
	}

	fmt.Println("配置加载")

	eg := errgroup.WithContext(context.Background())
	for _, c := range cfg.Proxy {
		cc := c
		eg.Go(func(ctx context.Context) error {
			return proxy.Adapter(cc).Handle()
		})
	}
	fmt.Println("服务启动")
	panic(eg.Wait())
}
