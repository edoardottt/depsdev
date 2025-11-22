/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

type Package struct {
	PackageKey PackageKey `json:"packageKey,omitempty"`
	Versions   []Version  `json:"versions,omitempty"`
	Purl       string     `json:"purl,omitempty"`
}

type PackageKey struct {
	System string `json:"system,omitempty"`
	Name   string `json:"name,omitempty"`
}

type VersionKey struct {
	System  string `json:"system,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}
