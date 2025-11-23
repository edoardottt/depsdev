/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

type ContainerImages struct {
	ContainerImages []ContainerImage `json:"results,omitempty"`
}

type ContainerImage struct {
	Registry   string `json:"registry,omitempty"`
	Repository string `json:"repository,omitempty"`
}
