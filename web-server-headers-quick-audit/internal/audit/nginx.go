package audit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AuditNginx(customPath string) {
	defaultPath := "/etc/nginx/"

	pathToCheck := defaultPath
	if customPath != "" {
		pathToCheck = customPath
	}

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
	// TODO: continue with NGINX audit logic here
}

func pathExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
