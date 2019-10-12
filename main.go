package main

import (
	"flag"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"time"
)

var dryRun bool
var timeZone string

func init() {
	flag.BoolVar(&dryRun, "dry-run", false, "enable dry-run mode (don't post to Slack)")
	flag.StringVar(&timeZone, "tz", "", "Time zone (default to use local time zone)")
}

func loadTimezone() {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		loc = time.Local
	}
	time.Local = loc
}

func main() {
	flag.Parse()
	if timeZone != "" {
		loadTimezone()
	}
	log.Printf("time zone is %s", time.Local.String())
	if (dryRun) {
		log.Println("dry-run mode.")
	}
	doIt(dryRun)
}
