/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

import "time"

type Version struct {
	VersionKey     VersionKey       `json:"versionKey,omitempty"`
	PublishedAt    time.Time        `json:"publishedAt,omitempty"`
	IsDefault      bool             `json:"isDefault"`
	Licenses       []string         `json:"licenses,omitempty"`
	AdvisoryKeys   []AdvisoryKey    `json:"advisoryKeys,omitempty"`
	Links          []Link           `json:"links,omitempty"`
	SlsaProvenance []SlsaProvenance `json:"slsaProvenances"`
	Attestation    []Attestation    `json:"attestations"`
	Registries     []string         `json:"registries,omitempty"`
	RelatedProject []RelatedProject `json:"relatedProjects,omitempty"`
}

type Link struct {
	Label string `json:"label,omitempty"`
	URL   string `json:"url,omitempty"`
}

type Attestation struct {
	Type             string `json:"type,omitempty"`
	URL              string `json:"url,omitempty"`
	Verified         bool   `json:"verified,omitempty"`
	SourceRepository string `json:"sourceRepository,omitempty"`
	Commit           string `json:"commit,omitempty"`
}

type RelatedProject struct {
	ProjectKey         ProjectKey `json:"projectKey"`
	RelationProvenance string     `json:"relationProvenance,omitempty"`
	RelationType       string     `json:"relationType,omitempty"`
}

type VersionBatchRequest struct {
	PackageManager string `json:"system,omitempty"`
	PackageName    string `json:"name,omitempty"`
	Version        string `json:"version,omitempty"`
}

type SlsaProvenance struct {
	URL              string `json:"url,omitempty"`
	Verified         bool   `json:"verified,omitempty"`
	SourceRepository string `json:"sourceRepository,omitempty"`
	Commit           string `json:"commit,omitempty"`
}
