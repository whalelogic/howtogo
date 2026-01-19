// Package handlers contains the HTTP route handlers for the web application.
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log"

	"github.com/a-h/templ"
	"github.com/whalelogic/howtogo/views/pages"
	"github.com/whalelogic/howtogo/database"
)

// render is a helper function to render a templ.Component with the given status code.
// from the templ.guide examples.
func render(c *gin.Context, status int, component templ.Component) {
	c.Status(status)
	if err := component.Render(c.Request.Context(), c.Writer); err != nil {
		log.Printf("render error: %v", err)
	}
}

// AppHandler holds dependencies for the HTTP handlers.
type AppHandler struct {
	Store *database.AnalyticsStore
}

func New(store *database.AnalyticsStore) *AppHandler {
	return &AppHandler{Store: store}
}

// AnalyticsMiddleware is a Gin middleware that tracks page visits and collects analytics data.
func (h *AppHandler) AnalyticsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		h.Store.RecordVisit(clientIP, c.Request.URL.Path)
		totalViews, uniqueVisitors := h.Store.GetStats()

		c.Set("client_ip", clientIP)
		c.Set("page_views", totalViews)
		c.Set("visitor_count", uniqueVisitors)

		c.Next()
	}
}

// AnalyticsPageHandler renders the analytics page with visit statistics.
func (h *AppHandler) AnalyticsPageHandler(c *gin.Context) {
	ip := c.MustGet("client_ip").(string)
	views := c.MustGet("page_views").(int)
	visitors := c.MustGet("visitor_count").(int)

	render(c, http.StatusOK, pages.Analytics(ip, views, visitors))
}

func AboutPageHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.About())
}



// HealthCheckHandler responds with a simple "200 OK" message for health checks.
func HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "200 OK\n")
}



// Page Handlers

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

func StatefulGoroutinesHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.StatefulGoroutines())
}

func ChannelBufferingHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.ChannelBuffering())
}

func ChannelSynchronizationHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.ChannelSynchronization())
}

func SelectHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Select())
}

func TimeoutsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Timeouts())
}

func TimersHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Timers())
}

func TickersHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Tickers())
}

func WorkerPoolsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.WorkerPools())
}

func MutexesHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Mutexes())
}

func WaitGroupsHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.WaitGroups())
}

func SortingHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Sorting())
}

func PanicHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Panic())
}

func DeferHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Defer())
}

func RecoverHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Recover())
}

func JSONHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.JSON())
}

func XMLHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.XML())
}

func TimeHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Time())
}

func ContextHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Context())
}

func ExitHandler(c *gin.Context) {
	render(c, http.StatusOK, pages.Exit())
}



