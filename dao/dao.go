package dao

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	db *sql.DB
}

var dao *Dao = nil
var once = sync.Once{}

func GetDao() *Dao {
	if dao == nil {
		once.Do(func() { InitDao() })
	}
	return dao
}

func InitDao() {
	username := os.Getenv("DB_USERNAME")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	db := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=true&loc=Local",
		username, password, host, db, "utf8")

	if conn, err := sql.Open("mysql", dsn); err != nil {
		fmt.Println("----error conn---", conn, err)
		panic(err.Error())
	} else {
		fmt.Println("----good conn---", conn)
		//set how long to wait for connection close, run `show global variables like '%timeout%';` on your db to get value wait_timeout, SetConnMaxLifetime should be smaller than wait_timeout
		conn.SetConnMaxLifetime(28800 * time.Second)

		//db.SetMaxIdleConns() is recommended to be set same to db.SetMaxOpenConns(). When it is smaller than SetMaxOpenConns(), connections can be opened and closed much more frequently than you expect
		conn.SetMaxOpenConns(10)
		conn.SetMaxIdleConns(10)
		// defer conn.Close()
		dao = &Dao{
			db: conn,
		}
	}
}

// Close close the resource.
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
