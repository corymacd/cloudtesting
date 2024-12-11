package cmd

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"testing"

	"github.com/cloudtesting/internal/version"
	"github.com/spf13/cobra"
)

func TestVersionCmd(t *testing.T) {
	// Save original version
	originalVersion := version.Version
	defer func() { version.Version = originalVersion }()

	tests := []struct {
		name        string
		format      string
		wantContent string
	}{
		{
			name:        "default format",
			format:      "",
			wantContent: "Version: unknown",
		},
		{
			name:        "explicit plain text",
			format:      "text",
			wantContent: "Version: unknown",
		},
		{
			name:        "json",
			format:      "json",
			wantContent: `{"version":"unknown"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new command for each test
			cmd := &cobra.Command{Use: "app"}
			vCmd := &cobra.Command{
				Use:   "version",
				Short: "Print version information",
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
			vCmd.Flags().StringP("format", "f", "", "Output format (json|xml)")
			cmd.AddCommand(vCmd)

			// Set up output capture
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)

			// Execute command
			cmd.SetArgs([]string{"version"})
			if tt.format != "" {
				cmd.SetArgs([]string{"version", "--format", tt.format})
			}

			err := cmd.Execute()
			if err != nil {
				t.Fatalf("failed to execute version command: %v", err)
			}

			if !strings.Contains(buf.String(), tt.wantContent) {
				t.Errorf("version command output = %q, want it to contain %q", buf.String(), tt.wantContent)
			}
		})
	}
}
