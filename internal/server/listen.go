package server

import (
	"data-center/internal/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			log.Fatal(err)
		}
		c.Next()
	}
}

func listenServer(server *gin.Engine, cfg models.Config) {
	debug := cfg.Server.Debug
	host := cfg.Server.Host
	port := cfg.Server.Port
	if cfg.Server.SSL.Enable {
		server.Use(TlsHandler())
		if !debug {
			fmt.Printf("[GIN] Listening and serving on https://%s:%d\n", host, port)
		}
		err := server.RunTLS(fmt.Sprintf("%s:%d", host, port), cfg.Server.SSL.CertFile, cfg.Server.SSL.KeyFile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if !debug {
			fmt.Printf("[GIN] Listening and serving on http://%s:%d\n", host, port)
		}
		err := server.Run(fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			log.Fatal(err)
		}
	}
}
