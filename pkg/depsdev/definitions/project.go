/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

import "time"

type Project struct {
	ProjectKey      ProjectKey `json:"projectKey,omitempty"`
	OpenIssuesCount int        `json:"openIssuesCount,omitempty"`
	StarsCount      int        `json:"starsCount,omitempty"`
	ForksCount      int        `json:"forksCount,omitempty"`
	License         string     `json:"license,omitempty"`
	Description     string     `json:"description,omitempty"`
	Homepage        string     `json:"homepage,omitempty"`
	Scorecard       Scorecard  `json:"scorecard,omitempty"`
	OssFuzz         OssFuzz    `json:"ossFuzz,omitempty"`
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

type OssFuzz struct {
	LineCount        string    `json:"lineCount,omitempty"`
	LineCoverCount   string    `json:"lineCoverCount,omitempty"`
	LineCoverPercent string    `json:"lineCoverPercent,omitempty"`
	Date             time.Time `json:"date,omitempty"`
	ConfigURL        string    `json:"configUrl,omitempty"`
}
