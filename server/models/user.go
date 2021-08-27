package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/isakgranqvist2021/surveys/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int
	Email         string
	Password      string
	AuthType      string
	VerifyCode    string
	CreatedAt     string
	UpdatedAt     string
	EmailVerified bool
}

func (u *User) Login() error {
	db := utils.Connect()
	defer db.Close()

	query := fmt.Sprintf("SELECT Password FROM users WHERE Email = '%s'", u.Email)
	row := db.QueryRow(query)

	var pw string

	if err := row.Scan(&pw); err != nil {
		return errors.New("scanning row failed")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(pw), []byte(u.Password)); err != nil {
		return errors.New("password does not match")
	}

	if err := u.PopulateFrom(fmt.Sprintf("SELECT * FROM users WHERE Email = '%s'", u.Email)); err != nil {
		return errors.New("an error occured while querying user")
	}

	if !u.EmailVerified {
		return errors.New("email has not been verified")
	}

	return nil
}

func (u *User) Register() error {
	db := utils.Connect()
	defer db.Close()

	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)

	if err != nil {
		return errors.New("could not hash password")
	}

	query := fmt.Sprintf(`
		INSERT INTO users (Email, Password, AuthType, VerifyCode) 
		VALUES('%s', '%s', '%s', '%s')`,
		u.Email, string(bytes), u.AuthType, u.VerifyCode)

	_, err = db.Exec(query)

	if err != nil {
		me, ok := err.(*mysql.MySQLError)

		if !ok {
			return errors.New("casting err -> mysqlErr failed")
		}

		switch me.Number {
		case 1062:
			return errors.New("email already in use.")
		case 1146:
			return errors.New("internal server error")
		}

		return errors.New("internal server error")
	}

	return nil
}

func (u *User) VerifyEmail() error {
	db := utils.Connect()
	defer db.Close()

	if err := u.PopulateFrom(fmt.Sprintf("SELECT * FROM users WHERE VerifyCode = '%s'", u.VerifyCode)); err != nil {
		return err
	}

	query := fmt.Sprintf(`
			UPDATE users
			SET VerifyCode = '%s', UpdatedAt = '%s', EmailVerified = 1
			WHERE VerifyCode = '%s'`, "", time.Now().Format("2006-01-02 15:04:05"), u.VerifyCode)

	_, err := db.Exec(query)

	if err != nil {
		return errors.New("error while performing update query")
	}

	return nil
}

// id exists on struct -> populate full struct
func (u *User) PopulateFrom(query string) error {
	db := utils.Connect()
	defer db.Close()

	row := db.QueryRow(query)

	if err := row.Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.AuthType,
		&u.VerifyCode,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.EmailVerified); err != nil {
		return errors.New("error while scanning user")
	}

	return nil
}

// email exists on struct -> populate full struct
// func (u *User) PopulateFromEmail() error {

// }
