package config

import (
	"time"
)

func setTimeZone() {
	time.Local = time.UTC
}
