package main

import (
	"fmt"
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/whalelogic/howtogo/database"
	"github.com/whalelogic/howtogo/handlers"
)

func main() {
	// Optional flags for database configuration
	dbPath := flag.String("db", "./database/analytics.db", "Path to SQLite database")
	flag.Parse()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize the database
	log.Printf("Connecting to database at: %s", *dbPath)
	store, err := database.InitAnalyticsDB(*dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer store.DB.Close()

	h := handlers.New(store)

	if os.Getenv("GIN_MODE") != "release" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	fmt.Println("Gin mode:", gin.Mode())
	mw := gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s\" \"%s\" %s\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
	r.Use(mw)

	// Caddy will be running as a reverse proxy, so we need to trust it
	err = r.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		log.Fatalf("failed to set trusted proxies: %v", err)
	}

	// Middleware
	r.Use(h.AnalyticsMiddleware())

	// Static files
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/icons", "./static/icons")
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.StaticFile("/robots.txt", "./static/robots.txt")
	r.StaticFile("/sitemap.xml", "./static/sitemap.xml")

	// Routes
	r.GET("/", handlers.HomePageHandler)
	r.GET("/analytics", h.AnalyticsPageHandler)
	r.GET("/health", handlers.HealthCheckHandler)
	r.GET("/about", handlers.AboutPageHandler)
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
	r.GET("/channel-buffering", handlers.ChannelBufferingHandler)
	r.GET("/channel-synchronization", handlers.ChannelSynchronizationHandler)
	r.GET("/select", handlers.SelectHandler)
	r.GET("/mutexes", handlers.MutexesHandler)
	r.GET("/timers", handlers.TimersHandler)
	r.GET("/timeouts", handlers.TimeoutsHandler)
	r.GET("/tickers", handlers.TickersHandler)
	r.GET("/waitgroups", handlers.WaitGroupsHandler)
	r.GET("/worker-pools", handlers.WorkerPoolsHandler)
	r.GET("/stateful-goroutines", handlers.StatefulGoroutinesHandler)
	r.GET("/sorting", handlers.SortingHandler)
	r.GET("/panic", handlers.PanicHandler)
	r.GET("/recover", handlers.RecoverHandler)
	r.GET("/defer", handlers.DeferHandler)
	r.GET("/json", handlers.JSONHandler)
	r.GET("/xml", handlers.XMLHandler)
	r.GET("/time", handlers.TimeHandler)
	r.GET("/context", handlers.ContextHandler)
	r.GET("/exit", handlers.ExitHandler)
 

	// Start the server using the configured port
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
