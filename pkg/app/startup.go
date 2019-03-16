package app

import (
	"fmt"
	"net/http"

	"github.com/dgarcia202/service-template/internal/logging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func startUp(cmd *cobra.Command, args []string) {

	logging.SetupLogger()
	log.Info("Starting service...")

	dumpConfiguration()

	log.Trace("Creating Gin engine")
	std.ginEngine = gin.Default()

	log.Trace("Adding logging middleware")
	std.ginEngine.Use(logging.ApplicationFileLogger())

	log.Trace("Adding custom routes")
	fnCount := 0
	for _, fn := range std.httpSetupFuncs {
		log.Trace("Executing 'SetupRoutes' function")
		fn(std.ginEngine)
		fnCount++
	}
	log.Trace(fnCount, " custom route functions processed")

	// start relational database
	db, err := gorm.Open(viper.GetString("dbdialect"), viper.GetString("db"))
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	} else {
		log.Info("Database connection successful")
	}

	db.SetLogger(&logging.GormLogger{})
	db.LogMode(true)
	if len(std.models) > 0 {
		db.AutoMigrate(std.models...)
	}

	db.Callback().Create().Before("gorm:create").Register("set_uuid_primary_key", setUUIDPrimaryKey)

	std.db = db

	// start HTTP server
	r := std.ginEngine
	log.Trace("Adding default '/ping' route")
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	std.registerConsulService()
	log.Info("Starting HTTP server at ", viper.GetString("address"), ":", viper.GetString("port"))
	r.Run(fmt.Sprintf("%s:%s", viper.GetString("address"), viper.GetString("port")))
}

func dumpConfiguration() {
	log.Trace("Config file used: ", viper.ConfigFileUsed())
	for _, key := range viper.AllKeys() {
		log.Trace("Config: ", key, " -> ", viper.GetString(key))
	}
}

func setUUIDPrimaryKey(scope *gorm.Scope) {
	if scope.HasColumn("ID") {
		scope.SetColumn("ID", uuid.New().String())
	}
}
