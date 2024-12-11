/*
Copyright © 2022 Cory M. MacDonald <corymacd@netsrv.co>
This file is part of the intel gathering honeypot application for myipintel.com
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// Version holds the current version of the application
// This is typically set at build time
var Version = "dev"

// VersionFormat defines how the version should be displayed
type VersionFormat int

const (
	// FormatSimple returns just the version number
	FormatSimple VersionFormat = iota
	// FormatPrefixed returns "Version: X.X.X"
	FormatPrefixed
	// FormatDetailed returns a more detailed version string
	FormatDetailed
)

// GetVersion returns the version string in the specified format
func GetVersion(format VersionFormat) string {
	switch format {
	case FormatSimple:
		return Version
	case FormatPrefixed:
		return fmt.Sprintf("Version: %s", Version)
	case FormatDetailed:
		return fmt.Sprintf("Application Version: %s (Built: %s)", Version, "unknown") // You can add build time info later
	default:
		return fmt.Sprintf("Version: %s", Version)
	}
}

// VersionHandler handles HTTP requests for version information
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, GetVersion(FormatPrefixed))
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the detailed version information of the application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(GetVersion(FormatDetailed))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
