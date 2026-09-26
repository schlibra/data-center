package server

import (
	"data-center/internal/router"
	"data-center/pkg/response"
	"data-center/pkg/utils"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"

	"github.com/gin-contrib/cors"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
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
func initStaticMiddlewares(router *gin.Engine, embedFS embed.FS) {
	staticFS, err := static.EmbedFolder(embedFS, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}
	router.Use(static.Serve("/", staticFS))
}
func initFrontendRouter(router *gin.Engine, embedFS embed.FS) {
	subFS, err := fs.Sub(embedFS, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}
	router.GET("/assets/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath")
		data, err := fs.ReadFile(subFS, path.Join("assets", filepath))
		if err != nil {
			c.Status(404)
			return
		}
		contentType := func() string {
			if strings.HasSuffix(filepath, ".css") {
				return "text/css"
			}
			if strings.HasSuffix(filepath, ".js") {
				return "text/javascript"
			}
			return http.DetectContentType(data)
		}()
		c.Data(200, contentType, data)
	})
	router.GET("/favicon.ico", func(c *gin.Context) {
		data, err := fs.ReadFile(subFS, "favicon.ico")
		if err != nil {
			c.Status(404)
			return
		}
		c.Data(200, "image/x-icon", data)
	})
	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			response.SendJson(c, 404, "Api not found", nil)
			return
		}
		data, err := fs.ReadFile(subFS, "index.html")
		if err != nil {
			c.String(500, "index.html not found")
			return
		}
		c.Data(200, "text/html; charset=utf-8", data)
	})
}
func Run(embedFS embed.FS, i18nFS embed.FS) {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	host := config.Server.Host
	port := config.Server.Port
	debug := config.Server.Debug
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	initMiddlewares(server)
	initI18nMiddlewares(server, i18nFS)
	initStaticMiddlewares(server, embedFS)
	initFrontendRouter(server, embedFS)
	router.InitRouter(server)
	if !debug {
		fmt.Printf("[GIN] Listening and serving HTTP on %s:%d\n", host, port)
	}
	err = server.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal(err)
	}
}
func Shutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
}
