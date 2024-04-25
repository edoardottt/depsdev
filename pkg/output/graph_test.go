/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package output_test

import (
	"encoding/json"
	"log"
	"testing"

	definitions "github.com/edoardottt/depsdev/pkg/depsdev/definitions"
	"github.com/edoardottt/depsdev/pkg/output"
	"github.com/stretchr/testify/require"
)

func TestGenerateGraph(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "cariddi v1.1.8",
			input: `{
"nodes": [
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/edoardottt/cariddi",
	  "version": "v1.1.8"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/andybalholm/cascadia",
	  "version": "v1.3.1"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/antchfx/htmlquery",
	  "version": "v1.2.4"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/antchfx/xmlquery",
	  "version": "v1.3.10"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/antchfx/xpath",
	  "version": "v1.2.0"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/fatih/color",
	  "version": "v1.13.0"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/gobwas/glob",
	  "version": "v0.2.3"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/gocolly/colly",
	  "version": "v1.2.0"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/golang/groupcache",
	  "version": "v0.0.0-20200121045136-8c9f03a8e57e"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/golang/protobuf",
	  "version": "v1.3.1"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/kennygrant/sanitize",
	  "version": "v1.2.4"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/mattn/go-colorable",
	  "version": "v0.1.9"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/mattn/go-isatty",
	  "version": "v0.0.14"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/PuerkitoBio/goquery",
	  "version": "v1.8.0"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/saintfish/chardet",
	  "version": "v0.0.0-20120816061221-3af4cd4741ca"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "github.com/temoto/robotstxt",
	  "version": "v1.1.2"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "golang.org/x/net",
	  "version": "v0.0.0-20210916014120-12bc252f5db8"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "golang.org/x/sys",
	  "version": "v0.0.0-20210630005230-0f9fa26af87c"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "golang.org/x/text",
	  "version": "v0.3.6"
	}
  },
  {
	"versionKey": {
	  "system": "GO",
	  "name": "google.golang.org/appengine",
	  "version": "v1.6.7"
	}
  }
],
"edges": [
  {
	"toNode": 5,
	"requirement": "v1.13.0"
  },
  {
	"toNode": 7,
	"requirement": "v1.2.0"
  },
  {
	"fromNode": 1,
	"toNode": 16,
	"requirement": "v0.0.0-20210916014120-12bc252f5db8"
  },
  {
	"fromNode": 2,
	"toNode": 4,
	"requirement": "v1.2.0"
  },
  {
	"fromNode": 2,
	"toNode": 8,
	"requirement": "v0.0.0-20200121045136-8c9f03a8e57e"
  },
  {
	"fromNode": 2,
	"toNode": 16,
	"requirement": "v0.0.0-20200421231249-e086a090c8fd"
  },
  {
	"fromNode": 3,
	"toNode": 4,
	"requirement": "v1.2.0"
  },
  {
	"fromNode": 3,
	"toNode": 8,
	"requirement": "v0.0.0-20200121045136-8c9f03a8e57e"
  },
  {
	"fromNode": 3,
	"toNode": 16,
	"requirement": "v0.0.0-20200813134508-3edf25e44fcc"
  },
  {
	"fromNode": 5,
	"toNode": 11,
	"requirement": "v0.1.9"
  },
  {
	"fromNode": 5,
	"toNode": 12,
	"requirement": "v0.0.14"
  },
  {
	"fromNode": 7,
	"toNode": 2
  },
  {
	"fromNode": 7,
	"toNode": 3
  },
  {
	"fromNode": 7,
	"toNode": 6
  },
  {
	"fromNode": 7,
	"toNode": 10
  },
  {
	"fromNode": 7,
	"toNode": 13
  },
  {
	"fromNode": 7,
	"toNode": 14
  },
  {
	"fromNode": 7,
	"toNode": 15
  },
  {
	"fromNode": 7,
	"toNode": 16
  },
  {
	"fromNode": 7,
	"toNode": 19
  },
  {
	"fromNode": 10,
	"toNode": 16
  },
  {
	"fromNode": 11,
	"toNode": 12,
	"requirement": "v0.0.12"
  },
  {
	"fromNode": 12,
	"toNode": 17,
	"requirement": "v0.0.0-20210630005230-0f9fa26af87c"
  },
  {
	"fromNode": 13,
	"toNode": 1,
	"requirement": "v1.3.1"
  },
  {
	"fromNode": 13,
	"toNode": 16,
	"requirement": "v0.0.0-20210916014120-12bc252f5db8"
  },
  {
	"fromNode": 16,
	"toNode": 18,
	"requirement": "v0.3.6"
  },
  {
	"fromNode": 19,
	"toNode": 9,
	"requirement": "v1.3.1"
  },
  {
	"fromNode": 19,
	"toNode": 16,
	"requirement": "v0.0.0-20190603091049-60506f45cf65"
  }
]
			  }`,
			want: `digraph {
  0 [label="github.com/edoardottt/cariddi@v1.1.8"];
  1 [label="github.com/andybalholm/cascadia@v1.3.1"];
  2 [label="github.com/antchfx/htmlquery@v1.2.4"];
  3 [label="github.com/antchfx/xmlquery@v1.3.10"];
  4 [label="github.com/antchfx/xpath@v1.2.0"];
  5 [label="github.com/fatih/color@v1.13.0"];
  6 [label="github.com/gobwas/glob@v0.2.3"];
  7 [label="github.com/gocolly/colly@v1.2.0"];
  8 [label="github.com/golang/groupcache@v0.0.0-20200121045136-8c9f03a8e57e"];
  9 [label="github.com/golang/protobuf@v1.3.1"];
  10 [label="github.com/kennygrant/sanitize@v1.2.4"];
  11 [label="github.com/mattn/go-colorable@v0.1.9"];
  12 [label="github.com/mattn/go-isatty@v0.0.14"];
  13 [label="github.com/PuerkitoBio/goquery@v1.8.0"];
  14 [label="github.com/saintfish/chardet@v0.0.0-20120816061221-3af4cd4741ca"];
  15 [label="github.com/temoto/robotstxt@v1.1.2"];
  16 [label="golang.org/x/net@v0.0.0-20210916014120-12bc252f5db8"];
  17 [label="golang.org/x/sys@v0.0.0-20210630005230-0f9fa26af87c"];
  18 [label="golang.org/x/text@v0.3.6"];
  19 [label="google.golang.org/appengine@v1.6.7"];
  0 -> 5 [label="v1.13.0"];
  0 -> 7 [label="v1.2.0"];
  1 -> 16 [label="v0.0.0-20210916014120-12bc252f5db8"];
  2 -> 4 [label="v1.2.0"];
  2 -> 8 [label="v0.0.0-20200121045136-8c9f03a8e57e"];
  2 -> 16 [label="v0.0.0-20200421231249-e086a090c8fd"];
  3 -> 4 [label="v1.2.0"];
  3 -> 8 [label="v0.0.0-20200121045136-8c9f03a8e57e"];
  3 -> 16 [label="v0.0.0-20200813134508-3edf25e44fcc"];
  5 -> 11 [label="v0.1.9"];
  5 -> 12 [label="v0.0.14"];
  7 -> 2 [label=""];
  7 -> 3 [label=""];
  7 -> 6 [label=""];
  7 -> 10 [label=""];
  7 -> 13 [label=""];
  7 -> 14 [label=""];
  7 -> 15 [label=""];
  7 -> 16 [label=""];
  7 -> 19 [label=""];
  10 -> 16 [label=""];
  11 -> 12 [label="v0.0.12"];
  12 -> 17 [label="v0.0.0-20210630005230-0f9fa26af87c"];
  13 -> 1 [label="v1.3.1"];
  13 -> 16 [label="v0.0.0-20210916014120-12bc252f5db8"];
  16 -> 18 [label="v0.3.6"];
  19 -> 9 [label="v1.3.1"];
  19 -> 16 [label="v0.0.0-20190603091049-60506f45cf65"];
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := definitions.Dependencies{}

			err := json.Unmarshal([]byte(tt.input), &d)
			if err != nil {
				log.Fatal("error while unmarshaling test input")
			}

			got, err := output.GenerateGraph(d)
			require.Nil(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
