package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	//queryOne()
	queryMulti()
	//insertData()
	// updateData()
	// deleteData()
}

//封装链接数据库
func dbConn() (db *sql.DB) {
	const (
		username = "root"
		password = "zhang19960920..."
		network  = "tcp"
		host     = "47.95.225.57"
		prot     = 3306
		database = "users"
	)
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, password, network, host, prot, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	db.SetConnMaxLifetime(100 * time.Second)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return db
}

//查询单行
func queryOne() {
	db := dbConn()
	user := new(User)
	row := db.QueryRow("select * from user where id=?", 1)
	if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Println(*user)
}

//查询多行
func queryMulti() {
	db := dbConn()
	user := new(User)
	rows, err := db.Query("select * from user where id > ?", 0)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return
		}
		fmt.Print(*user)
	}

}

//插入数据
func insertData() {
	db := dbConn()
	result, err := db.Exec("insert INTO user(name,age) values(?,?)", "YDZ", 23)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return
	}
	fmt.Println("LastInsertID:", lastInsertID)
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

//更新数据
func updateData() {
	db := dbConn()
	result, err := db.Exec("UPDATE user set age=? where id=?", "30", 3)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

//删除数据
func deleteData() {
	db := dbConn()
	result, err := db.Exec("delete from user where id=?", 1)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}
