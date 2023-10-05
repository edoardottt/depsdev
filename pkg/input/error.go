/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package input

import (
	"errors"
	"fmt"
)

var (
	ErrArgumentLeast         = errors.New("argument must be specified")
	ErrArgumentsLeast        = errors.New("arguments must be specified")
	ErrInvalidPackageManager = errors.New("invalid package manager specified")
)

func ErrWithMessage(err error, msg string) error {
	return errors.New(fmt.Sprintf("%v: %v", err, msg))
}
