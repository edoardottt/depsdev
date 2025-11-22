/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

type Requirements struct {
	Nuget    *Nuget          `json:"nuget,omitempty"`
	Npm      *Npm            `json:"npm,omitempty"`
	Maven    *Maven          `json:"maven,omitempty"`
	RubyGems *DependencyList `json:"rubygems,omitempty"`
}

type Nuget struct {
	DependependencyGroups []DependependencyGroup `json:"dependencyGroups,omitempty"`
}

type DependependencyGroup struct {
	TargetFramework string       `json:"targetFramework,omitempty"`
	Dependencies    []Dependency `json:"dependencies,omitempty"`
}

type Dependency struct {
	Name        string   `json:"name,omitempty"`
	Requirement string   `json:"requirement,omitempty"`
	Version     string   `json:"version,omitempty"`
	System      string   `json:"system,omitempty"`
	Classifier  string   `json:"classifier,omitempty"`
	Type        string   `json:"type,omitempty"`
	Scope       string   `json:"scope,omitempty"`
	Optional    string   `json:"optional,omitempty"`
	Exclusions  []string `json:"exclusions,omitempty"`
}

type Npm struct {
	Dependencies DependencyList `json:"dependencies,omitempty"`
	Bundled      []Bundled      `json:"bundled,omitempty"`
}

type Bundled struct {
	Path         string           `json:"path,omitempty"`
	Name         string           `json:"name,omitempty"`
	Version      string           `json:"version,omitempty"`
	Dependencies []DependencyList `json:"dependencies,omitempty"`
}

type DependencyList struct {
	Dependencies         []Dependency `json:"dependencies,omitempty"`
	DevDependencies      []Dependency `json:"devDependencies,omitempty"`
	OptionalDependencies []Dependency `json:"optionalDependencies,omitempty"`
	PeerDependencies     []Dependency `json:"peerDependencies,omitempty"`
	RuntimeDependencies  []Dependency `json:"runtimeDependencies,omitempty"`
	BundleDependencies   []string     `json:"bundleDependencies,omitempty"`
}

type Maven struct {
	ID                   string            `json:"id,omitempty"`
	Parent               Dependency        `json:"parent,omitempty"`
	Activation           Activation        `json:"activation,omitempty"`
	Dependencies         []Dependency      `json:"dependencies,omitempty"`
	DependencyManagement []Dependency      `json:"dependency_management,omitempty"`
	Properties           []Property        `json:"properties,omitempty"`
	Repositories         []MavenRepository `json:"repositories,omitempty"`
	Profiles             []Maven           `json:"profiles,omitempty"`
}

type Property struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type MavenRepository struct {
	ID               string `json:"id,omitempty"`
	URL              string `json:"url,omitempty"`
	Layout           string `json:"layout,omitempty"`
	ReleasesEnabled  string `json:"releasesEnabled,omitempty"`
	SnapshotsEnabled string `json:"snapshotsEnabled,omitempty"`
}

type Activation struct {
	JDK             JDK             `json:"jdk,omitempty"`
	OS              OS              `json:"os,omitempty"`
	Property        PropertyWrapper `json:"property,omitempty"`
	File            File            `json:"file,omitempty"`
	ActiveByDefault string          `json:"activeByDefault,omitempty"`
}

type JDK struct {
	JDK string `json:"jdk,omitempty"`
}

type OS struct {
	Name    string `json:"name,omitempty"`
	Family  string `json:"family,omitempty"`
	Arch    string `json:"arch,omitempty"`
	Version string `json:"version,omitempty"`
}

type PropertyWrapper struct {
	Property Property `json:"property,omitempty"`
}

type File struct {
	Exists  string `json:"exists,omitempty"`
	Missing string `json:"missing,omitempty"`
}
