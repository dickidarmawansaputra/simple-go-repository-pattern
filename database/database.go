package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	// ini tidak di close, karna akan di pake di func test lain (hanya untuk belajar)
	// parseTime=true
	/*
		● Secara default, Driver MySQL untuk Golang akan melakukan query tipe data DATE, DATETIME,
		TIMESTAMP menjadi []byte / []uint8. Dimana ini bisa dikonversi menjadi String, lalu di parsing
		menjadi time.Time
		● Namun hal ini merepotkan jika dilakukan manual, kita bisa meminta Driver MySQL untuk Golang
		secara otomatis melakukan parsing dengan menambahkan parameter parseTime=true
	*/
	db, err := sql.Open("mysql", "dickids:rahasia@tcp(localhost:3306)/belajar-golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
