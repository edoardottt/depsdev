/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

type Advisory struct {
	AdvisoryKey AdvisoryKey `json:"advisoryKey,omitempty"`
	URL         string      `json:"url,omitempty"`
	Title       string      `json:"title,omitempty"`
	Aliases     []string    `json:"aliases,omitempty"`
	Cvss3Score  float64     `json:"cvss3Score"`
	Cvss3Vector string      `json:"cvss3Vector,omitempty"`
}

type AdvisoryKey struct {
	ID string `json:"id,omitempty"`
}
