/*
Copyright © 2022 Cory M. MacDonald <corymacd@netsrv.co>
This file is part of the intel gathering honeypot application for myipintel.com
*/
package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/cloudtesting/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print version information in various formats`,
	Run: func(cmd *cobra.Command, args []string) {
		info := version.GetInfo()
		format, _ := cmd.Flags().GetString("format")

		switch format {
		case "json":
			json.NewEncoder(cmd.OutOrStdout()).Encode(info)
		case "xml":
			xml.NewEncoder(cmd.OutOrStdout()).Encode(info)
		default:
			fmt.Fprintf(cmd.OutOrStdout(), "Version: %s\nGitCommit: %s\nBuildTime: %s\nBuildUser: %s\nGoVersion: %s\n",
				info.Version, info.GitCommit, info.BuildTime, info.BuildUser, info.GoVersion)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().StringP("format", "f", "", "Output format (json|xml)")
}
