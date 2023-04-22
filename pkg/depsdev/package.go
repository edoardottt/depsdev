package depsdev

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
	VersionKey VersionKey `json:"versionKey,omitempty"`
	IsDefault  bool       `json:"isDefault,omitempty"`
}
