// Package handlers contains the HTTP route handlers for the web application.
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log"

	"github.com/a-h/templ"
	"github.com/whalelogic/howtogo/views/pages"
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

func ForHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.For())
}

func IfElseHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.IfElse())
}

func ArraysHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Arrays())
}

func SwitchHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Switch())
}

func SlicesHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Slices())
}

func MapsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Maps())
}

func RangeHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Range())
}

func FunctionsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Functions())
}

func MultipleReturnValuesHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.MultipleReturnValues())
}

func VariadicFunctionsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.VariadicFunctions())
}

func ClosuresHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Closures())
}

func RecursionHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Recursion())
}

func PointersHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Pointers())
}

func StringsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Strings())
}

func RunesHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Runes())
}

func StructsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Structs())
}

func MethodsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Methods())
}

func InterfacesHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Interfaces())
}

func GenericsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Generics())
}

func ErrorsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Errors())
}

func GoroutinesHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Goroutines())
}

func ChannelsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Channels())
}
