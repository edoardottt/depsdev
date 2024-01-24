/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package output

import "fmt"

const (
	Version = `0.0.9`
	Banner  = `depsdev v` + Version + `
  > https://github.com/edoardottt/depsdev
  > @edoardottt, https://www.edoardoottavianelli.it/`
	ShortDescription = `CLI client for deps.dev API`
)

// PrintBanner prints the banner in stdout.
func PrintBanner() {
	fmt.Println(Banner)
}
