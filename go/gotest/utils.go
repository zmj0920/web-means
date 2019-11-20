package main

import (
	"net/http"
	"strings"
)

// 请求成功 返回信息
func api_ret_json(writer http.ResponseWriter, request *http.Request, data []byte) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	writer.Header().Set("content-type", "application/json")             //返回数据格式是json
	_, _ = writer.Write(data)
}

// 请求失败 返回信息
func api_ret_err(writer http.ResponseWriter, request *http.Request, data []byte) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	writer.Header().Set("content-type", "text/html; charset=utf-8")     //返回数据格式是text
	writer.WriteHeader(201)                                             // 不是 200 就行， 201合适
	_, _ = writer.Write(data)
}

// 错误重定向
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// 检查用户提交的文本，至少有一个可见字符
func checkText(text []byte) (b bool) {
	b = false
	for _, v := range text {
		if v != ' ' && v != '\r' && v != '\n' && v != '\t' {
			b = true
			return
		}
	}
	return
}
