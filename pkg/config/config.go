package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// ParseFile 配置解析
func ParseFile(file string, cfg interface{}) (err error) {
	// 读取 YAML 文件
	f, err := os.ReadFile(file)
	if err != nil {
		return
	}

	// 解析 YAML 数据
	err = yaml.Unmarshal(f, cfg)
	return
}

// ParseData 配置解析
func ParseData(data string, cfg interface{}) (err error) {
	// 解析 YAML 数据
	err = yaml.Unmarshal([]byte(data), cfg)
	return
}
