package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/nlopes/slack"
	"github.com/xeonx/timeago"
	"golang.org/x/oauth2"
)

// Main function
func main() {
	githubToken := GetEnv("github-token", "")
	slackToken := GetEnv("slack-token", "")
	slackChannel := GetEnv("slack-channel", "engineering")

	flag.Parse()

	text := ""
	prs := GetPullRequests(githubToken)
	if len(prs) > 0 {
		for _, pr := range prs {
			ta := timeago.English.Format(*pr.CreatedAt)
			text = text + "— " + pr.GetTitle()
			text = text + " (" + ta + ") "
			text = text + " [" + pr.GetHTMLURL() + "] "
			text = text + "\n"
		}
		SendToSlack(slackToken, slackChannel, text)
	}
}

// Get the pull requests list
func GetPullRequests(token string) []github.Issue {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	options := &github.SearchOptions{}
	query := "is:open is:pr user:bureaux-a-partager label:RFR"
	result, _, _ := client.Search.Issues(ctx, query, options)

	return result.Issues
}

// Send text to a Slack chan
func SendToSlack(token string, channel string, text string) {
	api := slack.New(token)
	params := slack.PostMessageParameters{}
	params.IconEmoji = ":tada:"
	params.Username = "GitHub Reminder"
	channelID, timestamp, err := api.PostMessage(channel, text, params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

// Get environment var with a fallback
func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}
