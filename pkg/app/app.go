package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/dgarcia202/service-template/internal/cmd"
	"github.com/dgarcia202/service-template/internal/logging"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// App represents the service
type App struct {
	ServiceName      string
	ShortDescription string
	LongDescription  string
	Version          string

	ginEngine       *gin.Engine
	routeSetupFuncs []func(*gin.Engine)
}

var defaultApp App

var serveHandler = func(cmd *cobra.Command, args []string) {

	logging.SetupLogger()
	defaultApp.ginEngine = gin.Default()
	defaultApp.ginEngine.Use(logging.ApplicationFileLogger())

	for _, fn := range defaultApp.routeSetupFuncs {
		fn(defaultApp.ginEngine)
	}

	r := defaultApp.ginEngine
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	log.Info("Starting service...")
	r.Run(fmt.Sprintf("%s:%s", viper.GetString("address"), viper.GetString("port")))
}

// Instance returns a pointer to the created app
func Instance() *App {
	return &defaultApp
}

// Run runs the app either bringing up the service or other action like showing version number
func (a App) Run() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig.String() == "interrupt" {
				fmt.Println()
				log.Info("Shutting down service...")
				os.Exit(0)
			}
		}
	}()

	info := cmd.ServiceInfo{Name: a.ServiceName, Short: a.ShortDescription, Long: a.LongDescription, Version: a.Version}
	cmd.Execute(&info, serveHandler)
}

// SetupRoutes allows to modify routing configuration
func (a App) SetupRoutes(fn func(*gin.Engine)) {
	a.routeSetupFuncs = append(a.routeSetupFuncs, fn)
}

// Shutdown releases resources on application shutdown
func (a App) Shutdown() {
	// Perform clean up
}
