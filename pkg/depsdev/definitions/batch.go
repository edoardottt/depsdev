/*
depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.

@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

type VersionBatchRequest struct {
	VersionKey VersionKey `json:"versionKey,omitempty"`
}

type VersionBatchBody struct {
	Requests  []VersionBatchRequest `json:"requests"`
	PageToken string                `json:"pageToken"`
}

func (p *VersionBatchBody) SetNextPageToken(token string) {
	p.PageToken = token
}

type VersionBatchResponse struct {
	Responses []struct {
		Version Version             `json:"version"`
		Request VersionBatchRequest `json:"request"`
	} `json:"responses"`
	NextPageToken string `json:"nextPageToken"`
}

func (p *VersionBatchResponse) GetNextPageToken() string {
	return p.NextPageToken
}
func (p *VersionBatchResponse) Items() []Version {
	l := make([]Version, 0, len(p.Responses))
	for _, r := range p.Responses {
		l = append(l, r.Version)
	}

	return l
}

type ProjectBatchRequest struct {
	ProjectKey ProjectKey `json:"projectKey,omitempty"`
}

type ProjectBatchBody struct {
	Requests  []ProjectBatchRequest `json:"requests"`
	PageToken string                `json:"pageToken"`
}

func (p *ProjectBatchBody) SetNextPageToken(token string) {
	p.PageToken = token
}

type ProjectBatchResponse struct {
	Responses []struct {
		Project Project    `json:"project"`
		Request ProjectKey `json:"request"`
	} `json:"responses"`
	NextPageToken string `json:"nextPageToken"`
}

func (p *ProjectBatchResponse) GetNextPageToken() string {
	return p.NextPageToken
}

func (p *ProjectBatchResponse) Items() []Project {
	l := make([]Project, 0, len(p.Responses))
	for _, r := range p.Responses {
		l = append(l, r.Project)
	}

	return l
}

type PurlBatchRequest struct {
	Purl string `json:"purl,omitempty"`
}

type PurlBatchBody struct {
	Requests  []PurlBatchRequest `json:"requests"`
	PageToken string             `json:"pageToken"`
}

func (p *PurlBatchBody) SetNextPageToken(token string) {
	p.PageToken = token
}

type PurlBatchResponse struct {
	Responses []struct {
		Purl    Purl             `json:"result"`
		Request PurlBatchRequest `json:"request"`
	} `json:"responses"`
	NextPageToken string `json:"nextPageToken"`
}

func (p *PurlBatchResponse) GetNextPageToken() string {
	return p.NextPageToken
}

func (p *PurlBatchResponse) Items() []Purl {
	l := make([]Purl, 0, len(p.Responses))
	for _, r := range p.Responses {
		l = append(l, r.Purl)
	}

	return l
}
