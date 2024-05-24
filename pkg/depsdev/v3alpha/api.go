/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

import (
	"context"
	"fmt"
	"net/url"

	"github.com/edoardottt/depsdev/pkg/client"
	def "github.com/edoardottt/depsdev/pkg/depsdev/definitions"
	"github.com/edoardottt/depsdev/pkg/input"
)

type APIv3Alpha struct {
	client *client.Client
}

// NewV3AlphaAPI creates and returns a V3Alpha API object.
func NewV3AlphaAPI() *APIv3Alpha {
	return &APIv3Alpha{
		client: client.New(V3AlphaBasePath),
	}
}

// GetInfo returns information about a package for a specific package manager.
func (a *APIv3Alpha) GetInfo(packageManager, packageName string) (def.Package, error) {
	if !input.IsValidPackageManager(packageManager) {
		return def.Package{}, input.ErrInvalidPackageManager
	}

	return getPackage(a.client, packageManager, packageName)
}

// getPackage returns a Package object.
func getPackage(c *client.Client, packageManager, packageName string) (def.Package, error) {
	var response def.Package

	var path = fmt.Sprintf(GetPackagePath, packageManager, url.PathEscape(packageName))
	if err := c.Get(path, &response); err != nil {
		return def.Package{}, err
	}

	return response, nil
}

// GetVersion returns information about a specific version of a package
// for a specific package manager.
func (a *APIv3Alpha) GetVersion(packageManager, packageName, version string) (def.Version, error) {
	if !input.IsValidPackageManager(packageManager) {
		return def.Version{}, input.ErrInvalidPackageManager
	}

	return getVersion(a.client, packageManager, packageName, version)
}

// getVersion returns a Version object.
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
func (a *APIv3Alpha) GetDependencies(packageManager, packageName, version string) (def.Dependencies, error) {
	return getDependencies(a.client, packageManager, packageName, version)
}

// getDependencies returns a Dependencies object.
func getDependencies(c *client.Client, packageManager, packageName, version string) (def.Dependencies, error) {
	var response def.Dependencies

	var path = fmt.Sprintf(GetDependenciesPath, packageManager, url.PathEscape(packageName), version)
	if err := c.Get(path, &response); err != nil {
		return def.Dependencies{}, err
	}

	return response, nil
}

// GetProject returns information about a project (hosted on GitHub, GitLab or BitBucket).
func (a *APIv3Alpha) GetProject(projectName string) (def.Project, error) {
	return getProject(a.client, projectName)
}

// getProject returns a Project object.
func getProject(c *client.Client, projectName string) (def.Project, error) {
	var response def.Project

	var path = fmt.Sprintf(GetProjectPath, url.PathEscape(projectName))
	if err := c.Get(path, &response); err != nil {
		return def.Project{}, err
	}

	return response, nil
}

// GetAdvisory returns information about an advisory.
func (a *APIv3Alpha) GetAdvisory(advisory string) (def.Advisory, error) {
	return getAdvisory(a.client, advisory)
}

// getAdvisory returns an Advisory object.
func getAdvisory(c *client.Client, advisory string) (def.Advisory, error) {
	var response def.Advisory

	var path = fmt.Sprintf(GetAdvisoryPath, advisory)
	if err := c.Get(path, &response); err != nil {
		return def.Advisory{}, err
	}

	return response, nil
}

// Query returns the result of the inputted query.
func (a *APIv3Alpha) Query(query string) (def.Package, error) {
	return getQuery(a.client, query)
}

// getQuery returns a Package object given a query.
func getQuery(c *client.Client, query string) (def.Package, error) {
	var response def.Package

	var path = QueryPath + `?` + query
	if err := c.Get(path, &response); err != nil {
		return def.Package{}, err
	}

	return response, nil
}

// GetRequirements returns the requirements for a given version in a system-specific format.
// Requirements are currently available for Maven, npm and NuGet.
func (a *APIv3Alpha) GetRequirements(packageManager, packageName, version string) (def.Requirements, error) {
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
func (a *APIv3Alpha) GetPackageVersions(projectName string) (def.PackageVersions, error) {
	var response def.PackageVersions

	var path = fmt.Sprintf(GetProjectPackageVersionsPath, url.PathEscape(projectName))
	if err := a.client.Get(path, &response); err != nil {
		return def.PackageVersions{}, err
	}

	return response, nil
}

// GetVersionBatch retrieves a batch of versions for the given version batch requests.
// The response can be paginated, so the method returns an iterator that allows you to retrieve all the pages content sequentially.
func (a *APIv3Alpha) GetVersionBatch(req []def.VersionBatchRequest) (*Iterator[def.Version], error) {
	for _, r := range req {
		if !input.IsValidPackageManager(r.PackageManager) {
			return nil, input.ErrInvalidPackageManager
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	cIn := getVersionBatch(ctx, a.client, req)
	iter := Iterator[def.Version]{
		cIn:     cIn,
		item:    def.Version{},
		err:     nil,
		hasNext: true,
		cancel:  cancel,
	}

	return &iter, nil
}

type versionBatchResponse struct {
	Responses []struct {
		Version def.Version `json:"version"`
	} `json:"responses"`
	NextPageToken string `json:"nextPageToken"`
}

func getVersionBatch(ctx context.Context, c *client.Client, req []def.VersionBatchRequest) <-chan batchJob[def.Version] {
	cJob := make(chan batchJob[def.Version])

	go func() {
		defer close(cJob)

		requests := []map[string]def.VersionBatchRequest{}
		for _, r := range req {
			requests = append(requests, map[string]def.VersionBatchRequest{"versionKey": r})
		}

		body := map[string]any{
			"requests":  requests,
			"pageToken": "",
		}

		response := versionBatchResponse{
			NextPageToken: "first",
		}

		for response.NextPageToken != "" {
			if err := c.Post(GetVersionBatchPath, body, &response); err != nil {
				cJob <- batchJob[def.Version]{
					Err: err,
				}

				return
			}

			for _, r := range response.Responses {
				select {
				case <-ctx.Done():
					return
				default:
					cJob <- batchJob[def.Version]{
						Item: r.Version,
					}
				}
			}

			body["pageToken"] = response.NextPageToken
		}
	}()

	return cJob
}

// GetProjectBatch retrieves a batch of projects.
// The response can be paginated, so the method returns an iterator that allows you to retrieve all the pages content sequentially.
func (a *APIv3Alpha) GetProjectBatch(projectNames []string) (*Iterator[def.Project], error) {
	ctx, cancel := context.WithCancel(context.Background())
	cIn := getProjectBatch(ctx, a.client, projectNames)
	iter := Iterator[def.Project]{
		cIn:     cIn,
		item:    def.Project{},
		err:     nil,
		hasNext: true,
		cancel:  cancel,
	}

	return &iter, nil
}

type projectBatchResponse struct {
	Responses []struct {
		Project def.Project `json:"project"`
	} `json:"responses"`
	NextPageToken string `json:"nextPageToken"`
}

func getProjectBatch(ctx context.Context, c *client.Client, projectNames []string) <-chan batchJob[def.Project] {
	cJob := make(chan batchJob[def.Project])

	go func() {
		defer close(cJob)

		requests := []map[string]def.ProjectKey{}
		for _, n := range projectNames {
			requests = append(requests, map[string]def.ProjectKey{"projectKey": {ID: n}})
		}

		body := map[string]any{
			"requests":  requests,
			"pageToken": "",
		}

		response := projectBatchResponse{
			NextPageToken: "first",
		}

		for response.NextPageToken != "" {
			if err := c.Post(GetProjectBatchPath, body, &response); err != nil {
				cJob <- batchJob[def.Project]{
					Err: err,
				}
				
				return
			}

			for _, r := range response.Responses {
				select {
				case <-ctx.Done():
					return
				default:
					cJob <- batchJob[def.Project]{
						Item: r.Project,
					}
				}
			}

			body["pageToken"] = response.NextPageToken
		}
	}()

	return cJob
}
