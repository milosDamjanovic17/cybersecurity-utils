package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Web Server Security Audit ===")
	fmt.Println("Which web server are you using? Type 'n' for nginx, Type 'a' for Apache")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	server := strings.ToLower(strings.TrimSpace(input))

	switch server {
	case "n":
		auditNginx()
		// Placeholder for NGINX audit logic
	case "a":
		auditApache()
		// Placeholder for Apache audit logic
	default:
		fmt.Println("Invalid selection. Please choose either 'nginx' or 'apache'.")
	}
}

func auditNginx() {
	fmt.Println("You have chosen NGINX")
}

func auditApache() {
	fmt.Println("You have chosen Apache")
}
