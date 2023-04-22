package depsdev

type Version struct {
	VersionKey   VersionKey     `json:"versionKey,omitempty"`
	IsDefault    bool           `json:"isDefault,omitempty"`
	Licenses     []string       `json:"licenses,omitempty"`
	AdvisoryKeys []AdvisoryKeys `json:"advisoryKeys,omitempty"`
	Links        []Links        `json:"links,omitempty"`
}

type AdvisoryKeys struct {
	ID string `json:"id,omitempty"`
}

type Links struct {
	Label string `json:"label,omitempty"`
	URL   string `json:"url,omitempty"`
}
