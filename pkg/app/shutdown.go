package app

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func shutDown() {
	fmt.Println()
	log.Info("Shutting down service...")

	err := defaultApp.db.Close()
	if err != nil {
		log.Error(err)
	}

	os.Exit(0)
}
