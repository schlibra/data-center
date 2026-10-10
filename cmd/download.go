package cmd

import (
	"data-center/internal/models"
	"data-center/pkg/utils"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download latest version",
	Run: func(cmd *cobra.Command, args []string) {
		type Result models.GithubRelease
		var result Result
		_, err := resty.
			New().
			R().
			SetResult(&result).
			Get("https://api.github.com/repos/schlibra/data-center/releases/latest")
		if err != nil {
			log.Fatal(err)
		}
		if result.TagName == Version {
			log.Println("Already up to date")
		}
		fileExt := func() string {
			if runtime.GOOS == "windows" {
				return ".exe"
			}
			return ""
		}()
		for _, item := range result.Assets {
			if item.Name == fmt.Sprintf("data-center-%s-%s%s", runtime.GOOS, runtime.GOARCH, fileExt) {
				log.Printf("Downloading %s\n", item.Name)
				err = utils.DownloadWithProgressBar(item.BrowserDownloadUrl, item.Name)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("Download complete\n")
				os.Exit(0)
			}
		}
		log.Fatal("No matching asset found")
	},
}
