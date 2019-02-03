package logging

import (
	"fmt"
	"os"

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

	log.SetLevel(log.InfoLevel)
}
