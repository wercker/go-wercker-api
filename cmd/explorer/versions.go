package main

import "fmt"

var (
	// GitCommit is the git commit hash associated with this build.
	GitCommit = ""

	// MajorVersion is the semver major version.
	MajorVersion = "1"

	// MinorVersion is the semver minor version.
	MinorVersion = "0"

	// PatchVersion is the semver patch version. (use 0 for dev, build process
	// will inject a build number)
	PatchVersion = "0"
)

// Version returns a semver compatible version for this build.
func Version() string {
	return fmt.Sprintf("%s.%s.%s", MajorVersion, MinorVersion, PatchVersion)
}

// FullVersion returns the semver version and the git version if available.
func FullVersion() string {
	semver := Version()
	if GitCommit == "" {
		return semver
	}
	return fmt.Sprintf("%s (Git commit: %s)", semver, GitCommit)
}

// GetVersions returns a Versions struct filled with the current values.
func GetVersions() *Versions {
	return &Versions{
		GitCommit: GitCommit,
		Version:   Version(),
	}
}

// Versions contains GitCommit and Version as a JSON marshall friendly struct.
type Versions struct {
	GitCommit string `json:"gitCommit,omitempty"`
	Version   string `json:"version,omitempty"`
}
