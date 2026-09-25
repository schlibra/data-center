package cmd

import (
	"data-center/internal/server"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(EmbedFS)
	},
}
