/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

import (
	"testing"

	"github.com/edoardottt/depsdev/pkg/client"
	def "github.com/edoardottt/depsdev/pkg/depsdev/definitions"
	"github.com/stretchr/testify/require"
)

var (
	p   def.Package
	v   def.Version
	api = NewV3API()
)

func BenchmarkGetInfo(b *testing.B) {
	b.Run("function", func(b *testing.B) {
		var (
			info def.Package
			err  error
		)

		for i := 0; i < b.N; i++ {
			client := client.New(V3BasePath)
			info, err = getPackage(client, "npm", "react")
			require.NoError(b, err)
		}

		p = info
	})

	b.Run("method", func(b *testing.B) {
		var (
			info def.Package
			err  error
		)

		for i := 0; i < b.N; i++ {
			info, err = api.GetPackage("npm", "react")
			require.NoError(b, err)
		}

		p = info
	})
}

func BenchmarkGetVersion(b *testing.B) {
	b.Run("function", func(b *testing.B) {
		var (
			version def.Version
			err     error
		)

		for i := 0; i < b.N; i++ {
			client := client.New(V3BasePath)
			version, err = getVersion(client, "npm", "react", "18.3.0-next-fecc288b7-20221025")
			require.NoError(b, err)
		}

		v = version
	})

	b.Run("method", func(b *testing.B) {
		var (
			version def.Version
			err     error
		)

		for i := 0; i < b.N; i++ {
			version, err = api.GetVersion("npm", "react", "18.3.0-next-fecc288b7-20221025")
			require.NoError(b, err)
		}

		v = version
	})
}
