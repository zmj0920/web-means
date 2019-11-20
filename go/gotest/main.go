package main

import (
	"net/http"
	"time"
)

func main() {
	setup()
	// 路由
	mux := http.NewServeMux()

	// 获取文章列表
	mux.HandleFunc("/api/threads", getThreads)
	// 启动服务器
	server := &http.Server{
		Addr:           Config.ListeningAddress,
		Handler:        mux,
		ReadTimeout:    time.Duration(Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(Config.ReadTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
