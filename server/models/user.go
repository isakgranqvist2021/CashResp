package models

import "time"

type User struct {
	ID            int       `json:"id"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	AuthType      string    `json:"authType"`
	VerifyCode    string    `json:"verifyCode"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	EmailVerified bool      `json:"emailVerified"`
	Admin         bool      `json:"admin"`
}

func (u *User) Register() error {

	return nil
}

func (u *User) Login() error {

	return nil
}
