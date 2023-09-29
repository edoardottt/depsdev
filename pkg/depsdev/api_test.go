package depsdev

import (
	"testing"

	"github.com/edoardottt/depsdev/pkg/client"
	"github.com/stretchr/testify/require"
)

var (
	p   Package
	v   Version
	api = NewAPI()
)

func BenchmarkGetInfo(b *testing.B) {

	b.Run("function", func(b *testing.B) {
		var (
			info Package
			err  error
		)
		for i := 0; i < b.N; i++ {
			client := client.New(BasePath)
			info, err = getPackage(client, "npm", "react")
			require.NoError(b, err)
		}
		p = info
	})

	b.Run("method", func(b *testing.B) {
		var (
			info Package
			err  error
		)
		for i := 0; i < b.N; i++ {
			info, err = api.GetInfo("npm", "react")
			require.NoError(b, err)
		}
		p = info
	})
}

func BenchmarkGetVersion(b *testing.B) {
	b.Run("function", func(b *testing.B) {
		var (
			version Version
			err     error
		)
		for i := 0; i < b.N; i++ {
			client := client.New(BasePath)
			version, err = getVersion(client, "npm", "react", "18.3.0-next-fecc288b7-20221025")
			require.NoError(b, err)
		}
		v = version
	})

	b.Run("method", func(b *testing.B) {
		var (
			version Version
			err     error
		)
		for i := 0; i < b.N; i++ {
			version, err = api.GetVersion("npm", "react", "18.3.0-next-fecc288b7-20221025")
			require.NoError(b, err)
		}
		v = version
	})
}
