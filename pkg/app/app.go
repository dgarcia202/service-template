package app

import (
	"os"
	"os/signal"

	"github.com/dgarcia202/service-template/internal/cmd"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// App represents the service
type App struct {
	serviceName      string
	shortDescription string
	longDescription  string
	version          string

	ginEngine       *gin.Engine
	routeSetupFuncs []func(*gin.Engine)

	db *gorm.DB
}

var std App

// runs the app either bringing up the service or other action like showing version number
func (a *App) run() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig.String() == "interrupt" {
				shutDown()
			}
		}
	}()

	info := cmd.ServiceInfo{Name: a.serviceName, Short: a.shortDescription, Long: a.longDescription, Version: a.version}
	cmd.Execute(&info, startUp)
}

// setupRoutes allows to modify routing configuration
func (a *App) setupRoutes(fn func(*gin.Engine)) {
	a.routeSetupFuncs = append(a.routeSetupFuncs, fn)
}

// shutdown releases resources on application shutdown
func (a *App) shutdown() {
	// Perform clean up
}
