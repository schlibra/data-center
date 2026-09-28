package server

import (
	"data-center/pkg/response"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

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
	router.GET("/registerSW.js", func(c *gin.Context) {
		data, err := fs.ReadFile(subFS, "registerSW.js")
		if err != nil {
			c.Status(404)
			return
		}
		c.Data(200, "text/javascript; charset=utf-8", data)
	})
	router.GET("/manifest.webmanifest", func(c *gin.Context) {
		data, err := fs.ReadFile(subFS, "manifest.webmanifest")
		if err != nil {
			c.Status(404)
			return
		}
		c.Data(200, "application/json; charset=utf-8", data)
	})
	router.GET("/logo.svg", func(c *gin.Context) {
		data, err := fs.ReadFile(subFS, "favicon.ico")
		if err != nil {
			c.Status(404)
			return
		}
		c.Data(200, "image/svg+xml; charset=utf-8", data)
	})
	router.GET("/logo-192.png", func(c *gin.Context) {
		data, err := fs.ReadFile(subFS, "logo-192.png")
		if err != nil {
			c.Status(404)
			return
		}
		c.Data(200, "image/png; charset=utf-8", data)
	})
	router.GET("/logo-512.png", func(c *gin.Context) {
		data, err := fs.ReadFile(subFS, "logo-512.png")
		if err != nil {
			c.Status(404)
			return
		}
		c.Data(200, "image/png; charset=utf-8", data)
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
