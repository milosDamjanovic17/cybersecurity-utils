package audit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AuditApache(customPath string) {
	defaultPath := "/etc/apache2/"

	pathToCheck := defaultPath
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
	// TODO: continue with Apache audit logic here
}
