package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"github.com/whalelogic/howtogo/templates/pages"
)

func render(c *gin.Context, status int, component templ.Component) {
	c.Status(status)
	if err := component.Render(c.Request.Context(), c.Writer); err != nil {
		// Surface render errors to logs while keeping response simple.
		log.Printf("render error: %v", err)
	}
}

func main() {
	r := gin.Default()
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatalf("failed to set trusted proxies: %v", err)
	}

	r.Static("/css", "./public/css")
	r.Static("/icons", "./public/icons")

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "200 OK\n")
	})

	// test component

	component := pages.HelloWorld()
	r.GET("/test", func(c *gin.Context) {
		render(c, http.StatusOK, component)
	})

	r.GET("/", func(c *gin.Context) {
		render(c, http.StatusOK, pages.Index())
	})

	r.GET("/hello-world", func(c *gin.Context) {
		render(c, http.StatusOK, pages.HelloWorld())
	})

	r.GET("/values", func(c *gin.Context) {
		render(c, http.StatusOK, pages.Values())
	})

	r.GET("/variables", func(c *gin.Context) {
		render(c, http.StatusOK, pages.Variables())
	})

	r.GET("/constants", func(c *gin.Context) {
		render(c, http.StatusOK, pages.Constants())
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
