/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

// Dependencies represents the resolved dependency graph for a package version.
type Dependencies struct {
	Nodes []Node `json:"nodes,omitempty"`
	Edges []Edge `json:"edges,omitempty"`
	Error string `json:"error,omitempty"`
}

// Node represents a single package version within the dependency graph.
type Node struct {
	VersionKey VersionKey `json:"versionKey,omitempty"`
	Bundled    bool       `json:"bundled,omitempty"`
	// Relation can be SELF, DIRECT, or INDIRECT.
	Relation string   `json:"relation,omitempty"`
	Errors   []string `json:"errors,omitempty"`
}

// Edge represents a connection between a "from" node and a "to" node.
type Edge struct {
	// FromNode is the index into the Nodes list declaring the dependency.
	FromNode int `json:"fromNode,omitempty"`
	// ToNode is the index into the Nodes list resolving the dependency.
	ToNode      int    `json:"toNode,omitempty"`
	Requirement string `json:"requirement,omitempty"`
}

// Dependent contains counts of packages that depend on a version.
type Dependent struct {
	DependentCount         int `json:"dependentCount,omitempty"`
	DirectDependentCount   int `json:"directDependentCount,omitempty"`
	IndirectDependentCount int `json:"indirectDependentCount,omitempty"`
}
