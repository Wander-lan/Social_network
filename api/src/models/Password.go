package models

// Represents the request format of a password update
type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
