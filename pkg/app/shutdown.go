package app

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func shutDown() {
	fmt.Println()
	log.Info("Shutting down service...")

	std.deregisterConsulService()

	log.Trace("Closing database connection")
	err := std.db.Close()
	if err != nil {
		log.Error(err)
	}

	os.Exit(0)
}
