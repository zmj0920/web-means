// 处理配置文件和日志
package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/zmj/gotest/data"
)

type Configuration struct {
	ListeningAddress string
	ReadTimeout      int64
	WriteTimeout     int64
	StaticPath       string // 静态文件目录
	TemplatePath     string // 模板文件目录
	DBname           string // 数据库名字
	DBaddress        string // 数据库地址
	DBuser           string
	DBpasswd         string
	LogFile          string // 日志文件
}

var Config *Configuration // 全局配置
var Logger *log.Logger    // 日志

func init() {
	// 初始化为默认值
	Config = &Configuration{
		ListeningAddress: "0.0.0.0:8080",
		ReadTimeout:      10,
		WriteTimeout:     600,
		StaticPath:       "public",
		TemplatePath:     "template",
		DBname:           "go_test",
		DBaddress:        "47.95.225.57:3306",
		DBuser:           "root",
		DBpasswd:         "zhang19960920...",
	}

	Logger = log.New(os.Stdout, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Cannot open config file:", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(Config)
	if err != nil {
		log.Fatalln("Cannot decode config:", err)
	}
}

func setup() {
	if len(Config.LogFile) > 0 { // 设置日志输出 默认为 stdout
		f, err := os.OpenFile(Config.LogFile, os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Fatalln(err)
		}
		Logger.SetOutput(f)
	}

	// 启动数据库连接
	data.Setup(Config.DBaddress, Config.DBname, Config.DBuser, Config.DBpasswd, Logger)
}
