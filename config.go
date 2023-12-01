package main

import (
	"time"
)

func init() {
	time.Local = time.UTC
}
