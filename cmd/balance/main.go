package main

import (
	"github.com/tksasha/balance/config"
	"github.com/tksasha/balance/internal/app"
)

func main() {
	config := config.New()
	app.Run(config)
}
