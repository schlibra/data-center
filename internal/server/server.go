package server

import (
	"data-center/internal/router"
	"data-center/pkg/utils"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func initMiddlewares(router *gin.Engine) {
	router.Use(cors.Default())
	router.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		RootPath:         "./i18n",
		AcceptLanguage:   []language.Tag{language.English, language.Chinese},
		DefaultLanguage:  language.Chinese,
		UnmarshalFunc:    yaml.Unmarshal,
		FormatBundleFile: "yaml",
	})))
}

func Run() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	server := gin.Default()
	initMiddlewares(server)
	router.InitRouter(server)
	err = server.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
func Shutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
}
