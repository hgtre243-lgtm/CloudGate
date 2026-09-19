package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

// ServerInfo 定义服务状态信息的结构体（对应 Java 的 POJO / Entity）
// 后面的 `json:"..."` 叫做 Struct Tag（结构体标签），决定序列化为 JSON 时的字段名
type ServerInfo struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Status     string `json:"status"`
	GoVersion  string `json:"go_version"`
	Goroutines int    `json:"num_goroutines"`
	ServerTime string `json:"sever_time"`
}

// pingHandler 处理探活请求
func pingHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头为 JSON 格式
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 收集系统运行时数据
	info := ServerInfo{
		Name:       "CloudGate Core",
		Version:    "v0.0.1-alpha",
		Status:     "Up",
		GoVersion:  runtime.Version(),
		Goroutines: runtime.NumGoroutine(),
		ServerTime: time.Now().Format("2006-101-02 15:04:05"),
	}

	// 将结构体序列化为 JSON 并写入响应流
	if err := json.NewEncoder(w).Encode(info); err != nil {
		http.Error(w, "Internet Server Error", http.StatusInternalServerError)
	}
}
func main() {
	port := ":9000"

	// 注册路由：当请求 /ping 时由 pingHandler 函数处理
	http.HandleFunc("/ping", pingHandler)

	fmt.Printf("============================================")
	fmt.Printf("🚀 CloudGate 核心服务正在启动...\n")
	fmt.Printf("📡 监听端口: %s\n", port)
	fmt.Printf("👉 探活测试接口: http://localhost%s/ping\n", port)
	fmt.Printf("============================================\n")

	// 启动 HTTP 服务并监听端口（若端口被占用会返回错误）
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}

}
