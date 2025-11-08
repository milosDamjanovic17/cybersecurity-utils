package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"moose.com/web-server-quick-audit/internal/audit"
)

func main() {
	pathFlag := flag.String("path", "", "Custom path to web server configuration directory") // --path=/my/custom/path/to/web-server
	flag.Parse()

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
		audit.AuditNginx(*pathFlag)
		// Placeholder for NGINX audit logic
	case "a":
		audit.AuditApache(*pathFlag)
		// Placeholder for Apache audit logic
	default:
		fmt.Println("Invalid selection. Please choose either 'nginx' or 'apache'.")
	}
}