/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

type Dependencies struct {
	Nodes []Nodes `json:"nodes,omitempty"`
	Edges []Edges `json:"edges,omitempty"`
	Error string  `json:"error,omitempty"`
}

type Nodes struct {
	VersionKey VersionKey `json:"versionKey,omitempty"`
	Bundled    bool       `json:"bundled,omitempty"`
	Errors     []string   `json:"errors,omitempty"`
}

type Edges struct {
	FromNode    int    `json:"fromNode,omitempty"`
	ToNode      int    `json:"toNode,omitempty"`
	Requirement string `json:"requirement,omitempty"`
}
