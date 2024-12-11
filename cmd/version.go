/*
Copyright © 2022 Cory M. MacDonald <corymacd@netsrv.co>
This file is part of the intel gathering honeypot application for myipintel.com
*/
package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// VersionInfo represents version information
type VersionInfo struct {
	Version   string    `json:"version" xml:"version"`
	BuildTime time.Time `json:"build_time,omitempty" xml:"build_time,omitempty"`
	GitCommit string    `json:"git_commit,omitempty" xml:"git_commit,omitempty"`
	GoVersion string    `json:"go_version,omitempty" xml:"go_version,omitempty"`
}

// These variables are populated via ldflags during build
var (
	version   = "dev"
	buildTime = ""
	gitCommit = ""
	goVersion = ""
)

// GetVersionInfo returns the structured version information
func GetVersionInfo() VersionInfo {
	bt, _ := time.Parse(time.RFC3339, buildTime)
	return VersionInfo{
		Version:   version,
		BuildTime: bt,
		GitCommit: gitCommit,
		GoVersion: goVersion,
	}
}

// String implements fmt.Stringer for pretty console output
func (v VersionInfo) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Version:    %s\n", v.Version))
	if !v.BuildTime.IsZero() {
		b.WriteString(fmt.Sprintf("Built:      %s\n", v.BuildTime.Format(time.RFC3339)))
	}
	if v.GitCommit != "" {
		b.WriteString(fmt.Sprintf("Git commit: %s\n", v.GitCommit))
	}
	if v.GoVersion != "" {
		b.WriteString(fmt.Sprintf("Go version: %s\n", v.GoVersion))
	}
	return b.String()
}

// VersionHandler handles HTTP requests for version information
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	v := GetVersionInfo()
	accept := r.Header.Get("Accept")

	switch {
	case strings.Contains(accept, "application/json"):
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(v)
	case strings.Contains(accept, "application/xml"):
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(v)
	default:
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, v.String())
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print version information in various formats`,
	Run: func(cmd *cobra.Command, args []string) {
		v := GetVersionInfo()
		format, _ := cmd.Flags().GetString("format")

		switch format {
		case "json":
			json.NewEncoder(cmd.OutOrStdout()).Encode(v)
		case "xml":
			xml.NewEncoder(cmd.OutOrStdout()).Encode(v)
		default:
			fmt.Fprint(cmd.OutOrStdout(), v.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().StringP("format", "f", "", "Output format (json|xml)")
}
