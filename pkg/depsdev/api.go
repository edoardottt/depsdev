/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

import (
	"fmt"
	"net/url"

	"github.com/edoardottt/depsdev/pkg/client"
)

type API struct {
	client *client.Client
}

// NewAPI creates and returns an API object.
func NewAPI() *API {
	return &API{
		client: client.New(BasePath),
	}
}

// GetInfo returns information about a package for a specific package manager.
func (a *API) GetInfo(packageManager, packageName string) (Package, error) {
	return getPackage(a.client, packageManager, packageName)
}

// getPackage returns a Package object.
func getPackage(c *client.Client, packageManager, packageName string) (Package, error) {
	var response Package

	var path = fmt.Sprintf(GetPackagePath, packageManager, url.PathEscape(packageName))
	if err := c.Get(path, &response); err != nil {
		return Package{}, err
	}

	return response, nil
}

// GetVersion returns information about a specific version of a package
// for a specific package manager.
func (a *API) GetVersion(packageManager, packageName, version string) (Version, error) {
	return getVersion(a.client, packageManager, packageName, version)
}

// getVersion returns a Version object.
func getVersion(c *client.Client, packageManager, packageName, version string) (Version, error) {
	var response Version

	var path = fmt.Sprintf(GetVersionPath, packageManager, url.PathEscape(packageName), version)
	if err := c.Get(path, &response); err != nil {
		return Version{}, err
	}

	return response, nil
}

// GetDependencies returns information about dependencies for a specific version of a package
// for a specific package manager.
func (a *API) GetDependencies(packageManager, packageName, version string) (Dependencies, error) {
	return getDependencies(a.client, packageManager, packageName, version)
}

// getDependencies returns a Dependencies object.
func getDependencies(c *client.Client, packageManager, packageName, version string) (Dependencies, error) {
	var response Dependencies

	var path = fmt.Sprintf(GetDependenciesPath, packageManager, url.PathEscape(packageName), version)
	if err := c.Get(path, &response); err != nil {
		return Dependencies{}, err
	}

	return response, nil
}

// GetProject returns information about a project (hosted on GitHub, GitLab or BitBucket).
func (a *API) GetProject(projectName string) (Project, error) {
	return getProject(a.client, projectName)
}

// getProject returns a Project object.
func getProject(c *client.Client, projectName string) (Project, error) {
	var response Project

	var path = fmt.Sprintf(GetProjectPath, url.PathEscape(projectName))
	if err := c.Get(path, &response); err != nil {
		return Project{}, err
	}

	return response, nil
}

// GetAdvisory returns information about an advisory.
func (a *API) GetAdvisory(advisory string) (Advisory, error) {
	return getAdvisory(a.client, advisory)
}

// getAdvisory returns an Advisory object.
func getAdvisory(c *client.Client, advisory string) (Advisory, error) {
	var response Advisory

	var path = fmt.Sprintf(GetAdvisoryPath, advisory)
	if err := c.Get(path, &response); err != nil {
		return Advisory{}, err
	}

	return response, nil
}

// GetQuery returns the result of the inputted query.
func (a *API) GetQuery(query string) (Package, error) {
	return getQuery(a.client, query)
}

// getQuery returns a Package object given a query.
func getQuery(c *client.Client, query string) (Package, error) {
	var response Package

	var path = GetQueryPath + `?` + query
	if err := c.Get(path, &response); err != nil {
		return Package{}, err
	}

	return response, nil
}
