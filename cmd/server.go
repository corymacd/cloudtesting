/*
Copyright © 2022 Cory M. MacDonald <corymacd@netsrv.co>
*/
package cmd

import (
	"context"
	"log"

	"github.com/cloudtesting/internal/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the HTTP server",
	Long:  `Starts an HTTP server on port 8080 with health check and version endpoints`,
	Run:   runServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func runServer(cmd *cobra.Command, args []string) {
	srv := server.New()
	if err := srv.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
