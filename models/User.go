package models

import (
	"fmt"
	"net/http"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (user *User) Bind(request *http.Request) error {
	if user.Name == "" {
		fmt.Errorf("name is empty")
	}
	return nil
}
