// Package handlers contains the HTTP route handlers for the web application.
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/whalelogic/howtogo/templates/pages"
	"github.com/a-h/templ"
	"log"
)

func render(c *gin.Context, status int, component templ.Component) {
	c.Status(status)
	if err := component.Render(c.Request.Context(), c.Writer); err != nil {
		log.Printf("render error: %v", err)
	}
}

func HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "200 OK\n")
}

func HomePageHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Index())
}

func HelloWorldHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.HelloWorld())
}

func ValuesHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Values())
}

func VariablesHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Variables())
}

func ConstantsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Constants())
}


