package depsdev

type Dependencies struct {
	Nodes []Nodes `json:"nodes,omitempty"`
	Edges []Edges `json:"edges,omitempty"`
	Error string  `json:"error,omitempty"`
}

type Nodes struct {
	VersionKey VersionKey `json:"versionKey,omitempty"`
	Bundled    bool       `json:"bundled,omitempty"`
	Errors     []any      `json:"errors,omitempty"`
}

type Edges struct {
	FromNode    int    `json:"fromNode,omitempty"`
	ToNode      int    `json:"toNode,omitempty"`
	Requirement string `json:"requirement,omitempty"`
}
