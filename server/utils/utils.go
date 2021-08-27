package utils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func LongEnough(str string, minLength, maxLength int) bool {
	strLen := len(str)

	return strLen >= minLength && strLen <= maxLength
}

func RandKey(n int, ints bool) string {
	var runes []string
	var str string

	if !ints {
		runes = strings.Split("qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM", "")
	} else {
		runes = strings.Split("0123456789", "")
	}

	for i := 0; i < n; i++ {
		str += runes[rand.Intn(len(runes))]
	}

	return str
}

func ConsumeAlert(c *fiber.Ctx) Alert {
	session, err := Store.Get(c)

	if err != nil {
		fmt.Println(err)
	}

	alert, OK := session.Get("alert").(Alert)

	if !OK {
		return Alert{}
	}

	session.Delete("alert")
	if err := session.Save(); err != nil {
		fmt.Println(err)
	}

	return alert
}

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
