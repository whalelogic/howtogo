package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/whalelogic/howtogo/handlers"
)

func main() {
	r := gin.Default()
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatalf("failed to set trusted proxies: %v", err)
	}

	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/icons", "./static/icons")
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.StaticFile("/robots.txt", "./static/robots.txt")
	r.StaticFile("/sitemap.xml", "./static/sitemap.xml")

	r.GET("/health", handlers.HealthCheckHandler)
	r.GET("/", handlers.HomePageHandler)
	r.GET("/hello-world", handlers.HelloWorldHandler)
	r.GET("/values", handlers.ValuesHandler)
	r.GET("/variables", handlers.VariablesHandler)
	r.GET("/constants", handlers.ConstantsHandler)
	r.GET("/for", handlers.ForHandler)
	r.GET("/ifelse", handlers.IfElseHandler)
	r.GET("/arrays", handlers.ArraysHandler)
	r.GET("/switch", handlers.SwitchHandler)
	r.GET("/slices", handlers.SlicesHandler)
	r.GET("/maps", handlers.MapsHandler)
	r.GET("/range", handlers.RangeHandler)
	r.GET("/functions", handlers.FunctionsHandler)
	r.GET("/multiple-return-values", handlers.MultipleReturnValuesHandler)
	r.GET("/variadic-functions", handlers.VariadicFunctionsHandler)
	r.GET("/closures", handlers.ClosuresHandler)
	r.GET("/recursion", handlers.RecursionHandler)
	r.GET("/pointers", handlers.PointersHandler)
	r.GET("/strings", handlers.StringsHandler)
	r.GET("/runes", handlers.RunesHandler)
	r.GET("/structs", handlers.StructsHandler)
	r.GET("/methods", handlers.MethodsHandler)
	r.GET("/interfaces", handlers.InterfacesHandler)
	r.GET("/generics", handlers.GenericsHandler)
	r.GET("/errors", handlers.ErrorsHandler)
	r.GET("/goroutines", handlers.GoroutinesHandler)
	r.GET("/channels", handlers.ChannelsHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
