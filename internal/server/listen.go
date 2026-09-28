package server

import (
	"data-center/internal/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"golang.org/x/sync/errgroup"
)

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: false,
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
	sslPort := cfg.Server.SSL.Port
	var g errgroup.Group
	g.Go(func() error {
		if !debug {
			fmt.Printf("[GIN] Listening and serving on http://%s:%d\n", host, port)
		}
		return server.Run(fmt.Sprintf("%s:%d", host, port))

	})
	if cfg.Server.SSL.Enable {
		g.Go(func() error {
			server.Use(TlsHandler())
			if !debug {
				fmt.Printf("[GIN] Listening and serving on https://%s:%d\n", host, sslPort)
			}
			return server.RunTLS(fmt.Sprintf("%s:%d", host, sslPort), cfg.Server.SSL.CertFile, cfg.Server.SSL.KeyFile)
		})
	}
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
