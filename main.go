package main

import (
	"github.com/AyodejiO/okteto/app"
	"github.com/AyodejiO/okteto/logger"
)

func main() {
	logger.Info("Starting app...")
	app.Start()
}