package cmd

import (
	"bytes"
	"testing"

	"github.com/cloudtesting/internal/version"
	"github.com/spf13/cobra"
)

func TestVersionCmd(t *testing.T) {
	// Save original version and command state
	originalVersion := version.Version
	originalRoot := rootCmd
	defer func() {
		version.Version = originalVersion
		rootCmd = originalRoot
	}()

	// Reset for test, using "unknown" as default
	version.Version = "unknown"
	rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(versionCmd)

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
			wantContent: `{"version":"unknown"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			versionCmd.SetOut(buf)
			versionCmd.SetArgs([]string{"--format", tt.format})

			if err := versionCmd.Execute(); err != nil {
				t.Fatalf("failed to execute version command: %v", err)
			}

			got := buf.String()
			if !bytes.Contains(buf.Bytes(), []byte(tt.wantContent)) {
				t.Errorf("version command output = %q, want %q", got, tt.wantContent)
			}
		})
	}
}
