package handlers

import (
    _ "database/sql"
    "encoding/json"
    "log"
    "net/http"
    
    "moose.com/login-app-sqli-demo/models"
    "golang.org/x/crypto/bcrypt"
)

// SECURE REGISTRATION - Uses Prepared Statements
func (h *AuthHandler) RegisterSecure(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    var req models.RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Error hashing password: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    
    // SECURE: Using prepared statements
    query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
    
    _, err = h.DB.Exec(query, req.Username, req.Email, string(hashedPassword))
    if err != nil {
        log.Printf("Database error: %v", err)
        response := models.AuthResponse{
            Success: false,
            Message: "Registration failed",
        }
        json.NewEncoder(w).Encode(response)
        return
    }
    
    response := models.AuthResponse{
        Success: true,
        Message: "User registered successfully",
    }
    json.NewEncoder(w).Encode(response)
}

// SECURE LOGIN - Uses Prepared Statements
func (h *AuthHandler) LoginSecure(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    var req models.LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    // SECURE: Using prepared statements
    query := "SELECT id, username, email, password, created_at FROM users WHERE username = ?"
    
    var user models.User
    var hashedPassword string
    
    err := h.DB.QueryRow(query, req.Username).Scan(&user.ID, &user.Username, &user.Email, &hashedPassword, &user.CreatedAt)
    if err != nil {
        response := models.AuthResponse{
            Success: false,
            Message: "Invalid username or password",
        }
        json.NewEncoder(w).Encode(response)
        return
    }
    
    // Check password
    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
    if err != nil {
        response := models.AuthResponse{
            Success: false,
            Message: "Invalid username or password",
        }
        json.NewEncoder(w).Encode(response)
        return
    }
    
    user.Password = ""
    
    response := models.AuthResponse{
        Success: true,
        Message: "Login successful",
        User:    &user,
    }
    json.NewEncoder(w).Encode(response)
}