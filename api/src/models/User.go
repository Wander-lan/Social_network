package models

import (
	"api/src/safety"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(operationType string) error {
	if err := user.validate(operationType); err != nil {
		return err
	}

	if err := user.format(operationType); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(operationType string) error {
	if user.Name == "" {
		return errors.New("Name is mandatory")
	}

	if user.Nick == "" {
		return errors.New("Nick is mandatory")
	}

	if user.Email == "" {
		return errors.New("Email is mandatory")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("The given email is invalid")
	}

	if operationType == "create" && user.Password == "" {
		return errors.New("Password is mandatory")
	}

	return nil
}

func (user *User) format(operationType string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if operationType == "create" {
		hashPassword, err := safety.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPassword)
	}

	return nil
}
