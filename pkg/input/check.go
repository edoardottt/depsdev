package input

import "strings"

// IsValidPackageManager checks if the inputted package name is valid.
func IsValidPackageManager(input string) bool {
	var PackageManagers = []string{"npm", "go", "maven", "cargo", "pypi", "nuget"}
	for _, packageName := range PackageManagers {
		if strings.ToLower(input) == packageName {
			return true
		}
	}

	return false
}
