package models

import (
	"errors"
	"strings"
	"time"
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

	user.format()
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
	if operationType == "create" && user.Password == "" {
		return errors.New("Password is mandatory")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
