package app

import (
	"os"
	"os/signal"

	"github.com/dgarcia202/service-template/internal/cmd"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// App represents the service
type App struct {
	ServiceName      string
	ShortDescription string
	LongDescription  string
	Version          string

	ginEngine       *gin.Engine
	routeSetupFuncs []func(*gin.Engine)

	db *gorm.DB
}

var defaultApp App

// Instance returns a pointer to the created app
func Instance() *App {
	return &defaultApp
}

// Run runs the app either bringing up the service or other action like showing version number
func (a *App) Run() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig.String() == "interrupt" {
				shutDown()
			}
		}
	}()

	info := cmd.ServiceInfo{Name: a.ServiceName, Short: a.ShortDescription, Long: a.LongDescription, Version: a.Version}
	cmd.Execute(&info, startUp)
}

// SetupRoutes allows to modify routing configuration
func (a *App) SetupRoutes(fn func(*gin.Engine)) {
	log.Trace("Registering custom route config function")
	a.routeSetupFuncs = append(a.routeSetupFuncs, fn)
}

// Shutdown releases resources on application shutdown
func (a *App) Shutdown() {
	// Perform clean up
}
