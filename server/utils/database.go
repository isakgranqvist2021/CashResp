package utils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	// read .env file in root directory
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

func GetKeys() (map[string]map[string]string, error) {
	bytes, err := ioutil.ReadFile("./store/keys.json")

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	data := make(map[string]map[string]string)
	if err := json.Unmarshal(bytes, &data); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	return data, nil
}
