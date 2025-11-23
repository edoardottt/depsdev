/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

import (
	"fmt"
	"net/url"

	"github.com/edoardottt/depsdev/pkg/client"
	def "github.com/edoardottt/depsdev/pkg/depsdev/definitions"
	"github.com/edoardottt/depsdev/pkg/input"
)

type APIv3 struct {
	client *client.Client
}

// NewV3API creates and returns a V3 API object.
func NewV3API() *APIv3 {
	return &APIv3{
		client: client.New(V3BasePath),
	}
}

// GetPackage returns information about a package, including a list of its available versions,
// with the default version marked if known.
func (a *APIv3) GetPackage(packageManager, packageName string) (def.Package, error) {
	if !input.IsValidPackageManager(packageManager, input.AllValidPackageManagers) {
		return def.Package{}, input.ErrInvalidPackageManager
	}

	return getPackage(a.client, packageManager, packageName)
}

func getPackage(c *client.Client, packageManager, packageName string) (def.Package, error) {
	var response def.Package

	var path = fmt.Sprintf(GetPackagePath, packageManager, url.PathEscape(packageName))
	if err := c.Get(path, &response); err != nil {
		return def.Package{}, err
	}

	return response, nil
}

// GetVersion returns information about a specific package version,
// including its licenses and any security advisories known to affect it.
func (a *APIv3) GetVersion(packageManager, packageName, version string) (def.Version, error) {
	if !input.IsValidPackageManager(packageManager, input.AllValidPackageManagers) {
		return def.Version{}, input.ErrInvalidPackageManager
	}

	return getVersion(a.client, packageManager, packageName, version)
}

func getVersion(c *client.Client, packageManager, packageName, version string) (def.Version, error) {
	var response def.Version

	var path = fmt.Sprintf(GetVersionPath, packageManager, url.PathEscape(packageName), version)
	if err := c.Get(path, &response); err != nil {
		return def.Version{}, err
	}

	return response, nil
}

// GetDependencies returns a resolved dependency graph for the given package version.
// Dependencies are currently available for npm, Cargo, Maven and PyPI.
// Dependencies are the resolution of the requirements (dependency constraints) specified by a version.
// The dependency graph should be similar to one produced by installing the package version on a generic 64-bit Linux system,
// with no other dependencies present. The precise meaning of this varies from system to system.
// Example: npm react 18.2.0.
func (a *APIv3) GetDependencies(packageManager, packageName, version string) (def.Dependencies, error) {
	return getDependencies(a.client, packageManager, packageName, version)
}

func getDependencies(c *client.Client, packageManager, packageName, version string) (def.Dependencies, error) {
	var response def.Dependencies

	var path = fmt.Sprintf(GetDependenciesPath, packageManager, url.PathEscape(packageName), version)
	if err := c.Get(path, &response); err != nil {
		return def.Dependencies{}, err
	}

	return response, nil
}

// GetProject returns information about projects hosted by GitHub, GitLab, or BitBucket, when known to us.
// Example: github.com/facebook/react.
func (a *APIv3) GetProject(projectName string) (def.Project, error) {
	return getProject(a.client, projectName)
}

func getProject(c *client.Client, projectName string) (def.Project, error) {
	var response def.Project

	var path = fmt.Sprintf(GetProjectPath, url.PathEscape(projectName))
	if err := c.Get(path, &response); err != nil {
		return def.Project{}, err
	}

	return response, nil
}

// GetAdvisory returns information about security advisories hosted by OSV.
// Example: GHSA-2qrg-x229-3v8q.
func (a *APIv3) GetAdvisory(advisory string) (def.Advisory, error) {
	return getAdvisory(a.client, advisory)
}

func getAdvisory(c *client.Client, advisory string) (def.Advisory, error) {
	var response def.Advisory

	var path = fmt.Sprintf(GetAdvisoryPath, advisory)
	if err := c.Get(path, &response); err != nil {
		return def.Advisory{}, err
	}

	return response, nil
}

// Query returns information about multiple package versions, which can be specified by name, content hash, or both.
// If a hash was specified in the request, it returns the artifacts that matched the hash.
// Querying by content hash is currently supported for npm, Cargo, Maven, NuGet, PyPI and RubyGems.
// It is typical for hash queries to return many results; hashes are matched against multiple release artifacts
// (such as JAR files) that comprise package versions, and any given artifact may appear in several package versions.
// Examples:
// hash.type=SHA1&hash.value=ulXBPXrC%2FUTfnMgHRFVxmjPzdbk%3D
// versionKey.system=NPM&versionKey.name=react&versionKey.version=18.2.0
// End of examples.
func (a *APIv3) Query(query string) (def.Results, error) {
	return getQuery(a.client, query)
}

func getQuery(c *client.Client, query string) (def.Results, error) {
	var response def.Results

	var path = QueryPath + `?` + query
	if err := c.Get(path, &response); err != nil {
		return def.Results{}, err
	}

	return response, nil
}

// GetRequirements returns the requirements for a given version in a system-specific format.
// Requirements are currently available for Maven, npm, NuGet and RubyGems.
// Requirements are the dependency constraints specified by the version.
// Example: nuget castle.core 5.1.1.
func (a *APIv3) GetRequirements(packageManager, packageName, version string) (def.Requirements, error) {
	var response def.Requirements

	var path = fmt.Sprintf(GetRequirementsPath, packageManager, url.PathEscape(packageName), version)
	if err := a.client.Get(path, &response); err != nil {
		return def.Requirements{}, err
	}

	return response, nil
}

// GetProjectPackageVersions returns known mappings between the requested project and package versions.
// At most 1500 package versions are returned.
// Mappings which were derived from attestations are served first.
// Example: github.com/facebook/react.
func (a *APIv3) GetProjectPackageVersions(projectName string) (def.PackageVersions, error) {
	var response def.PackageVersions

	var path = fmt.Sprintf(GetProjectPackageVersionsPath, url.PathEscape(projectName))
	if err := a.client.Get(path, &response); err != nil {
		return def.PackageVersions{}, err
	}

	return response, nil
}
