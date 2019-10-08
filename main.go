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

func postMessage(api *slack.Client, channelID, text string) (string, string, error) {
	params := slack.NewPostMessageParameters()
	params.LinkNames = 1
	msgOption := slack.MsgOptionCompose(
		slack.MsgOptionText(text, false),
		slack.MsgOptionPostMessageParameters(params),
	)
	return api.PostMessage(channelID, msgOption)
}

func main () {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	apiKey, ok := os.LookupEnv("SLACK_TOKEN")
	if !ok {
		log.Fatalln("API key not set in SLACK_TOKEN")
	}
	postChannelID, ok :=os.LookupEnv("POST_CHANNEL_ID")
	if !ok {
		log.Fatalln("Post destination channel ID not set in POST_CHANNEL_ID")
	}

	api := slack.New(apiKey)

	channels, err := api.GetChannels(true)
	if err != nil {
		log.Fatalln(err)
	}

	channelNames := make([]string, len(channels))
	for i, channel := range channels {
		channelName := channel.GroupConversation.Conversation.NameNormalized
		channelNames[i] = fmt.Sprintf("#%s", channelName)
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(channelNames))
	todaysRecommendChannel := channelNames[index]
	text := fmt.Sprintf("今日のおすすめSlackチャンネルは……これ！！！！👉👉👉👉👉 %s\n", todaysRecommendChannel)
	postMessage(api, postChannelID, text)
}
