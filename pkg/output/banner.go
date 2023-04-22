package output

import "fmt"

const (
	Version = `0.0.1`
	Banner  = `depsdev v` + Version + `
  > https://github.com/edoardottt/depsdev
  > @edoardottt, https://www.edoardoottavianelli.it/`
)

// PrintBanner prints the banner in stdout.
func PrintBanner() {
	fmt.Println(Banner)
}
