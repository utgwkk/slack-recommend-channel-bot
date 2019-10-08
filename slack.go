package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/nlopes/slack"
)

func postMessage(api *slack.Client, channelID, text string) (string, string, error) {
	params := slack.NewPostMessageParameters()
	params.LinkNames = 1
	msgOption := slack.MsgOptionCompose(
		slack.MsgOptionUsername("今日のおすすめチャンネル"),
		slack.MsgOptionIconEmoji("tada"),
		slack.MsgOptionText(text, false),
		slack.MsgOptionPostMessageParameters(params),
	)
	return api.PostMessage(channelID, msgOption)
}

func chooseChannel(channelNames []string) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(channelNames))
	return channelNames[index]
}

func doIt() {
	apiKey, ok := os.LookupEnv("SLACK_TOKEN")
	if !ok {
		log.Fatalln("API key not set in SLACK_TOKEN")
	}
	postChannelID, ok := os.LookupEnv("POST_CHANNEL_ID")
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

	todaysRecommendChannel := chooseChannel(channelNames)
	text := fmt.Sprintf("今日のおすすめチャンネルは……これ！！！！👉👉👉👉👉 %s 👈👈👈👈👈", todaysRecommendChannel)
	log.Printf("post %s to %s", text, postChannelID)
	postMessage(api, postChannelID, text)
}
