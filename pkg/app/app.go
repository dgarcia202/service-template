package app

import (
	"fmt"
	"net/http"

	"github.com/dgarcia202/service-template/internal/cmd"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// App represents the service
type App struct {
	ServiceName      string
	ShortDescription string
	LongDescription  string
	Version          string

	ginEngine *gin.Engine
}

var defaultApp App

var serveHandler = func(cmd *cobra.Command, args []string) {
	r := defaultApp.ginEngine

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(fmt.Sprintf("%s:%s", viper.GetString("address"), viper.GetString("port")))
}

func init() {
	defaultApp.ginEngine = gin.Default()
}

// Instance returns a pointer to the created app
func Instance() *App {
	return &defaultApp
}

// Run runs the app either bringing up the service or other action like showing version number
func (a App) Run() {
	info := cmd.ServiceInfo{Name: a.ServiceName, Short: a.ShortDescription, Long: a.LongDescription, Version: a.Version}
	cmd.Execute(&info, serveHandler)
}
