package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
}

type JWTToken struct {
	Name   string
	Token  string
	Expire time.Time
}
