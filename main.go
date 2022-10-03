package main

import (
	"github.com/pedramaghasian/hexagonal-in-go/app"
	"github.com/pedramaghasian/hexagonal-in-go/logger"
)

func main() {

	logger.Info("starting the application...")
	app.Start()

}
