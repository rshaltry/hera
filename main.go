package main

import (
	"github.com/rshaltry/hera/certificate"
	"github.com/rshaltry/hera/listener"
	"github.com/rshaltry/hera/logger"
	"github.com/rshaltry/hera/version"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("hera")

func main() {
	logger.Init("hera")

	listener, err := listener.New()
	if err != nil {
		log.Errorf("Unable to start: %s", err)
	}

	log.Infof("Hera v%s has started", version.Current)

	err = certificate.Verify(listener.Fs)
	if err != nil {
		log.Error(err.Error())
	}

	err = listener.Revive()
	if err != nil {
		log.Error(err.Error())
	}

	listener.Listen()
}
