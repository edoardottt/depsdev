/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package input

import (
	"errors"
)

var (
	ErrArgumentLeast                        = errors.New("argument must be specified")
	ErrArgumentsLeast                       = errors.New("arguments must be specified")
	ErrInvalidPackageManager                = errors.New("invalid package manager specified")
	ErrInvalidPackageManagerForRequirements = errors.New("invalid package manager specified: requirements are currently available only for Maven, npm, RubyGems and NuGet")
)
