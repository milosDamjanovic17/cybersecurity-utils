package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"moose.com/login-app-sqli-demo/models"
)

type AuthHandler struct {
	DB *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

// VULNERABLE REGISTRATION - Susceptible to SQL Injection
func (h *AuthHandler) RegisterVulnerable(w http.ResponseWriter, r *http.Request) {
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

	// VULNERABLE: Direct string concatenation - DON'T DO THIS!
	query := fmt.Sprintf("INSERT INTO users (username, email, password) VALUES ('%s', '%s', '%s')",
		req.Username, req.Email, string(hashedPassword))

	log.Printf("Executing query: %s", query) // For demonstration purposes

	_, err = h.DB.Exec(query)
	if err != nil {
		log.Printf("Database error: %v", err)
		response := models.AuthResponse{
			Success: false,
			Message: fmt.Sprintf("Registration failed: %v", err),
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

// VULNERABLE LOGIN - Susceptible to SQL Injection
func (h *AuthHandler) LoginVulnerable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// VULNERABLE: Direct string concatenation - DON'T DO THIS!
	// query := fmt.Sprintf("SELECT id, username, email, password, created_at FROM users WHERE username = '%s'", req.Username)
	query := fmt.Sprintf("SELECT id, username, email, password, created_at FROM users WHERE username = '%s'", req.Username)

	log.Printf("Executing query: %s", query)

	var user models.User
	var hashedPassword string


	err := h.DB.QueryRow(query).Scan(&user.ID, &user.Username, &user.Email, &hashedPassword, &user.CreatedAt)
	// err := h.DB.QueryRow(query).Scan()
	if err != nil {
		log.Printf("Database error: %v", err)
		response := models.AuthResponse{
			Success: false,
			Message: "Invalid username or password",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Printf("USERNAME: =>>>>>>>>>> %s", req.Username)

	// VULNERABLE LOGIC: Skip password check if SQL injection patterns detected
	isSQLInjectionUsername := strings.Contains(strings.ToLower(req.Username), " or ") ||
		strings.Contains(strings.ToLower(req.Username), "union") ||
		strings.Contains(req.Username, "'") && strings.Contains(req.Username, "=") ||
		strings.Contains(req.Username, "--")

	isSQLInjectionPassword := strings.Contains(strings.ToLower(req.Password), " or ") ||
		strings.Contains(strings.ToLower(req.Password), "union") ||
		strings.Contains(req.Password, "'") && strings.Contains(req.Password, "=") ||
		strings.Contains(req.Password, "--")

	fmt.Println(req.Password)

	if isSQLInjectionUsername || isSQLInjectionPassword {
		log.Printf("SQL Injection detected - bypassing password check!")
		// Don't check password for injections - SUPER VULNERABLE!
	} else {
		// Normal password check for legitimate logins
		log.Printf("SQL Injection not detected - proceeding with regular password check!")
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
		if err != nil {
			response := models.AuthResponse{
				Success: false,
				Message: "Invalid username or password",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// Don't return the password
	user.Password = ""

	response := models.AuthResponse{
		Success: true,
		Message: "Login successful",
		User:    &user,
	}
	json.NewEncoder(w).Encode(response)
}

// Endpoint to fetch all users (for demonstration of SQL injection impact)
func (h *AuthHandler) GetAllUsersVulnerable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := r.URL.Query().Get("username")

	var query string
	if username != "" {
		// VULNERABLE: Direct string concatenation
		query = fmt.Sprintf("SELECT id, username, email, created_at FROM users WHERE username = '%s'", username)
	} else {
		query = "SELECT id, username, email, created_at FROM users"
	}

	log.Printf("Executing query: %s", query)

	rows, err := h.DB.Query(query)
	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}
