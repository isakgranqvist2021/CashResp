package utils

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql"
	"github.com/joho/godotenv"
)

var Store *session.Store

type Alert struct {
	Severity string
	Message  string
}

func CreateStore() {
	// read .env file in root directory
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	storage := mysql.New(mysql.Config{
		Host:       "localhost",
		Port:       3306,
		Username:   os.Getenv("DATABASE_USER"),
		Password:   os.Getenv("DATABASE_PASSWORD"),
		Database:   os.Getenv("DATABASE"),
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})

	Store = session.New(session.Config{ // cookie config
		Expiration: time.Hour * 12,
		Storage:    storage,
	})

	var data map[string]interface{}
	Store.RegisterType(data)
}

func GetStore() *session.Store {
	return Store
}
