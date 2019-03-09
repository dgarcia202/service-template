package app

import (
	"fmt"
	"net/http"

	"github.com/dgarcia202/service-template/internal/logging"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func startUp(cmd *cobra.Command, args []string) {

	log.Trace("Setting up logging artifacts")
	logging.SetupLogger()

	log.Trace("Creating Gin engine")
	defaultApp.ginEngine = gin.Default()

	log.Trace("Adding logging middleware")
	defaultApp.ginEngine.Use(logging.ApplicationFileLogger())

	log.Trace("Adding custom routes")
	fnCount := 0
	for _, fn := range defaultApp.routeSetupFuncs {
		log.Trace("Executing 'SetupRoutes' function")
		fn(defaultApp.ginEngine)
		fnCount++
	}
	log.Trace(fnCount, " custom route functions processed")

	db, err := gorm.Open("sqlite3", "temp/test.db")
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	defaultApp.db = db

	r := defaultApp.ginEngine
	log.Trace("Adding default '/ping' route")
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	log.Info("Starting service...")
	r.Run(fmt.Sprintf("%s:%s", viper.GetString("address"), viper.GetString("port")))
}
