package main

import (
	"flag"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	doIt(dryRun)
}
