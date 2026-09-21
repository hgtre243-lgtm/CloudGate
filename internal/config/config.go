package config

import "fmt"

// ServerConfig 网关服务端基础配置
type ServerConfig struct {
	Port         int `json:"port"`
	ReadTimeout  int `json:"read_timeout"`  // 读取超时时间(秒)
	WriteTimeout int `json:"write_timeout"` // 写入超时时间(秒)
}

// Route 路由转发规则配置
type RouteConfig struct {
	Path     string `json:"path"`     // 前缀匹配路径，例如 "/api/user"
	Upstream string `json:"upstream"` // 后端服务地址，例如 "http://localhost:8080"
}

// GatewayConfig 网关全局总配置
type GatewayConfig struct {
	Server ServerConfig  `json:"server"`
	Routes []RouteConfig `json:"routes"`
}

// NewDefaultConfig 构造函数：返回一个指向默认配置的指针 (*GatewayConfig)
// 思考：为什么这里返回的是指针 *GatewayConfig 而不是普通结构体？（稍后在逃逸分析中解答）
func NewDefaultConfig() *GatewayConfig {
	return &GatewayConfig{
		Server: ServerConfig{
			Port:         9000,
			ReadTimeout:  5,
			WriteTimeout: 5,
		},
		Routes: []RouteConfig{
			{
				Path:     "/api/user",
				Upstream: "http://localhost:8080", // 预留 Java SpringBoot 服务
			},
		},
	}
}

// UpdatePortByValue [反例演示] 值接收者：传递的是 GatewayConfig 的全量副本
func (c GatewayConfig) UpdatePortByValue(newPort int) {
	c.Server.Port = newPort
	fmt.Printf("[UpdatePortByValue 内部] 端口被修改为了: %d (地址: %p)\n", c.Server.Port, &c)
}

// UpdatePortByPointer [正例实践] 指针接收者：传递的是内存地址，直接原地修改原对象
func (c *GatewayConfig) UpdatePortByPointer(newPort int) {
	c.Server.Port = newPort
	fmt.Printf("[UpdatePortByPointer 内部] 端口被修改为了: %d (地址: %p)\n", c.Server.Port, c)
}
