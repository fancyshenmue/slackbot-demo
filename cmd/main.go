package main

import (
	"log"
	"os"
	"strings"

	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"

	"github.com/slack-go/slack"

	slackMiddleware "slackbot-demo/cmd/middleware"
)

func main() {
	appToken := os.Getenv("SLACK_APP_TOKEN")
	if appToken == "" {
		panic("SLACK_APP_TOKEN must be set.\n")
	}

	if !strings.HasPrefix(appToken, "xapp-") {
		panic("SLACK_APP_TOKEN must have the prefix \"xapp-\".")
	}

	botToken := os.Getenv("SLACK_BOT_TOKEN")
	if botToken == "" {
		panic("SLACK_BOT_TOKEN must be set.\n")
	}

	if !strings.HasPrefix(botToken, "xoxb-") {
		panic("SLACK_BOT_TOKEN must have the prefix \"xoxb-\".")
	}

	api := slack.New(
		botToken,
		// slack.OptionDebug(true),
		// slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
		slack.OptionAppLevelToken(appToken),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	// connect
	socketmodeHandler := socketmode.NewSocketmodeHandler(client)
	socketmodeHandler.Handle(socketmode.EventTypeConnecting, slackMiddleware.Connecting)
	socketmodeHandler.Handle(socketmode.EventTypeConnectionError, slackMiddleware.ConnectionError)
	socketmodeHandler.Handle(socketmode.EventTypeConnected, slackMiddleware.Connected)

	// EventTypeEventsAPI //
	// Handle all EventsAPI
	socketmodeHandler.Handle(socketmode.EventTypeEventsAPI, slackMiddleware.EventsAPI)

	// Handle a specific event from EventsAPI
	socketmodeHandler.HandleEvents(slackevents.AppMention, slackMiddleware.AppMentionEvent)

	// EventTypeInteractive //
	// Handle all Interactive Events
	socketmodeHandler.Handle(socketmode.EventTypeInteractive, slackMiddleware.Interactive)

	// Handle a specific Interaction
	socketmodeHandler.HandleInteraction(slack.InteractionTypeBlockActions, slackMiddleware.InteractionTypeBlockActions)

	// Handle all SlashCommand
	socketmodeHandler.Handle(socketmode.EventTypeSlashCommand, slackMiddleware.SlashCommand)
	socketmodeHandler.HandleSlashCommand("/demo", slackMiddleware.SlashCommand)

	socketmodeHandler.HandleDefault(slackMiddleware.Default)

	socketmodeHandler.RunEventLoop()
}
