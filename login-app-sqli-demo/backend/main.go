package main

import (
    "log"
    "net/http"
    
    "moose.com/login-app-sqli-demo/config"
    "moose.com/login-app-sqli-demo/handlers"
    "moose.com/login-app-sqli-demo/middleware"
    
    "github.com/gorilla/mux"
)

func main() {
    // Initialize database
    db := config.NewDatabase()
    defer db.Close()
    
    // Initialize handlers
    authHandler := handlers.NewAuthHandler(db.Connection)
    
    // Setup router
    router := mux.NewRouter()
    
    // API routes
    api := router.PathPrefix("/api").Subrouter()
    
    // Vulnerable endpoints (for demonstration)
    api.HandleFunc("/register-vulnerable", authHandler.RegisterVulnerable).Methods("POST")
    api.HandleFunc("/login-vulnerable", authHandler.LoginVulnerable).Methods("POST")
    api.HandleFunc("/users-vulnerable", authHandler.GetAllUsersVulnerable).Methods("GET")
    
    // Secure endpoints
    api.HandleFunc("/register-secure", authHandler.RegisterSecure).Methods("POST")
    api.HandleFunc("/login-secure", authHandler.LoginSecure).Methods("POST")
    
    // Apply CORS middleware
    handler := middleware.EnableCORS(router)
    
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}