package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

// 初始化 Db ，只在启动时候调用一次
func Setup(dbaddr, dbname, dbuser, dbpasswd string, l *log.Logger) {
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8", dbuser, dbpasswd, dbaddr, dbname)
	var err error
	Db, err = sql.Open("mysql", str)
	if err != nil {
		log.Fatalln(err)
	}

	// 更改内部日志输出
	logger = l

	// 测试连接
	_, err = Db.Query("select * from users where id = 0")
	if err != nil {
		logger.Fatalln("Connect database fail", str, err)
	}
	logger.SetPrefix("INFO ")
	logger.Println("Set up Database success!")
}
