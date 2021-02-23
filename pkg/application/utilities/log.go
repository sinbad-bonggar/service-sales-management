package utilities

import (
	"log"
)

// Log ...
func Log(s string, v ...interface{}) {
	log.Printf(s, v...)
}
