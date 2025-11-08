package audit

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var defaultApachePath = "/etc/apache2/sites-enabled/"

var apacheChecks = map[string][]string{
	"SSLProtocol": {
		"TLSv1.2", "TLSv1.3",
	},
	"SSLHonorCipherOrder": {
		"on",
	},
	"SSLSessionCache": {
		"shmcb",
	},
	"Header always set Strict-Transport-Security": {
		"max-age=31536000", "includeSubDomains", "preload",
	},
	"Header always set Referrer-Policy": {
		"strict-origin-when-cross-origin",
	},
	"Header always set Permissions-Policy": {
		"camera=()", "microphone=()", "geolocation=()",
	},
	"Header always set X-Content-Type-Options": {
		"nosniff",
	},
	"Header always set X-Frame-Options": {
		"SAMEORIGIN",
	},
}

func AuditApache(customPath string) {

	pathToCheck := defaultApachePath
	if customPath != "" {
		pathToCheck = customPath
	}

	if !pathExists(pathToCheck) {
		fmt.Printf("Apache path not found at '%s'.\n", pathToCheck)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the correct Apache path manually: ")
		input, _ := reader.ReadString('\n')
		pathToCheck = strings.TrimSpace(input)

		if !pathExists(pathToCheck) {
			fmt.Printf("The provided path '%s' does not exist.\n", pathToCheck)
			return
		}
	}

	fmt.Printf("Apache path verified at: %s\n", pathToCheck)

	fmt.Printf("\nChecking Apache configuration directory: %s\n", pathToCheck)

	if _, err := os.Stat(pathToCheck); os.IsNotExist(err) {
		fmt.Println("Path does not exist. Please provide a valid Apache sites-enabled directory.")
		return
	}

	files, err := filepath.Glob(filepath.Join(pathToCheck, "*"))
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No Apache files found in the specified directory.")
		return
	}

	for _, file := range files {
		fmt.Printf("\nAuditing file: %s\n", file)
		auditApacheFile(file)
	}
	// TODO: continue with Apache audit logic here
}


func auditApacheFile(filePath string) {
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

	for directive, expectedValues := range apacheChecks {
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
			fmt.Printf("%s OK - values => %s\n", directive, expectedValues)
		}
	}
}
