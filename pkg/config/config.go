package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Parse 配置解析
func Parse(file string, cfg interface{}) (err error) {
	// 读取 YAML 文件
	f, err := os.ReadFile(file)
	if err != nil {
		return
	}

	// 解析 YAML 数据
	err = yaml.Unmarshal(f, cfg)
	return
}
