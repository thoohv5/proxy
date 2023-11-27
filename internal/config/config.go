package config

// NetConfig 网络配置
type NetConfig struct {
	/**
	"tcp"、"tcp4"、"tcp6"、"unix" 或 "unixpacket"
	"udp"、"udp4"、"udp6"、"unixgram" 或 IP 传输协议。IP 传输协议可以是 "ip"、"ip4" 或 "ip6"，后面跟着冒号和字面协议号或协议名，例如 "ip:1" 或 "ip:icmp"。
	*/
	Network string `yaml:"network"`
	/**
	对于 TCP 网络，address 参数可以是空或未指定 IP 地址。
	如果主机为空或为文字未指定的 IP 地址，Listen 会监听本地系统所有可用的单播和任播 IP 地址。
	若要仅使用 IPv4，请使用网络类型 "tcp4"。
	地址可以使用主机名，但不推荐，因为这将最多为主机的一个 IP 地址创建监听器。
	如果地址参数中的端口为空或为 "0"，比如 "127.0.0.1:" 或 "[::1]:0"，会自动选择一个端口号。可以使用 Listener 的 Addr 方法来查找选择的端口。
	*/
	Address string `yaml:"address"`
}

// Config 配置
type Config struct {
	// 类型 http/tcp/udp
	Type string `yaml:"type"`
	// 监听配置
	Listen *NetConfig `yaml:"listen"`
	// 连接配置
	Dial *NetConfig `yaml:"dial"`
}
