package app

import (
	"os"
	"os/signal"

	"github.com/dgarcia202/service-template/internal/cmd"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// HTTPSetupFunc is a function that is able to configure the web server in terms
// of routes and middleware
type HTTPSetupFunc func(*gin.Engine)

// App represents the service
type App struct {
	serviceName      string
	shortDescription string
	longDescription  string
	version          string

	httpSetupFuncs []HTTPSetupFunc
	models         []interface{}

	ginEngine *gin.Engine
	db        *gorm.DB
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
func (a *App) addHTTPSetup(fn HTTPSetupFunc) {
	a.httpSetupFuncs = append(a.httpSetupFuncs, fn)
}

func (a *App) addModel(value interface{}) {
	a.models = append(a.models, value)
}
