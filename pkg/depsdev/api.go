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
	"log"
	"net/url"

	"github.com/edoardottt/depsdev/pkg/client"
)

// InfoHandler is the info subcommand handler.
func InfoHandler(args []string) {
	c := client.New(BasePath)

	p, err := GetPackage(c, args[0], args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}

// GetPackage returns a Package object.
func GetPackage(c *client.Client, packageManager, packageName string) (Package, error) {
	var response Package

	var path = fmt.Sprintf(GetPackagePath, packageManager, url.QueryEscape(packageName))

	if err := c.Get(path, &response); err != nil {
		return Package{}, err
	}

	return response, nil
}

// VersionHandler is the info subcommand handler if version is specified.
func VersionHandler(args []string) {
	c := client.New(BasePath)

	v, err := GetVersion(c, args[0], args[1], args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}

// GetVersion returns a Version object.
func GetVersion(c *client.Client, packageManager, packageName, version string) (Version, error) {
	var response Version

	var path = fmt.Sprintf(GetVersionPath, packageManager, url.QueryEscape(packageName), version)

	if err := c.Get(path, &response); err != nil {
		return Version{}, err
	}

	return response, nil
}

// DepsHandler is the deps subcommand handler.
func DepsHandler(args []string) {
	c := client.New(BasePath)

	d, err := GetDependencies(c, args[0], args[1], args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)
}

// GetDependencies returns a Dependencies object.
func GetDependencies(c *client.Client, packageManager, packageName, version string) (Dependencies, error) {
	var response Dependencies

	var path = fmt.Sprintf(GetDependenciesPath, packageManager, url.QueryEscape(packageName), version)

	if err := c.Get(path, &response); err != nil {
		return Dependencies{}, err
	}

	return response, nil
}

// ProjectHandler is the project subcommand handler.
func ProjectHandler(args []string) {
	c := client.New(BasePath)

	p, err := GetProject(c, args[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}

// GetProject returns a Project object.
func GetProject(c *client.Client, projectName string) (Project, error) {
	var response Project

	var path = fmt.Sprintf(GetProjectPath, url.QueryEscape(projectName))
	if err := c.Get(path, &response); err != nil {
		return Project{}, err
	}

	return response, nil
}

// AdvisoryHandler is the advisory subcommand handler.
func AdvisoryHandler(args []string) {
	c := client.New(BasePath)

	a, err := GetAdvisory(c, args[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a)
}

// GetAdvisory returns an Advisory object.
func GetAdvisory(c *client.Client, advisory string) (Advisory, error) {
	var response Advisory

	var path = fmt.Sprintf(GetAdvisoryPath, advisory)
	if err := c.Get(path, &response); err != nil {
		return Advisory{}, err
	}

	return response, nil
}

// QueryHandler is the query subcommand handler.
func QueryHandler(args []string) {
	c := client.New(BasePath)

	q, err := GetQuery(c, args[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}

// GetQuery returns a Package object given a query.
func GetQuery(c *client.Client, query string) (Package, error) {
	var response Package

	var path = GetQueryPath + `?` + query
	if err := c.Get(path, &response); err != nil {
		return Package{}, err
	}

	return response, nil
}
