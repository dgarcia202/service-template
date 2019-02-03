package logging

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// SetupLogger configure logrus default logger according to loaded configuration
func SetupLogger() {
	log.SetFormatter(&log.TextFormatter{})

	file, err := os.OpenFile(viper.GetString("logfile"), os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
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
