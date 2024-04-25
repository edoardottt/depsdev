/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

Partially taken from https://github.com/google/deps.dev/blob/main/examples/go/dependencies_dot/main.go

*/

package output

import (
	"errors"
	"fmt"

	definitions "github.com/edoardottt/depsdev/pkg/depsdev/definitions"
)

var (
	ErrGraphDependencies = errors.New("error in dependencies: %s")
	ErrGraphNode         = errors.New("error in node: %s")
)

// GenerateGraph takes as input a Dependencies struct and
// returns a Graphviz compatible graph.
func GenerateGraph(d definitions.Dependencies) (string, error) {
	if d.Error != "" {
		return "", fmt.Errorf("%w: %s", ErrGraphDependencies, d.Error)
	}

	var result = ""

	for _, n := range d.Nodes {
		for _, e := range n.Errors {
			if e != "" {
				return "", fmt.Errorf("%w: %s", ErrGraphNode,
					n.VersionKey.System+" "+
						n.VersionKey.Name+" "+
						n.VersionKey.Version+
						": "+e)
			}
		}
	}

	result += "digraph {\n"

	for i, n := range d.Nodes {
		result += fmt.Sprintf("  %d [label=%q];\n", i, n.VersionKey.Name+"@"+n.VersionKey.Version)
	}

	for _, e := range d.Edges {
		result += fmt.Sprintf("  %d -> %d [label=%q];\n", e.FromNode, e.ToNode, e.Requirement)
	}

	result += "}\n"

	return result, nil
}
