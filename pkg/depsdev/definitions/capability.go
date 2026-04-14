/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

// Capabilities represents the collection of Capslock capabilities for a package version.
type Capabilities struct {
	Capabilities []CapabilityDetail `json:"capabilities,omitempty"`
}

// CapabilityDetail describes a specific capability and the frequency of its usage.
type CapabilityDetail struct {
	// Capability is the name of the Capslock capability (e.g., "CAPABILITY_NETWORK").
	Capability string `json:"capability,omitempty"`
	// DirectCount is the number of direct calls from the package to this capability.
	DirectCount int `json:"directCount,omitempty"`
	// IndirectCount is the number of calls to this capability via other dependencies.
	IndirectCount int `json:"indirectCount,omitempty"`
}
