package input

import "strings"

var (
	PackageManagers = []string{"npm", "go", "maven", "cargo", "pypi", "nuget"}
)

// IsValidPackageManager checks if the inputted package name is valid.
func IsValidPackageManager(input string) bool {
	for _, packageName := range PackageManagers {
		if strings.ToLower(input) == packageName {
			return true
		}
	}

	return false
}

// IsValidPackageName checks if the inputted package name is valid.
// TODO.
func IsValidPackageName(input string) bool {
	return true
}

// IsValidAdvisory checks if the inputted advisory name is valid.
// TODO.
func IsValidAdvisory(input string) bool {
	return true
}

// IsValidProject checks if the inputted project name is valid.
// TODO.
func IsValidProject(input string) bool {
	return true
}

// IsValidQuery checks if the inputted query is valid.
// TODO.
func IsValidQuery(input string) bool {
	return true
}
