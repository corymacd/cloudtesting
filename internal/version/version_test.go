package version

import (
	"testing"
)

func TestGetInfo(t *testing.T) {
	// Save original values
	origVersion := Version
	origGitCommit := GitCommit
	origBuildTime := BuildTime
	origBuildUser := BuildUser
	origGoVersion := GoVersion

	// Restore after test
	defer func() {
		Version = origVersion
		GitCommit = origGitCommit
		BuildTime = origBuildTime
		BuildUser = origBuildUser
		GoVersion = origGoVersion
	}()

	// Set test values
	Version = "test-version"
	GitCommit = "test-commit"
	BuildTime = "test-time"
	BuildUser = "test-user"
	GoVersion = "test-go"

	info := GetInfo()

	if info.Version != Version {
		t.Errorf("Version = %q, want %q", info.Version, Version)
	}
	if info.GitCommit != GitCommit {
		t.Errorf("GitCommit = %q, want %q", info.GitCommit, GitCommit)
	}
	if info.BuildTime != BuildTime {
		t.Errorf("BuildTime = %q, want %q", info.BuildTime, BuildTime)
	}
	if info.BuildUser != BuildUser {
		t.Errorf("BuildUser = %q, want %q", info.BuildUser, BuildUser)
	}
	if info.GoVersion != GoVersion {
		t.Errorf("GoVersion = %q, want %q", info.GoVersion, GoVersion)
	}
}
