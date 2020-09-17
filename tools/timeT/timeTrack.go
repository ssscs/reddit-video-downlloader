package timeT

import(
	"time"
	"log"
)

func TimeTrack(start time.Time, name string) {
	log.Printf("%s took %s", name, time.Since(start))
}
