package main

import (
	"database/sql"
	_ "github.com/lib/pq" // 添加 Postgres支持
	_ "github.com/go-sql-driver/mysql" // 添加 mysql 支持
)
func main()  {
	db, err = sql.Open("postgres", dbname) // ok
	db, err = sql.Open("mysql", dbname) // ok
	db, err = sql.Open("sqlite3", dbname) // unknown driver sqlite3
}

