package main

import (
	"log"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"github.com/whalelogic/howtogo/handlers"
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

	r.Static("/css", "./static/css")
	r.Static("/icons", "./static/icons")

	r.GET("/health", handlers.HealthCheckHandler) 
	r.GET("/", handlers.HomePageHandler)
	r.GET("/hello-world", handlers.HelloWorldHandler)
	r.GET("/values", handlers.ValuesHandler)
	r.GET("/variables", handlers.VariablesHandler)
	r.GET("/constants", handlers.ConstantsHandler)


	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
