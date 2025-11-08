package audit

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const defaultNginxPath = "/etc/nginx/sites-enabled/"

var nginxHeadersChecks = map[string][]string{
	"ssl_protocols": {
		"TLSv1.2", "TLSv1.3",
	},
	"add_header Strict-Transport-Security": {
		"max-age=31536000", "includeSubDomains", "preload",
	},
	"add_header Referrer-Policy": {
		"strict-origin-when-cross-origin",
	},
	"add_header Permissions-Policy": {
		"camera=()", "microphone=()", "geolocation=()",
	},
	"add_header X-Content-Type-Options": {
		"nosniff",
	},
	"add_header X-Frame-Options": {
		"SAMEORIGIN",
	},
}

func AuditNginx(customPath string) {

	pathToCheck := defaultNginxPath
	if customPath != "" {
		pathToCheck = customPath
	}
	// first of all check if the path exists
	if !pathExists(pathToCheck) {
		fmt.Printf("NGINX path not found at '%s'.\n", pathToCheck)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the correct NGINX path manually: ")
		input, _ := reader.ReadString('\n')
		pathToCheck = strings.TrimSpace(input)

		if !pathExists(pathToCheck) {
			fmt.Printf("The provided path '%s' does not exist.\n", pathToCheck)
			return
		}
	}

	fmt.Printf("NGINX path verified at: %s\n", pathToCheck)

	// check the NGINX configuration directory
	fmt.Printf("\nChecking NGINX Configuration Directory:%s\n", pathToCheck)

	if _, err := os.Stat(pathToCheck); os.IsNotExist(err) {
		fmt.Println("Path does not exist. Please provide a valid NGINX sites-enabled directory.")
		return
	}

	files, err := filepath.Glob(filepath.Join(pathToCheck, "*"))
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No .conf files found in the specified directory.")
		return
	}

	for _, file := range files {
		fmt.Printf("\nAuditing file: %s\n", file)
		auditFile(file)
	}
	// TODO: continue with NGINX audit logic here
}

func pathExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func auditFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := []string{}

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	for directive, expectedValues := range nginxHeadersChecks {
		found := false
		missingValues := []string{}

		for _, line := range content {
			if strings.Contains(line, directive) {
				found = true
				for _, val := range expectedValues {
					if !strings.Contains(line, val) {
						missingValues = append(missingValues, val)
					}
				}
				break
			}
		}

		if !found {
			fmt.Printf("Missing directive: %s\n", directive)
		} else if len(missingValues) > 0 {
			fmt.Printf("Directive '%s' found but missing values: %v\n\t recommended: %s\n", directive, missingValues, expectedValues)
		} else {
			fmt.Printf("%s OK - values %s\n", directive, expectedValues)
		}
	}
}
