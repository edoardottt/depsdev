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
	VersionKey          VersionKey           `json:"versionKey,omitempty"`
	PublishedAt         time.Time            `json:"publishedAt,omitempty"`
	IsDefault           bool                 `json:"isDefault"`
	IsDeprecated        bool                 `json:"IsDeprecated,omitempty"`
	Licenses            []string             `json:"licenses,omitempty"`
	LicenseDetails      []LicenseDetail      `json:"licenseDetails,omitempty"`
	AdvisoryKeys        []AdvisoryKey        `json:"advisoryKeys,omitempty"`
	Links               []Link               `json:"links,omitempty"`
	Purl                string               `json:"purl,omitempty"`
	SlsaProvenance      []SlsaProvenance     `json:"slsaProvenances"`
	Attestations        []Attestation        `json:"attestations"`
	Registries          []string             `json:"registries,omitempty"`
	RelatedProjects     []RelatedProject     `json:"relatedProjects,omitempty"`
	UpstreamIdentifiers []UpstreamIdentifier `json:"upstreamIdentifiers,omitempty"`
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

type SlsaProvenance struct {
	SourceRepository string `json:"sourceRepository,omitempty"`
	Commit           string `json:"commit,omitempty"`
	URL              string `json:"url,omitempty"`
	Verified         bool   `json:"verified,omitempty"`
}

type LicenseDetail struct {
	License string `json:"license,omitempty"`
	Spdx    string `json:"spdx,omitempty"`
}

type UpstreamIdentifier struct {
	PackageName   string `json:"packageName,omitempty"`
	VersionString string `json:"versionString,omitempty"`
	Source        string `json:"source,omitempty"`
}
