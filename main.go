package main

import (
	"os"
	"log"
	"fmt"
	"math/rand"
	"time"

	"github.com/nlopes/slack"
	"github.com/joho/godotenv"
)

func main () {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	apiKey, ok := os.LookupEnv("SLACK_TOKEN")
	if !ok {
		log.Fatalln("API key not set in SLACK_TOKEN")
	}

	api := slack.New(apiKey)

	channels, err := api.GetChannels(true)
	if err != nil {
		log.Fatalln(err)
	}

	channelNames := make([]string, len(channels))
	for i, channel := range channels {
		channelName := channel.GroupConversation.Conversation.NameNormalized
		channelNames[i] = channelName
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(channelNames))
	todaysRecommendChannel := channelNames[index]
	fmt.Printf("今日のおすすめSlackチャンネルは……これ！！！！👉👉👉👉👉 #%s\n", todaysRecommendChannel)
}
