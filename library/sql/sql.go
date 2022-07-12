package sql

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	// TODO: 熔断器
	// "go-common/library/net/netutil/breaker"

	"golang.org/x/net/trace"
)

type DB struct {
	write  *conn
	read   []*conn
	idx    int64
	master *DB
}

// conn database connection
type conn struct {
	*sql.DB
	// breaker breaker.Breaker
	// conf    *Config
}
type Tx struct {
	db *conn
	tx *sql.Tx
	t  trace.Trace
	// c      context.Context
	cancel func()
}

// Row row.
type Row struct {
	err error
	*sql.Row
	db     *conn
	query  string
	args   []interface{}
	t      trace.Trace
	cancel func()
}

var db *sql.DB = nil
var once = sync.Once{}

func InitDb() *sql.DB {
	username := os.Getenv("DB_USERNAME")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	db := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=true&loc=Local",
		username, password, host, db, "utf8")

	if conn, err := sql.Open("mysql", dsn); err != nil {
		panic(err.Error())
	} else {
		//set how long to wait for connection close, run `show global variables like '%timeout%';` on your db to get value wait_timeout, SetConnMaxLifetime should be smaller than wait_timeout
		conn.SetConnMaxLifetime(7 * time.Second)

		//db.SetMaxIdleConns() is recommended to be set same to db.SetMaxOpenConns(). When it is smaller than SetMaxOpenConns(), connections can be opened and closed much more frequently than you expect
		conn.SetMaxOpenConns(10)
		conn.SetMaxIdleConns(10)
		defer conn.Close()
		return conn
	}
}

func GetDb() *sql.DB {
	if db == nil {
		once.Do(func() { InitDb() })
	}
	return db
}
