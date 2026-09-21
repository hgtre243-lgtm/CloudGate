package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	// 引入你自己仓库的 config 包 (请将你的 GitHub 用户名替换进去)
	"github.com/hgtre243-lgtm/CloudGate/internal/config"
)

// 全局配置实例指针
var globalConfig *config.GatewayConfig

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 将当前全局配置和运行时状态一并输出
	data := map[string]any{
		"name":           "CloudGate Core",
		"version":        "v0.0.2-alpha",
		"status":         "UP",
		"go_version":     runtime.Version(),
		"num_goroutines": runtime.NumGoroutine(),
		"server_time":    time.Now().Format("2006-01-02 15:04:05"),
		"current_config": globalConfig, // 直接返回当前运行的配置结构体
	}

	_ = json.NewEncoder(w).Encode(data)
}

func main() {
	// 1. 初始化默认配置（返回指针）
	globalConfig = config.NewDefaultConfig()
	fmt.Printf("1. 初始全局配置指针地址: %p, 初始端口: %d\n", globalConfig, globalConfig.Server.Port)

	// 2. 尝试使用【值接收者】修改端口
	fmt.Println("\n--- 试验 A: 值接收者修改 ---")
	globalConfig.UpdatePortByValue(9001)
	fmt.Printf("执行 UpdatePortByValue 后，外部全局端口为: %d (根本没有变!)\n", globalConfig.Server.Port)

	// 3. 尝试使用【指针接收者】修改端口
	fmt.Println("\n--- 试验 B: 指针接收者修改 ---")
	globalConfig.UpdatePortByPointer(9002)
	fmt.Printf("执行 UpdatePortByPointer 后，外部全局端口为: %d (成功被原地修改!)\n\n", globalConfig.Server.Port)

	// 将端口格式化
	addr := fmt.Sprintf(":%d", globalConfig.Server.Port)

	http.HandleFunc("/ping", pingHandler)

	fmt.Printf("============================================\n")
	fmt.Printf("🚀 CloudGate 启动成功，监听端口: %s\n", addr)
	fmt.Printf("👉 测试地址: http://localhost%s/ping\n", addr)
	fmt.Printf("============================================\n")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
