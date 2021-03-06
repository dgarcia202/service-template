package logging

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var usingFileSystem = false

// GormLogger encapsulates logrus for use inside GORM
type GormLogger struct{}

// Print logs GORM infomation
func (*GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		log.WithFields(log.Fields{"module": "gorm", "type": "sql"}).Print(v[3])
	}
	if v[0] == "log" {
		log.WithFields(log.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}

// SetupLogger configure logrus default logger according to loaded configuration
func SetupLogger() {
	log.SetFormatter(&log.TextFormatter{})

	fileHandler, err := os.OpenFile(viper.GetString("logfile"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(fileHandler)
		usingFileSystem = true
	} else {
		log.Error(fmt.Sprintf("Failed to log to file, using default stderr (%s)", err))
	}

	lvl, err := log.ParseLevel(viper.GetString("loglevel"))
	if err != nil {
		log.Error("Couldn't parse error level ", viper.GetString("loglevel"), " using default level (INFO)")
		log.SetLevel(log.InfoLevel)
	} else {
		log.Debug("Loglevel is set to ", strings.ToUpper(lvl.String()))
		log.SetLevel(lvl)
	}
}

// ApplicationFileLogger web framework middleware that logs HTTP operations to the main application log
func ApplicationFileLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if usingFileSystem {
			path := c.Request.URL.Path
			raw := c.Request.URL.RawQuery
			c.Next()

			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()
			comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

			log.WithFields(log.Fields{
				"path":   path,
				"raw":    raw,
				"ip":     clientIP,
				"method": method,
				"status": statusCode,
			}).Debug(comment)
		}
	}
}
