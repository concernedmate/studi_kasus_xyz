package models

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var dbPool *sql.DB

func InitDb() error {
	config := mysql.Config{
		User:                 os.Getenv("MYSQL_UID"),
		Passwd:               os.Getenv("MYSQL_PWD"),
		Addr:                 os.Getenv("MYSQL_ADDR"),
		DBName:               os.Getenv("MYSQL_DB"),
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal("Error initalizing database: ", err)
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Error connecting to database: ", pingErr)
		return err
	}

	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)

	dbPool = db

	return nil
}
