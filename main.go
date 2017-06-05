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
	// Required
	githubToken := GetEnv("GITHUB_TOKEN", "")
	slackToken := GetEnv("SLACK_TOKEN", "")

	// Optional
	githubSearchQuery := GetEnv("GITHUB_QUERY", "is:open is:pr")
	slackChannel := GetEnv("SLACK_CHANNEL", "engineering")
	slackIconEmoji := GetEnv("SLACK_ICON_EMOJI", ":tada:")
	slackUsername := GetEnv("SLACK_USERNAME", "GitHub Reminder")

	flag.Parse()

	text := ""
	prs := GetPullRequests(githubToken, githubSearchQuery)
	if len(prs) > 0 {
		for _, pr := range prs {
			ta := timeago.English.Format(*pr.CreatedAt)
			text = text + "— " + pr.GetTitle()
			text = text + " (" + ta + ") "
			text = text + " [" + pr.GetHTMLURL() + "] "
			text = text + "\n"
		}
		SendToSlack(slackToken, slackChannel, text, slackIconEmoji, slackUsername)
	}
}

// Get the pull requests list
func GetPullRequests(token string, query string) []github.Issue {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	options := &github.SearchOptions{}
	result, _, _ := client.Search.Issues(ctx, query, options)

	return result.Issues
}

// Send text to a Slack chan
func SendToSlack(token string, channel string, text string, iconEmoji string, username string) {
	api := slack.New(token)
	params := slack.PostMessageParameters{}
	params.IconEmoji = iconEmoji
	params.Username = username
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
