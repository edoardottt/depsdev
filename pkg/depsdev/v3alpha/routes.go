/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

const (
	// API base path.
	V3AlphaBasePath = `https://api.deps.dev/v3alpha`

	// API routes.
	GetPackagePath                = `/systems/%s/packages/%s`
	GetSimilarlyNamedPackagesPath = `/systems/%s/packages/%s:similarlyNamedPackages`
	GetVersionPath                = `/systems/%s/packages/%s/versions/%s`
	GetRequirementsPath           = `/systems/%s/packages/%s/versions/%s:requirements`
	GetCapabilitiesPath           = `/systems/%s/packages/%s/versions/%s:capabilities`
	GetDependenciesPath           = `/systems/%s/packages/%s/versions/%s:dependencies`
	GetDependentsPath             = `/systems/%s/packages/%s/versions/%s:dependents`
	GetProjectPath                = `/projects/%s`
	GetAdvisoryPath               = `/advisories/%s`
	QueryPath                     = `/query`
	GetProjectPackageVersionsPath = `/projects/%s:packageversions`
	GetVersionBatchPath           = `/versionbatch`
	GetProjectBatchPath           = `/projectbatch`
	PurlLookupPath                = `/purl/%s`
	PurlLookupBatchPath           = `/purlbatch`
	QueryContainerImagesPath      = `/querycontainerimages/%s`
)
