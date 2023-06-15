package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email: "test@test.com",
		Username: "test",
		Password: "password",
	}
}