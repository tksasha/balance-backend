package main

import (
	"github.com/tksasha/balance/config"
	"github.com/tksasha/balance/internal/app"
)

func main() {
	cfg := config.New()

	app.Run(cfg)
}
