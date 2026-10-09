package server

import (
	"data-center/internal/router"
	"data-center/pkg/utils"
	"embed"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func initMiddlewares(router *gin.Engine) {
	router.Use(cors.Default())

}
func initI18nMiddlewares(router *gin.Engine, i18nFS embed.FS) {
	router.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		AcceptLanguage:   []language.Tag{language.English, language.Chinese},
		DefaultLanguage:  language.Chinese,
		UnmarshalFunc:    yaml.Unmarshal,
		FormatBundleFile: "yaml",
		RootPath:         "i18n",
		Loader: &ginI18n.EmbedLoader{
			FS: i18nFS,
		},
	})))
}
func initSwagger(router *gin.Engine) {
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
func Run(embedFS embed.FS, i18nFS embed.FS) {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	debug := config.Server.Debug
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	initSwagger(server)
	initMiddlewares(server)
	initI18nMiddlewares(server, i18nFS)
	initStaticMiddlewares(server, embedFS)
	initFrontendRouter(server, embedFS)
	router.InitRouter(server)
	listenServer(server, config)
}
func Shutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
}
