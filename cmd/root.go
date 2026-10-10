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
var I18nFS embed.FS

func Execute(embedFs embed.FS, i18nFS embed.FS) {
	EmbedFS = embedFs
	I18nFS = i18nFS

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
	rootCmd.AddCommand(initCmd)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
