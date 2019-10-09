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

func filterChannels(channels []slack.Channel) []slack.Channel {
	var filtered []slack.Channel
	for _, channel := range channels {
		conversation := channel.GroupConversation.Conversation
		if channel.IsGeneral {
			continue
		}
		if conversation.NumMembers > 20 {
			continue
		}
		if conversation.IsPrivate {
			continue
		}
		filtered = append(filtered, channel)
	}
	return filtered
}

func buildText(channel string) string {
	var aisatsu string
	now := time.Now()
	if 5 <= now.Hour() && now.Hour() < 12 {
		aisatsu = "おはようございます！！！"
	} else if 12 <= now.Hour() && now.Hour() < 16 {
		aisatsu = "こんにちは！！！"
	} else if 16 <= now.Hour() && now.Hour() < 23 {
		aisatsu = "こんばんは！！！"
	} else { // midnight
		aisatsu = "夜分に失礼します！！！"
	}
	return fmt.Sprintf("%s 今日のおすすめチャンネルは……これ！！！！👉👉👉👉👉 %s 👈👈👈👈👈", aisatsu, channel)
}

func doIt(dryRun bool) {
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
	channels = filterChannels(channels)
	log.Printf("number of target channels is %d\n", len(channels))

	channelNames := make([]string, len(channels))
	for i, channel := range channels {
		channelName := channel.Name
		channelNames[i] = fmt.Sprintf("#%s", channelName)
	}

	todaysRecommendChannel := chooseChannel(channelNames)
	text := buildText(todaysRecommendChannel)
	log.Printf("message: %s", text)
	if !dryRun {
		log.Printf("post %s to %s", text, postChannelID)
		postMessage(api, postChannelID, text)
	}
}
