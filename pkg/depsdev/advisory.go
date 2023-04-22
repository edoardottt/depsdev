package depsdev

type Advisory struct {
	AdvisoryKey AdvisoryKey `json:"advisoryKey,omitempty"`
	URL         string      `json:"url,omitempty"`
	Title       string      `json:"title,omitempty"`
	Aliases     []string    `json:"aliases,omitempty"`
	Cvss3Score  float64     `json:"cvss3Score,omitempty"`
	Cvss3Vector string      `json:"cvss3Vector,omitempty"`
}

type AdvisoryKey struct {
	ID string `json:"id,omitempty"`
}
