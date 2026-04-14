/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

type Requirements struct {
	Nuget    Nuget    `json:"nuget,omitempty"`
	Npm      Npm      `json:"npm,omitempty"`
	Maven    Maven    `json:"maven,omitempty"`
	RubyGems RubyGems `json:"rubygems,omitempty"`
	Go       Go       `json:"go,omitempty"`
	Pypi     Pypi     `json:"pypi,omitempty"`
	Cargo    Cargo    `json:"cargo,omitempty"`
}

// --- NuGet ---

type Nuget struct {
	DependencyGroups      []NugetDependencyGroup `json:"dependencyGroups,omitempty"`
	TargetFrameworks      []string               `json:"targetFrameworks,omitempty"`
	DevelopmentDependency bool                   `json:"developmentDependency,omitempty"`
	FrameworkAssemblies   []FrameworkAssembly    `json:"frameworkAssemblies,omitempty"`
	FrameworkReferences   []FrameworkReference   `json:"frameworkReferences,omitempty"`
}

type NugetDependencyGroup struct {
	TargetFramework string            `json:"targetFramework,omitempty"`
	Dependencies    []NugetDependency `json:"dependencies,omitempty"`
}

type NugetDependency struct {
	Name        string `json:"name,omitempty"`
	Requirement string `json:"requirement,omitempty"`
	Include     string `json:"include,omitempty"`
	Exclude     string `json:"exclude,omitempty"`
}

type FrameworkAssembly struct {
	AssemblyName    string `json:"assemblyName,omitempty"`
	TargetFramework string `json:"targetFramework,omitempty"`
}

type FrameworkReference struct {
	Name            string `json:"name,omitempty"`
	TargetFramework string `json:"targetFramework,omitempty"`
}

// --- NPM ---

type Npm struct {
	Dependencies NpmDependencyObject `json:"dependencies,omitempty"`
	Bundled      []NpmBundled        `json:"bundled,omitempty"`
	OS           []string            `json:"os,omitempty"`
	CPU          []string            `json:"cpu,omitempty"`
}

type NpmDependencyObject struct {
	Dependencies           []NpmDependency          `json:"dependencies,omitempty"`
	DevDependencies        []NpmDependency          `json:"devDependencies,omitempty"`
	OptionalDependencies   []NpmDependency          `json:"optionalDependencies,omitempty"`
	PeerDependencies       []NpmDependency          `json:"peerDependencies,omitempty"`
	BundleDependencies     []string                 `json:"bundleDependencies,omitempty"`
	PeerDependencyMetadata []PeerDependencyMetadata `json:"peerDependencyMetadata,omitempty"`
}

type NpmDependency struct {
	Name        string `json:"name,omitempty"`
	Requirement string `json:"requirement,omitempty"`
}

type PeerDependencyMetadata struct {
	Name     string `json:"name,omitempty"`
	Optional bool   `json:"optional,omitempty"`
}

type NpmBundled struct {
	Path         string              `json:"path,omitempty"`
	Name         string              `json:"name,omitempty"`
	Version      string              `json:"version,omitempty"`
	Dependencies NpmDependencyObject `json:"dependencies,omitempty"`
}

// --- Maven ---

type Maven struct {
	Parent               MavenParent       `json:"parent,omitempty"`
	Dependencies         []MavenDependency `json:"dependencies,omitempty"`
	DependencyManagement []MavenDependency `json:"dependencyManagement,omitempty"`
	Properties           []Property        `json:"properties,omitempty"`
	Repositories         []MavenRepository `json:"repositories,omitempty"`
	Profiles             []MavenProfile    `json:"profiles,omitempty"`
}

type MavenParent struct {
	System  string `json:"system,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type MavenDependency struct {
	Name            string   `json:"name,omitempty"`
	Version         string   `json:"version,omitempty"`
	Classifier      string   `json:"classifier,omitempty"`
	Type            string   `json:"type,omitempty"`
	Scope           string   `json:"scope,omitempty"`
	Optional        string   `json:"optional,omitempty"`
	Exclusions      []string `json:"exclusions,omitempty"`
	ResolvedVersion string   `json:"resolvedVersion,omitempty"`
	ResolvedName    string   `json:"resolvedName,omitempty"`
	Origin          string   `json:"origin,omitempty"`
}

type MavenProfile struct {
	ID                   string            `json:"id,omitempty"`
	Activation           Activation        `json:"activation,omitempty"`
	Dependencies         []MavenDependency `json:"dependencies,omitempty"`
	DependencyManagement []MavenDependency `json:"dependencyManagement,omitempty"`
	Properties           []Property        `json:"properties,omitempty"`
	Repositories         []MavenRepository `json:"repositories,omitempty"`
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
	ResolvedUrl      string `json:"resolvedUrl,omitempty"`
}

type Activation struct {
	ActiveByDefault string          `json:"activeByDefault,omitempty"`
	JDK             JDK             `json:"jdk,omitempty"`
	OS              OS              `json:"os,omitempty"`
	Property        PropertyWrapper `json:"property,omitempty"`
	File            File            `json:"file,omitempty"`
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

// --- RubyGems ---

type RubyGems struct {
	RuntimeDependencies     []RubyGemsDependency `json:"runtimeDependencies,omitempty"`
	DevDependencies         []RubyGemsDependency `json:"devDependencies,omitempty"`
	Platform                string               `json:"platform,omitempty"`
	RequiredRubyVersion     string               `json:"requiredRubyVersion,omitempty"`
	RequiredRubygemsVersion string               `json:"requiredRubygemsVersion,omitempty"`
}

type RubyGemsDependency struct {
	Name        string `json:"name,omitempty"`
	Requirement string `json:"requirement,omitempty"`
}

// --- Go ---

type Go struct {
	DirectDependencies   []GoDependency `json:"directDependencies,omitempty"`
	IndirectDependencies []GoDependency `json:"indirectDependencies,omitempty"`
	Replaces             []GoReplace    `json:"replaces,omitempty"`
	Excludes             []GoExclude    `json:"excludes,omitempty"`
}

type GoDependency struct {
	Name        string `json:"name,omitempty"`
	Requirement string `json:"requirement,omitempty"`
}

type GoReplace struct {
	Src         GoDependency `json:"src,omitempty"`
	Replacement GoDependency `json:"replacement,omitempty"`
	LocalPath   string       `json:"localPath,omitempty"`
}

type GoExclude struct {
	Name        string `json:"name,omitempty"`
	Requirement string `json:"requirement,omitempty"`
}

// --- PyPI ---

type Pypi struct {
	Dependencies          []PypiDependency `json:"dependencies,omitempty"`
	ProvidedExtras        []ProvidedExtra  `json:"providedExtras,omitempty"`
	ExternalDependencies  []PypiDependency `json:"externalDependencies,omitempty"`
	RequiredPythonVersion string           `json:"requiredPythonVersion,omitempty"`
}

type PypiDependency struct {
	ProjectName       string `json:"projectName,omitempty"`
	Extras            string `json:"extras,omitempty"`
	VersionSpecifier  string `json:"versionSpecifier,omitempty"`
	EnvironmentMarker string `json:"environmentMarker,omitempty"`
}

type ProvidedExtra struct {
	Name string `json:"name,omitempty"`
}

// --- Cargo ---

type Cargo struct {
	Dependencies []CargoDependency `json:"dependencies,omitempty"`
	Features     []CargoFeature    `json:"features,omitempty"`
}

type CargoDependency struct {
	Name                string   `json:"name,omitempty"`
	Requirement         string   `json:"requirement,omitempty"`
	Kind                string   `json:"kind,omitempty"`
	Optional            bool     `json:"optional,omitempty"`
	PackageAlias        string   `json:"packageAlias,omitempty"`
	UsesDefaultFeatures bool     `json:"usesDefaultFeatures,omitempty"`
	Features            []string `json:"features,omitempty"`
	Target              string   `json:"target,omitempty"`
}

type CargoFeature struct {
	Name    string   `json:"name,omitempty"`
	Implies []string `json:"implies,omitempty"`
}
