package utils

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	URI := fmt.Sprintf("%s:%s@/%s",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE"))

	db, err := sql.Open("mysql", URI)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	/*
		always call defer db.Close() when using database
	*/

	return db
}
