package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

//User represents a user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

//Prepare calls methods to format the user
func (user *User) Prepare(step string) error {
	if error := user.validate(step); error != nil {
		return error
	}
	if error := user.format(step); error != nil {
		return error
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Name can't be empty ")
	}

	if user.Nick == "" {
		return errors.New("Nick can't be empty ")
	}

	if user.Email == "" {
		return errors.New("Email can't be empty ")
	}

	if error := checkmail.ValidateFormat(user.Email); error != nil {
		return error
	}

	if step == "register" && user.Password == "" {
		return errors.New("Password can't be empty ")
	}
	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	if step == "register" {
		passwordWithHash, error := security.Hash(user.Password)
		if error != nil {
			return error
		}
		user.Password = string(passwordWithHash)
	}
	return nil
}
