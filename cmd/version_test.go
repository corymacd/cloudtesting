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

	// Reset for test
	version.Version = "1.0.0-test"
	rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(versionCmd)

	tests := []struct {
		name        string
		format      string
		wantContent string
	}{
		{
			name:        "plain text",
			format:      "",
			wantContent: "Version: 1.0.0-test",
		},
		{
			name:        "json",
			format:      "json",
			wantContent: `{"version":"1.0.0-test"`,
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
