package version

// Info represents version information
type Info struct {
	Version   string `json:"version" xml:"version"`
	GitCommit string `json:"gitCommit" xml:"gitCommit"`
	BuildTime string `json:"buildTime" xml:"buildTime"`
	BuildUser string `json:"buildUser" xml:"buildUser"`
	GoVersion string `json:"goVersion" xml:"goVersion"`
}

var (
	Version   = "unknown"
	GitCommit = "unknown"
	BuildTime = "unknown"
	BuildUser = "unknown"
	GoVersion = "unknown"
)

// GetInfo returns structured version information
func GetInfo() Info {
	return Info{
		Version:   Version,
		GitCommit: GitCommit,
		BuildTime: BuildTime,
		BuildUser: BuildUser,
		GoVersion: GoVersion,
	}
}
