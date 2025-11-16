package models

import (
	"time"
)

type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"password,omitempty"` // omitempty for security
    CreatedAt time.Time `json:"created_at"`
}

type RegisterRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type AuthResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    User    *User  `json:"user,omitempty"`
}