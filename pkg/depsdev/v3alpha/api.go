/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

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

// GetPackage returns information about a package, including a list of its available versions,
// with the default version marked if known.
func (a *APIv3Alpha) GetPackage(packageManager, packageName string) (def.Package, error) {
	if !input.IsValidPackageManager(packageManager, input.AllValidPackageManagers) {
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

// GetVersion returns information about a specific package version,
// including its licenses and any security advisories known to affect it.
func (a *APIv3Alpha) GetVersion(packageManager, packageName, version string) (def.Version, error) {
	if !input.IsValidPackageManager(packageManager, input.AllValidPackageManagers) {
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

// GetDependencies returns a resolved dependency graph for the given package version.
// Dependencies are currently available for npm, Cargo, Maven and PyPI.
// Dependencies are the resolution of the requirements (dependency constraints) specified by a version.
// The dependency graph should be similar to one produced by installing the package version on a generic 64-bit Linux system,
// with no other dependencies present. The precise meaning of this varies from system to system.
func (a *APIv3Alpha) GetDependencies(packageManager, packageName, version string) (def.Dependencies, error) {
	if !input.IsValidPackageManager(packageManager, input.DepsValidPackageManagers) {
		return def.Dependencies{}, input.ErrInvalidPackageManagerForDependencies
	}

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

// GetProject returns information about projects hosted by GitHub, GitLab, or BitBucket, when known to us.
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

// GetAdvisory returns information about security advisories hosted by OSV.
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

// Query returns information about multiple package versions, which can be specified by name, content hash, or both.
// If a hash was specified in the request, it returns the artifacts that matched the hash.
// Querying by content hash is currently supported for npm, Cargo, Maven, NuGet, PyPI, and RubyGems.
// It is typical for hash queries to return many results; hashes are matched against multiple release artifacts
// (such as JAR files) that comprise package versions, and any given artifact may appear in several package versions.
func (a *APIv3Alpha) Query(query string) (def.Results, error) {
	return getQuery(a.client, query)
}

// getQuery returns a Package object given a query.
func getQuery(c *client.Client, query string) (def.Results, error) {
	var response def.Results

	var path = QueryPath + `?` + query
	if err := c.Get(path, &response); err != nil {
		return def.Results{}, err
	}

	return response, nil
}

// GetRequirements returns the requirements for a given version in a system-specific format.
// Requirements are currently available for Maven, npm, NuGet, and RubyGems.
// Requirements are the dependency constraints specified by the version.
// Example: /v3alpha/systems/nuget/packages/castle.core/versions/5.1.1:requirements.
func (a *APIv3Alpha) GetRequirements(packageManager, packageName, version string) (def.Requirements, error) {
	if !input.IsValidPackageManager(packageManager, input.ReqsValidPackageManagers) {
		return def.Requirements{}, input.ErrInvalidPackageManager
	}

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
func (a *APIv3Alpha) GetPackageVersions(projectName string) (def.PackageVersions, error) {
	var response def.PackageVersions

	var path = fmt.Sprintf(GetProjectPackageVersionsPath, url.PathEscape(projectName))
	if err := a.client.Get(path, &response); err != nil {
		return def.PackageVersions{}, err
	}

	return response, nil
}

// GetDependents returns information about the number of distinct packages known to depend
// on the given package version. Dependent counts are currently available for npm, Cargo, Maven and PyPI.
//
// Dependent counts are derived from the dependency graphs computed by deps.dev, which means that only public dependents are counted.
// As such, dependent counts should be treated as indicative of relative popularity rather than precisely accurate.
func (a *APIv3Alpha) GetDependents(packageManager, packageName, version string) (def.Dependent, error) {
	if !input.IsValidPackageManager(packageManager, input.AllValidPackageManagers) {
		return def.Dependent{}, input.ErrInvalidPackageManager
	}

	var response def.Dependent

	var path = fmt.Sprintf(GetDependentsPath, packageManager, url.PathEscape(packageName), version)
	if err := a.client.Get(path, &response); err != nil {
		return def.Dependent{}, err
	}

	return response, nil
}

// GetCapabilityRequest returns counts for direct and indirect calls to Capslock capabilities for a given package version.
// Currently only available for Go.
func (a *APIv3Alpha) GetCapabilities(packageManager, packageName, version string) (def.Capabilities, error) {
	if !input.IsValidPackageManager(packageManager, input.AllValidPackageManagers) {
		return def.Capabilities{}, input.ErrInvalidPackageManager
	}

	var response def.Capabilities

	var path = fmt.Sprintf(GetCapabilitiesPath, packageManager, url.PathEscape(packageName), version)
	if err := a.client.Get(path, &response); err != nil {
		return def.Capabilities{}, err
	}

	return response, nil
}

// GetSimilarlyNamedPackages returns packages with names that are similar to the requested package.
// This similarity relation is computed by deps.dev.
func (a *APIv3Alpha) GetSimilarlyNamedPackages(packageManager, packageName string) (def.SimilarlyNamedPackages, error) {
	if !input.IsValidPackageManager(packageManager, input.AllValidPackageManagers) {
		return def.SimilarlyNamedPackages{}, input.ErrInvalidPackageManager
	}

	var response def.SimilarlyNamedPackages

	var path = fmt.Sprintf(GetSimilarlyNamedPackagesPath, packageManager, url.PathEscape(packageName))
	if err := a.client.Get(path, &response); err != nil {
		return def.SimilarlyNamedPackages{}, err
	}

	return response, nil
}

// GetVersionBatch performs GetVersion requests for a batch of versions.
// Large result sets may be paginated.
func (a *APIv3Alpha) GetVersionBatch(req def.VersionBatchBody) (*Iterator[def.Version], error) {
	for _, v := range req.Requests {
		if !input.IsValidPackageManager(v.VersionKey.System, input.AllValidPackageManagers) {
			return nil, input.ErrInvalidPackageManager
		}
	}

	response := &def.VersionBatchResponse{
		NextPageToken: "first",
	}

	ctx, cancel := context.WithCancel(context.Background())
	cIn := getBatch(ctx, a.client, GetVersionBatchPath, &req, response)

	iter := Iterator[def.Version]{
		cIn:     cIn,
		hasNext: true,
		cancel:  cancel,
	}

	return &iter, nil
}

// GetProjectBatch performs GetProjectBatch requests for a batch of projects.
// Large result sets may be paginated.
func (a *APIv3Alpha) GetProjectBatch(req def.ProjectBatchBody) (*Iterator[def.Project], error) {
	response := &def.ProjectBatchResponse{
		NextPageToken: "first",
	}

	ctx, cancel := context.WithCancel(context.Background())
	cIn := getBatch(ctx, a.client, GetProjectBatchPath, &req, response)

	iter := Iterator[def.Project]{
		cIn:     cIn,
		hasNext: true,
		cancel:  cancel,
	}

	return &iter, nil
}

// PurlLookup searches for a package or package version specified via purl, and returns the corresponding result
// from GetPackage or GetVersion as appropriate.
// For a package lookup, the purl should be in the form pkg:type/namespace/name for a namespaced package name,
// or pkg:type/name for a non-namespaced package name.
// For a package version lookup, the purl should be in the form pkg:type/namespace/name@version, or pkg:type/name@version.
// Extra fields in the purl must be empty, otherwise the request will fail. In particular, there must be no subpath or qualifiers.
// Supported values for type are cargo, gem, golang, maven, npm, nuget, and pypi. Further details on types,
// and how to form purls of each type, can be found in the purl spec.
// Special characters in purls must be percent-encoded. This is described in detail by the purl spec.
func (a *APIv3Alpha) PurlLookup(purl string) (def.Purl, error) {
	var response def.Purl

	var path = fmt.Sprintf(PurlLookupPath, url.PathEscape(purl))
	if err := a.client.Get(path, &response); err != nil {
		return def.Purl{}, err
	}

	return response, nil
}

// PurlLookupBatch performs PurlLookup requests for a batch of purls.
// This endpoint only supports version lookups. Purls in requests must include a version field.
// Supported purl forms are pkg:type/namespace/name@version for a namespaced package name,
// or pkg:type/name@version for a non-namespaced package name.
// Extra fields in the purl must be empty, otherwise the request will fail.
// In particular, there must be no subpath or qualifiers. Large result sets may be paginated.
func (a *APIv3Alpha) PurlLookupBatch(req def.PurlBatchBody) (*Iterator[def.Purl], error) {
	response := &def.PurlBatchResponse{
		NextPageToken: "first",
	}

	ctx, cancel := context.WithCancel(context.Background())
	cIn := getBatch(ctx, a.client, PurlLookupBatchPath, &req, response)

	iter := Iterator[def.Purl]{
		cIn:     cIn,
		hasNext: true,
		cancel:  cancel,
	}

	return &iter, nil
}

// QueryContainerImages searches for container image repositories on DockerHub that match the requested OCI Chain ID.
// At most 1000 image repositories are returned.
// An image repository is identifier (eg. ‘tensorflow’) that refers to a collection of images.
// An OCI Chain ID is a hashed encoding of an ordered sequence of OCI layers.
// For further details see the OCI Chain ID spec. If an image contains empty layers,
// it is available from this endpoint under two different chain IDs
// - one computed by including empty layers in the calculation and one computed by excluding them.
func (a *APIv3Alpha) QueryContainerImages(chainId string) (def.ContainerImages, error) {
	var response def.ContainerImages

	var path = fmt.Sprintf(QueryContainerImagesPath, chainId)
	if err := a.client.Get(path, &response); err != nil {
		return def.ContainerImages{}, err
	}

	return response, nil
}
