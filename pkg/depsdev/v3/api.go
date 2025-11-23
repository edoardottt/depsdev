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

// GetDependencies returns information about dependencies for a specific version of a package
// for a specific package manager.
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

// GetProject returns information about a project (hosted on GitHub, GitLab or BitBucket).
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

// GetAdvisory returns information about an advisory.
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

// Query returns the result of the inputted query.
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
// Requirements are currently available for Maven, npm and NuGet.
func (a *APIv3) GetRequirements(packageManager, packageName, version string) (def.Requirements, error) {
	var response def.Requirements

	var path = fmt.Sprintf(GetRequirementsPath, packageManager, url.PathEscape(packageName), version)
	if err := a.client.Get(path, &response); err != nil {
		return def.Requirements{}, err
	}

	return response, nil
}

// GetPackageVersions returns the package versions which attest to being created from the specified
// source code repository (hosted on GitHub, GitLab or BitBucket).
// At most 1500 package versions are returned.
func (a *APIv3) GetPackageVersions(projectName string) (def.PackageVersions, error) {
	var response def.PackageVersions

	var path = fmt.Sprintf(GetProjectPackageVersionsPath, url.PathEscape(projectName))
	if err := a.client.Get(path, &response); err != nil {
		return def.PackageVersions{}, err
	}

	return response, nil
}
