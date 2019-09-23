package profiling

import (
	"log"
	"time"
)

// Duration measurement
func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)
	log.Printf("%s lasted %s", name, elapsed)
}
