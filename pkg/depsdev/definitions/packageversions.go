/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

type PackageVersions struct {
	Versions []PackageVersion `json:"versions,omitempty"`
}

// PackageVersion is a package version built from the source code of a project.
type PackageVersion struct {
	VersionKey VersionKey `json:"versionKey,omitempty"`
	// RelationType can be SOURCE_REPO or ISSUE_TRACKER.
	RelationType string `json:"relationType,omitempty"`
	// RelationProvenance can be SLSA_ATTESTATION, GO_ORIGIN, PYPI_PUBLISH_ATTESTATION,
	// RUBYGEMS_PUBLISH_ATTESTATION or UNVERIFIED_METADATA.
	RelationProvenance string           `json:"relationProvenance,omitempty"`
	SlsaProvenance     []SlsaProvenance `json:"slsaProvenances"`
	Attestations       []Attestation    `json:"attestations"`
}
