/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

// Findings represents the root JSON object.
type Findings struct {
	PackageKey          *PackageKey      `json:"packageKey,omitempty"`
	VersionKey          *VersionKey      `json:"versionKey,omitempty"`
	RecommendedVersions []VersionDetails `json:"recommendedVersions,omitempty"`
	RequestedVersion    *VersionDetails  `json:"requestedVersion,omitempty"`
	DefaultVersion      *VersionDetails  `json:"defaultVersion,omitempty"`
	PackageFindings     []Finding        `json:"packageFindings,omitempty"`
}

// VersionDetails contains findings and metadata for a specific version.
type VersionDetails struct {
	VersionKey  *VersionKey `json:"versionKey,omitempty"`
	IsDefault   bool        `json:"isDefault,omitempty"`
	Findings    []Finding   `json:"findings,omitempty"`
	CooldownEnd string      `json:"cooldownEnd,omitempty"`
}

// Finding indicates specific risks or statuses associated with a package or version.
type Finding struct {
	Type              string             `json:"type,omitempty"`
	Risk              string             `json:"risk,omitempty"`
	DeprecatedContext *DeprecatedContext `json:"deprecatedContext,omitempty"`
	CooldownContext   *CooldownContext   `json:"cooldownContext,omitempty"`
	LowUsageContext   *LowUsageContext   `json:"lowUsageContext,omitempty"`
}

// DeprecatedContext provides the reason for deprecation.
type DeprecatedContext struct {
	Reason string `json:"reason,omitempty"`
}

// CooldownContext indicates the cooldown period for a specific version.
type CooldownContext struct {
	End string `json:"end,omitempty"`
}

// LowUsageContext provides alternative packages with higher usage.
type LowUsageContext struct {
	AlternativePackages []string `json:"alternativePackages,omitempty"`
}
