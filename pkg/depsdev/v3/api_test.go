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
	"log"
	"testing"

	def "github.com/edoardottt/depsdev/pkg/depsdev/definitions"
	"github.com/edoardottt/depsdev/pkg/depsdev/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	api = depsdev.NewV3API()
)

func TestGetPackage(t *testing.T) {
	result := `{
		"packageKey": {
		  "system": "NPM",
		  "name": "defangjs"
		},
		"versions": [
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "defangjs",
			  "version": "1.0.0"
			},
			"publishedAt": "2022-03-10T19:53:06Z"
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "defangjs",
			  "version": "1.0.1"
			},
			"publishedAt": "2022-03-11T13:44:30Z"
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "defangjs",
			  "version": "1.0.2"
			},
			"publishedAt": "2022-03-11T13:52:09Z"
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "defangjs",
			  "version": "1.0.3"
			},
			"publishedAt": "2022-03-12T15:41:16Z"
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "defangjs",
			  "version": "1.0.4"
			},
			"publishedAt": "2022-03-17T18:13:31Z"
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "defangjs",
			  "version": "1.0.5"
			},
			"publishedAt": "2022-05-07T13:17:15Z"
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "defangjs",
			  "version": "1.0.6"
			},
			"publishedAt": "2022-09-17T12:46:20Z"
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "defangjs",
			  "version": "1.0.7"
			},
			"isDefault": true,
			"publishedAt": "2023-05-16T09:48:31Z"
		  }
		]
	  }`

	t.Run("GetPackage npm defangjs", func(t *testing.T) {
		got, err := api.GetPackage("npm", "defangjs")
		require.Nil(t, err)

		var d def.Package

		if err := json.Unmarshal([]byte(result), &d); err != nil {
			log.Fatal(err)
		}

		require.Equal(t, d, got)
	})
}

func TestGetVersion(t *testing.T) {
	result := `{
	"versionKey":{
		"system":"NPM",
		"name":"defangjs",
		"version":"1.0.7"
	},
	"publishedAt":"2023-05-16T09:48:31Z",
	"isDefault":true,
	"licenses":[
		"GPL-3.0"
	],
	"advisoryKeys":[],
	"links":[
		{
			"label":"HOMEPAGE",
			"url":"https://github.com/edoardottt/defangjs#readme"
		},
		{
			"label":"ISSUE_TRACKER",
			"url":"https://github.com/edoardottt/defangjs/issues"
		},
		{
			"label":"ORIGIN",
			"url":"https://registry.npmjs.org/defangjs/1.0.7"
		},
		{
			"label":"SOURCE_REPO",
			"url":"git+https://github.com/edoardottt/defangjs.git"
		}
	],
	"slsaProvenances":[],
	"attestations":[],
	"registries":[
		"https://registry.npmjs.org/"
	],
	"relatedProjects":[
		{
			"projectKey":{
				"id":"github.com/edoardottt/defangjs"
			},
			"relationProvenance":"UNVERIFIED_METADATA",
			"relationType":"ISSUE_TRACKER"
		},
		{
			"projectKey":{
				"id":"github.com/edoardottt/defangjs"
			},
			"relationProvenance":"UNVERIFIED_METADATA",
			"relationType":"SOURCE_REPO"
		}
	]
}`

	t.Run("GetVersion npm defangjs 1.0.7", func(t *testing.T) {
		got, err := api.GetVersion("npm", "defangjs", "1.0.7")
		require.Nil(t, err)

		var v def.Version

		if err := json.Unmarshal([]byte(result), &v); err != nil {
			log.Fatal(err)
		}

		require.Equal(t, v, got)
	})
}

func TestGetRequirements(t *testing.T) {
	result := `{
		"npm": {
		  "dependencies": {
			"dependencies": [
			  {
				"name": "fast-redact",
				"requirement": "^3.0.0"
			  },
			  {
				"name": "fast-safe-stringify",
				"requirement": "^2.0.8"
			  },
			  {
				"name": "flatstr",
				"requirement": "^1.0.12"
			  },
			  {
				"name": "pino-std-serializers",
				"requirement": "^3.1.0"
			  },
			  {
				"name": "process-warning",
				"requirement": "^1.0.0"
			  },
			  {
				"name": "quick-format-unescaped",
				"requirement": "^4.0.3"
			  },
			  {
				"name": "sonic-boom",
				"requirement": "^1.0.2"
			  }
			],
			"devDependencies": [
			  {
				"name": "airtap",
				"requirement": "4.0.3"
			  },
			  {
				"name": "benchmark",
				"requirement": "^2.1.4"
			  },
			  {
				"name": "bole",
				"requirement": "^4.0.0"
			  },
			  {
				"name": "bunyan",
				"requirement": "^1.8.14"
			  },
			  {
				"name": "docsify-cli",
				"requirement": "^4.4.1"
			  },
			  {
				"name": "eslint",
				"requirement": "^7.17.0"
			  },
			  {
				"name": "eslint-config-standard",
				"requirement": "^16.0.2"
			  },
			  {
				"name": "eslint-plugin-import",
				"requirement": "^2.22.1"
			  },
			  {
				"name": "eslint-plugin-node",
				"requirement": "^11.1.0"
			  },
			  {
				"name": "eslint-plugin-promise",
				"requirement": "^5.1.0"
			  },
			  {
				"name": "execa",
				"requirement": "^5.0.0"
			  },
			  {
				"name": "fastbench",
				"requirement": "^1.0.1"
			  },
			  {
				"name": "flush-write-stream",
				"requirement": "^2.0.0"
			  },
			  {
				"name": "import-fresh",
				"requirement": "^3.2.1"
			  },
			  {
				"name": "log",
				"requirement": "^6.0.0"
			  },
			  {
				"name": "loglevel",
				"requirement": "^1.6.7"
			  },
			  {
				"name": "pino-pretty",
				"requirement": "^5.0.0"
			  },
			  {
				"name": "pre-commit",
				"requirement": "^1.2.2"
			  },
			  {
				"name": "proxyquire",
				"requirement": "^2.1.3"
			  },
			  {
				"name": "pump",
				"requirement": "^3.0.0"
			  },
			  {
				"name": "semver",
				"requirement": "^7.0.0"
			  },
			  {
				"name": "split2",
				"requirement": "^3.1.1"
			  },
			  {
				"name": "steed",
				"requirement": "^1.1.3"
			  },
			  {
				"name": "strip-ansi",
				"requirement": "^6.0.0"
			  },
			  {
				"name": "tap",
				"requirement": "^15.0.1"
			  },
			  {
				"name": "tape",
				"requirement": "^5.0.0"
			  },
			  {
				"name": "through2",
				"requirement": "^4.0.0"
			  },
			  {
				"name": "winston",
				"requirement": "^3.3.3"
			  }
			],
			"optionalDependencies": [],
			"peerDependencies": [],
			"bundleDependencies": []
		  },
		  "Bundled": []
		}
	  }`

	t.Run("GetRequirements npm pino 6.14.0", func(t *testing.T) {
		got, err := api.GetRequirements("npm", "pino", "6.14.0")
		require.Nil(t, err)

		var r def.Requirements

		if err := json.Unmarshal([]byte(result), &r); err != nil {
			log.Fatal(err)
		}

		require.Equal(t, r, got)
	})
}

func TestGetDependencies(t *testing.T) {
	result := `{
		"nodes": [
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "pino",
			  "version": "6.14.0"
			},
			"relation": "SELF",
			"errors": []
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "atomic-sleep",
			  "version": "1.0.0"
			},
			"relation": "INDIRECT",
			"errors": []
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "fast-redact",
			  "version": "3.5.0"
			},
			"relation": "DIRECT",
			"errors": []
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "fast-safe-stringify",
			  "version": "2.1.1"
			},
			"relation": "DIRECT",
			"errors": []
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "flatstr",
			  "version": "1.0.12"
			},
			"relation": "DIRECT",
			"errors": []
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "pino-std-serializers",
			  "version": "3.2.0"
			},
			"relation": "DIRECT",
			"errors": []
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "process-warning",
			  "version": "1.0.0"
			},
			"relation": "DIRECT",
			"errors": []
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "quick-format-unescaped",
			  "version": "4.0.4"
			},
			"relation": "DIRECT",
			"errors": []
		  },
		  {
			"versionKey": {
			  "system": "NPM",
			  "name": "sonic-boom",
			  "version": "1.4.1"
			},
			"relation": "DIRECT",
			"errors": []
		  }
		],
		"edges": [
		  {
			"toNode": 2,
			"requirement": "^3.0.0"
		  },
		  {
			"toNode": 3,
			"requirement": "^2.0.8"
		  },
		  {
			"toNode": 4,
			"requirement": "^1.0.12"
		  },
		  {
			"toNode": 5,
			"requirement": "^3.1.0"
		  },
		  {
			"toNode": 6,
			"requirement": "^1.0.0"
		  },
		  {
			"toNode": 7,
			"requirement": "^4.0.3"
		  },
		  {
			"toNode": 8,
			"requirement": "^1.0.2"
		  },
		  {
			"fromNode": 8,
			"toNode": 1,
			"requirement": "^1.0.0"
		  },
		  {
			"fromNode": 8,
			"toNode": 4,
			"requirement": "^1.0.12"
		  }
		]
	  }`

	t.Run("GetDependencies npm pino 6.14.0", func(t *testing.T) {
		got, err := api.GetDependencies("npm", "pino", "6.14.0")
		require.Nil(t, err)

		var d def.Dependencies

		if err := json.Unmarshal([]byte(result), &d); err != nil {
			log.Fatal(err)
		}

		require.Equal(t, d, got)
	})
}

func TestGetProject(t *testing.T) {
	t.Run("GetProject npm defangjs", func(t *testing.T) {
		got, err := api.GetProject("github.com/edoardottt/defangjs")
		require.Nil(t, err)

		// no checking of the actual value because they can change over time
		// we just ensure the call to the API and unmarshaling of the response works properly
		assert.NotEmpty(t, got)
	})
}

func TestGetProjectPackageVersions(t *testing.T) {
	t.Run("GetProjectPackageVersions npm defangjs", func(t *testing.T) {
		got, err := api.GetProjectPackageVersions("github.com/edoardottt/defangjs")
		require.Nil(t, err)

		// no checking of the actual value because they can change over time
		// we just ensure the call to the API and unmarshaling of the response works properly
		assert.NotEmpty(t, got)
	})
}

func TestGetAdvisory(t *testing.T) {
	result := `{
		"advisoryKey": {
		  "id": "GHSA-jfh8-c2jp-5v3q"
		},
		"url": "https://osv.dev/vulnerability/GHSA-jfh8-c2jp-5v3q",
		"title": "Remote code injection in Log4j",
		"aliases": [
		  "CVE-2021-44228"
		],
		"cvss3Score": 10,
		"cvss3Vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:H/I:H/A:H/E:H"
	  }`

	t.Run("GetAdvisory GHSA-jfh8-c2jp-5v3q", func(t *testing.T) {
		got, err := api.GetAdvisory("GHSA-jfh8-c2jp-5v3q")
		require.Nil(t, err)

		var a def.Advisory

		if err := json.Unmarshal([]byte(result), &a); err != nil {
			log.Fatal(err)
		}

		require.Equal(t, a, got)
	})
}

func TestQuery(t *testing.T) {
	result := `{
		"results": [
			{
			"version": {
				"versionKey": {
				"system": "NPM",
				"name": "defangjs",
				"version": "1.0.7"
				},
				"publishedAt": "2023-05-16T09:48:31Z",
				"isDefault": true,
				"licenses": [
				"GPL-3.0"
				],
				"advisoryKeys": [],
				"links": [
				{
					"label": "HOMEPAGE",
					"url": "https://github.com/edoardottt/defangjs#readme"
				},
				{
					"label": "ISSUE_TRACKER",
					"url": "https://github.com/edoardottt/defangjs/issues"
				},
				{
					"label": "ORIGIN",
					"url": "https://registry.npmjs.org/defangjs/1.0.7"
				},
				{
					"label": "SOURCE_REPO",
					"url": "git+https://github.com/edoardottt/defangjs.git"
				}
				],
				"slsaProvenances": [],
				"attestations": [],
				"registries": [
				"https://registry.npmjs.org/"
				],
				"relatedProjects": [
				{
					"projectKey": {
					"id": "github.com/edoardottt/defangjs"
					},
					"relationProvenance": "UNVERIFIED_METADATA",
					"relationType": "ISSUE_TRACKER"
				},
				{
					"projectKey": {
					"id": "github.com/edoardottt/defangjs"
					},
					"relationProvenance": "UNVERIFIED_METADATA",
					"relationType": "SOURCE_REPO"
				}
				]
			}
			}
		]
		}`

	t.Run("Query versionKey.system=NPM&versionKey.name=defangjs&versionKey.version=1.0.7", func(t *testing.T) {
		got, err := api.Query("versionKey.system=NPM&versionKey.name=defangjs&versionKey.version=1.0.7")
		require.Nil(t, err)

		var a def.Results

		if err := json.Unmarshal([]byte(result), &a); err != nil {
			log.Fatal(err)
		}

		require.Equal(t, a, got)
	})
}
