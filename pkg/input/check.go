/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package input

import "strings"

var (
	AllValidPackageManagers  = []string{"go", "rubygems", "npm", "cargo", "maven", "pypi", "nuget"}
	DepsValidPackageManagers = []string{"npm", "cargo", "maven", "pypi"}
	ReqsValidPackageManagers = []string{"npm", "nuget", "maven", "rubygems"}
)

// IsValidPackageManager checks if the inputted package name is valid.
func IsValidPackageManager(input string, validpackageManagers []string) bool {
	for _, packageName := range validpackageManagers {
		if strings.ToLower(input) == packageName {
			return true
		}
	}

	return false
}

func Contains(v string, sl []string) bool {
	for _, item := range sl {
		if strings.ToLower(v) == item {
			return true
		}
	}

	return false
}
