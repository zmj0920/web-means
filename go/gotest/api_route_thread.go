package main

import (
	"encoding/json"
	"net/http"

	"github.com/zmj/gotest/data"
)

// 获取文章列表， 支持GET/POST, 返回json
// 失败返回 错误信息 状态码设置为201
// TODO: 每次返回部分文章，分页控制
func getThreads(writer http.ResponseWriter, request *http.Request) {
	ts, err := data.Threads()
	if err != nil {
		api_ret_err(writer, request, []byte("Get threads fail"))
		return
	}
	js, err := json.Marshal(ts)
	if err != nil {
		api_ret_err(writer, request, []byte("Marshal threads fail"))
		return
	}
	api_ret_json(writer, request, js)
}
