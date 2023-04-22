/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

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
