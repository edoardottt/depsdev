package depsdev

import "time"

type Project struct {
	ProjectKey      ProjectKey `json:"projectKey,omitempty"`
	OpenIssuesCount string     `json:"openIssuesCount,omitempty"`
	StarsCount      string     `json:"starsCount,omitempty"`
	ForksCount      string     `json:"forksCount,omitempty"`
	License         string     `json:"license,omitempty"`
	Description     string     `json:"description,omitempty"`
	Homepage        string     `json:"homepage,omitempty"`
	Scorecard       Scorecard  `json:"scorecard,omitempty"`
}

type ProjectKey struct {
	ID string `json:"id,omitempty"`
}

type Repository struct {
	Name   string `json:"name,omitempty"`
	Commit string `json:"commit,omitempty"`
}

type ScorecardReference struct {
	Version string `json:"version,omitempty"`
	Commit  string `json:"commit,omitempty"`
}

type Documentation struct {
	ShortDescription string `json:"shortDescription,omitempty"`
	URL              string `json:"url,omitempty"`
}

type Checks struct {
	Name          string        `json:"name,omitempty"`
	Documentation Documentation `json:"documentation,omitempty"`
	Score         string        `json:"score,omitempty"`
	Reason        string        `json:"reason,omitempty"`
	Details       []string      `json:"details,omitempty"`
}

type Scorecard struct {
	Repository   Repository         `json:"repository,omitempty"`
	Scorecard    ScorecardReference `json:"scorecard,omitempty"`
	Checks       []Checks           `json:"checks,omitempty"`
	OverallScore float64            `json:"overallScore,omitempty"`
	Metadata     []string           `json:"metadata,omitempty"`
	Date         time.Time          `json:"date,omitempty"`
}
