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

// GetInfo returns information about a package for a specific package manager.
func GetInfo(packageManager, packageName string) (Package, error) {
	c := client.New(BasePath)

	p, err := getPackage(c, packageManager, packageName)
	if err != nil {
		return Package{}, err
	}

	return p, nil
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
func GetVersion(packageManager, packageName, version string) (Version, error) {
	c := client.New(BasePath)

	v, err := getVersion(c, packageManager, packageName, version)
	if err != nil {
		return Version{}, err
	}

	return v, nil
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
func GetDependencies(packageManager, packageName, version string) (Dependencies, error) {
	c := client.New(BasePath)

	d, err := getDependencies(c, packageManager, packageName, version)
	if err != nil {
		return Dependencies{}, err
	}

	return d, nil
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
func GetProject(projectName string) (Project, error) {
	c := client.New(BasePath)

	p, err := getProject(c, projectName)
	if err != nil {
		return Project{}, err
	}

	return p, nil
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
func GetAdvisory(advisory string) (Advisory, error) {
	c := client.New(BasePath)

	a, err := getAdvisory(c, advisory)
	if err != nil {
		return Advisory{}, err
	}

	return a, nil
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
func GetQuery(query string) (Package, error) {
	c := client.New(BasePath)

	q, err := getQuery(c, query)
	if err != nil {
		return Package{}, err
	}

	return q, nil
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
