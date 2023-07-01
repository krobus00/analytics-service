package utils

import (
	log "github.com/sirupsen/logrus"
)

func ContinueOrFatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
