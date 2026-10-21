/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev_test

import (
	"encoding/json"
	"testing"

	def "github.com/edoardottt/depsdev/pkg/depsdev/definitions"
	"github.com/stretchr/testify/require"
)

func TestFindingBatchRequestSetsOneKey(t *testing.T) {
	tests := []struct {
		name     string
		request  def.FindingBatchRequest
		expected string
	}{
		{
			name:     "package key",
			request:  def.FindingBatchRequest{PackageKey: def.PackageKey{System: "NPM", Name: "react"}},
			expected: `{"packageKey":{"system":"NPM","name":"react"}}`,
		},
		{
			name:     "version key",
			request:  def.FindingBatchRequest{VersionKey: def.VersionKey{System: "NPM", Name: "react", Version: "18.2.0"}},
			expected: `{"versionKey":{"system":"NPM","name":"react","version":"18.2.0"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.request)
			require.NoError(t, err)
			require.JSONEq(t, tt.expected, string(got))
		})
	}
}

func TestResultsArtifacts(t *testing.T) {
	var r def.Results

	err := json.Unmarshal([]byte(`{"results":[{"version":{},"artifacts":[{"url":"https://example.com/a.jar"}]}]}`), &r)
	require.NoError(t, err)
	require.Equal(t, []def.Artifact{{URL: "https://example.com/a.jar"}}, r.Results[0].Artifacts)
}

func TestPackageVersionsRelation(t *testing.T) {
	var p def.PackageVersions

	err := json.Unmarshal([]byte(`{"versions":[{
		"versionKey":{"system":"GO","name":"github.com/edoardottt/depsdev","version":"v0.2.3"},
		"relationType":"SOURCE_REPO",
		"relationProvenance":"GO_ORIGIN",
		"slsaProvenances":[],
		"attestations":[]
	}]}`), &p)
	require.NoError(t, err)

	require.Equal(t, []def.PackageVersion{{
		VersionKey:         def.VersionKey{System: "GO", Name: "github.com/edoardottt/depsdev", Version: "v0.2.3"},
		RelationType:       "SOURCE_REPO",
		RelationProvenance: "GO_ORIGIN",
		SlsaProvenance:     []def.SlsaProvenance{},
		Attestations:       []def.Attestation{},
	}}, p.Versions)
}

func TestPypiExternalDependencies(t *testing.T) {
	var r def.Requirements

	err := json.Unmarshal([]byte(`{"pypi":{"externalDependencies":[
		{"name":"libpq","versionSpecifier":">=9","environmentMarker":"os_name == \"posix\""}
	]}}`), &r)
	require.NoError(t, err)

	require.Equal(t, []def.PypiExternalDependency{{
		Name:              "libpq",
		VersionSpecifier:  ">=9",
		EnvironmentMarker: `os_name == "posix"`,
	}}, r.Pypi.ExternalDependencies)
}

func TestProjectBatchResponseRequest(t *testing.T) {
	var r def.ProjectBatchResponse

	err := json.Unmarshal([]byte(`{"responses":[{
		"request":{"projectKey":{"id":"github.com/facebook/react"}},
		"project":{"projectKey":{"id":"github.com/facebook/react"}}
	}],"nextPageToken":""}`), &r)
	require.NoError(t, err)
	require.Equal(t, "github.com/facebook/react", r.Responses[0].Request.ProjectKey.ID)
}

func TestVersionIsDeprecatedKey(t *testing.T) {
	got, err := json.Marshal(def.Version{IsDeprecated: true})
	require.NoError(t, err)

	var m map[string]any

	require.NoError(t, json.Unmarshal(got, &m))
	require.Contains(t, m, "isDeprecated")
	require.NotContains(t, m, "IsDeprecated")
}
