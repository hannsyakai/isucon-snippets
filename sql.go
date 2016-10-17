package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func initDB(db *sql.DB) error {
	dbUser := os.Getenv("ISU_DB_USER")
	dbPassword := os.Getenv("ISU_DB_PASSWORD")
	dbName := os.Getenv("ISU_DB_NAME")
	dbUnix := "/var/run/mysqld/mysqld.sock"

	dsn := fmt.Sprintf(
		"%s:%s@unix(%s)/%s?loc=Local&parseTime=true&interpolateParams=true&collation=utf8mb4_bin",
		dbUser, dbPassword, dbUnix, dbName)
	log.Println("dsn: ", dsn)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(16)
	db.SetMaxIdleConns(16)

	maxRetryCount := 5
	cnt := 0
	for {
		cnt++
		err := db.Ping()
		if err == nil {
			return nil
		}
		log.Println(err)
		if cnt >= maxRetryCount {
			return err
		}
		time.Sleep(time.Millisecond * 100)
	}
}
