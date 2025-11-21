<h1 align="center">
  depsdev
  <br>
</h1>
<h4 align="center">CLI client (and Golang module) for deps.dev API.<br>Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.</h4>

<h6 align="center"> Coded with 💙 by edoardottt </h6>

<p align="center">

  <a href="https://goreportcard.com/report/github.com/edoardottt/depsdev">
      <img src="https://goreportcard.com/badge/github.com/edoardottt/depsdev" alt="go report card">
  </a>

  <a href="https://github.com/edoardottt/depsdev/actions">
      <img src="https://github.com/edoardottt/depsdev/actions/workflows/go.yml/badge.svg" alt="go action">
  </a>

<br>
  <!--Tweet button-->
  <a href="https://twitter.com/intent/tweet?text=depsdev%20-%20CLI%20client%20for%20deps.dev%20API.%20Free%20access%20to%20dependencies%2C%20licenses%2C%20advisories%2C%20and%20other%20critical%20health%20and%20security%20signals%20for%20open%20source%20package%20versions.%20https%3A%2F%2Fgithub.com%2Fedoardottt%2Fdepsdev%20%23golang%20%23github%20%23linux%20%23infosec%20%23bugbounty%20%23security" target="_blank">Share on Twitter!
  </a>
</p>

<p align="center">
  <a href="#install-">Install</a> •
  <a href="#get-started-">Get Started</a> •
  <a href="#examples-bulb">Examples</a> •
  <a href="#changelog-">Changelog</a> •
  <a href="#contributing-">Contributing</a> •
  <a href="#license-">License</a>
</p>

<p align="center">
  <img src="https://github.com/edoardottt/images/blob/main/depsdev/depsdev.gif">
</p>
  
Install 📡
----------

### Using Snap

```console
sudo snap install depsdev
```

### Using Go

```console
go install github.com/edoardottt/depsdev@latest
```

Get Started 🎉
----------

```console
Usage:
  depsdev [command]

Available Commands:
  advisory    Get info about an (OSV) advisory
  completion  Generate the autocompletion script for the specified shell
  deps        Get info about a package's dependencies
  graph       Generate a Graphviz compatible dependencies graph
  help        Help about any command
  info        Get info about a package or a specific version of that
  packages    Get info about a project's package versions (GitHub, GitLab, or BitBucket)
  project     Get info about a project (GitHub, GitLab, or BitBucket)
  query       Get info about multiple package versions using a query
  reqs        Get info about a package's requirements

Flags:
  -h, --help   help for depsdev

Use "depsdev [command] --help" for more information about a command.
```

Examples 💡
----------

> **Note**
> The supported package managers are `go`, `npm`, `cargo`, `rubygems`, `maven`, `pypi` and `nuget`.  

For more information [read the API documentation](https://docs.deps.dev/api/v3).

### CLI

<br>

Get information about a package, including a list of its available versions, with the default version marked if known.

```console
depsdev info npm @colors/colors
```

<br>

Get information about a specific package version including its licenses and any security advisories known to affect it.

```console
depsdev info npm @colors/colors 1.5.0
```

<br>

Get information about a resolved dependency graph for the given package version.

```console
depsdev deps npm @colors/colors 1.5.0
```

<br>

Get information about projects hosted by GitHub, GitLab, or BitBucket (if available).

```console
depsdev project github.com/facebook/react
```

<br>

Get information about security advisories hosted by OSV.

```console
depsdev advisory GHSA-2qrg-x229-3v8q
```

<br>

Get information about multiple package versions, which can be specified by name, content hash, or both.

```console
depsdev query "versionKey.system=NPM&versionKey.name=react&versionKey.version=18.2.0"
```

<br>

Generate a Graphviz compatible dependencies graph for a specific version of a package.

```console
depsdev graph npm slice-ansi 6.0.0
```

<br>

Get information about the package requirements for a given version in a system-specific format.

```console
depsdev reqs npm slice-ansi 6.0.0
```

<br>

Returns known mappings between the requested project and package versions.

```console
depsdev packages github.com/eslint/espree
```

### Use depsdev as a Go module

You can use *v3* or *v3alpha*.

#### v3

Core features with a stability guarantee and deprecation policy. Recommended for most users.

```Go
package main

import (
    "fmt"
    depsdev "github.com/edoardottt/depsdev/pkg/depsdev/v3"
)

func main() {
    client := depsdev.NewV3API()
    i, err := client.GetInfo("npm", "defangjs")
    if err != nil {
      fmt.Println(err)
    }
    
    fmt.Println(i)
}
```

#### v3alpha

All the features of v3, with additional experimental features. May change in incompatible ways from time to time.

```Go
package main

import (
    "fmt"
    depsdev "github.com/edoardottt/depsdev/pkg/depsdev/v3alpha"
)

func main() {
    client := depsdev.NewV3AlphaAPI()
    i, err := client.GetInfo("npm", "defangjs")
    if err != nil {
      fmt.Println(err)
    }
    
    fmt.Println(i)
}
```

Read the full [`package documentation here`](https://pkg.go.dev/github.com/edoardottt/depsdev/pkg/depsdev)

Changelog 📌
-------

Detailed changes for each release are documented in the [release notes](https://github.com/edoardottt/depsdev/releases).

Contributing 🛠
-------

Just open an [issue](https://github.com/edoardottt/depsdev/issues) / [pull request](https://github.com/edoardottt/depsdev/pulls).

Before opening a pull request, download [golangci-lint](https://golangci-lint.run/usage/install/) and run

```console
golangci-lint run
```

If there aren't errors, go ahead :)

The HTTP client implementation is partially taken from [@liamg/hackerone](https://github.com/liamg/hackerone).

License 📝
-------

This repository is under [Apache2.0 License](https://github.com/edoardottt/depsdev/blob/main/LICENSE).  
[edoardottt.com](https://edoardottt.com) to contact me.
