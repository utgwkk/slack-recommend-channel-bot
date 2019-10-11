package main

import (
	"flag"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

var dryRun bool

func init() {
	flag.BoolVar(&dryRun, "dry-run", false, "enable dry-run mode (don't post to Slack)")
}

func main() {
	flag.Parse()
	if (dryRun) {
		log.Println("dry-run mode.")
	}
	doIt(dryRun)
}
