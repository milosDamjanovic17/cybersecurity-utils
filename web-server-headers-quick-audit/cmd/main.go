package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Web Server Header Security Audit ===")
	// fmt.Println("=== Which webserver are you using? (nginx(n)/apache(a)) ===")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the website URL (e.g., https://example.com): ")
	// fmt.Print("> ")
	url, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	url = strings.TrimSpace(url)
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url // default this to HTTPS
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Failed to reach the site:", err)
		return
	}
	defer resp.Body.Close()

	serverHeader := resp.Header.Get("Server")
	if serverHeader == "" {
		fmt.Println("No 'Server' header detected.")
		return
	}

	serverLower := strings.ToLower(serverHeader)
	switch {
	case strings.Contains(serverLower, "nginx"):
		fmt.Println("Identified as NGINX.")
		// TODO: call auditNginx(resp.Header)
	case strings.Contains(serverLower, "apache"):
		fmt.Println("Identified as Apache.")
		// TODO: call auditApache(resp.Header)
	default:
		fmt.Println("Unknown or custom web server type. The script only checks for NGINX and Apache")
	}
	// switch server {
	// case "n":
	// 	fmt.Println("=== NGINX Chosen Selected ===")
	// 	// Placeholder for NGINX audit logic
	// case "a":
	// 	fmt.Println("=== Apache Selected ===")
	// default:
	// 	fmt.Println("Invalid selection. Please choose either NGINX(n) or Apache(a)")
	// }
}
