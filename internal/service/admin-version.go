package service

import (
	"bufio"
	"data-center/internal/models"
	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/inconshreveable/go-update"
)

func GetAdminVersionService(c *gin.Context, version string, commit string, buildTime string) {
	var result models.GithubRelease
	_, err := resty.New().R().SetResult(&result).Get("https://api.github.com/repos/schlibra/data-center/releases/latest")
	var latestVersion string
	if err != nil {
		latestVersion = "unknown"
	}
	latestVersion = result.TagName
	sendJson(c, 200, "success", H{
		"version":       version,
		"commit":        commit,
		"buildTime":     buildTime,
		"latestVersion": latestVersion,
	})
}
func UpgradeAdminVersionService(c *gin.Context) {
	fileExt := func() string {
		if runtime.GOOS == "windows" {
			return ".exe"
		}
		return ""
	}()
	open, err := os.Open(fmt.Sprintf("data-center-%s-%s%s", runtime.GOOS, runtime.GOARCH, fileExt))
	if err != nil {
		sendJson(c, 400, "更新文件不存在", nil)
		return
	}
	defer func(open *os.File) {
		_ = open.Close()
	}(open)
	reader := bufio.NewReader(open)
	err = update.Apply(reader, update.Options{})
	if err != nil {
		sendError(c, err)
	}
	sendJson(c, 200, "更新成功", nil)
}
