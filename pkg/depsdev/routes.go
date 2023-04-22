package depsdev

const (
	// API base path
	BasePath = `https://api.deps.dev/v3alpha`

	// API routes
	GetPackagePath      = `/systems/%s/packages/%s`
	GetVersionPath      = `/systems/%s/packages/%s/versions/%s`
	GetDependenciesPath = `/systems/%s/packages/%s/versions/%s:dependencies`
	GetProjectPath      = `/projects/%s`
	GetAdvisoryPath     = `/advisories/%s`
	GetQueryPath        = `/query`
)
