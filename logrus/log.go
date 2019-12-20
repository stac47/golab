package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	log.Debug("Another one appears.")
}
