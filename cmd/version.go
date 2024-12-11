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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
