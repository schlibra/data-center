package cmd

import (
	"embed"
	"log"
	"os"
	"path/filepath"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

var EmbedFS embed.FS

func Execute(embedFs embed.FS) {
	EmbedFS = embedFs
	if !service.Interactive() {
		RunService()
		return
	}

	var rootCmd = &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "Data-Center",
		Long:  "Data-Center",
	}
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(serviceCmd)
	rootCmd.AddCommand(testCmd)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
