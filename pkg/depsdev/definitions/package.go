/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

import "time"

type Package struct {
	PackageKey PackageKey `json:"packageKey,omitempty"`
	Versions   []Versions `json:"versions,omitempty"`
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

type Versions struct {
	VersionKey         VersionKey        `json:"versionKey,omitempty"`
	IsDefault          bool              `json:"isDefault,omitempty"`
	PublishedAt        time.Time         `json:"publishedAt,omitempty"`
	SlsaProvenances    []SLSAProvenances `json:"slsaProvenances,omitempty"`
	RelationProvenance string            `json:"relationProvenance,omitempty"`
	RelationType       string            `json:"relationType,omitempty"`
}
