package middleware

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func InteractiveSubmitResponseOperator(client *socketmode.Client, callback slack.InteractionCallback, option OptionsStruct) {
	msg := slack.NewBlockMessage(
		SlackBlockKitHeader("Action"),
		SlackBlockKitDiv(),
		SlackBlockKitResponse(OptionsStruct{
			Operator: option.Operator,
			Action:   option.Action,
			DryRun:   option.DryRun,
		}),
	)

	// Post Message
	_, _, err := client.Client.PostMessage(callback.Channel.ID, slack.MsgOptionBlocks(msg.Msg.Blocks.BlockSet...))
	if err != nil {
		fmt.Println("--------------------------------------------------")
		fmt.Printf("\tfailed posting message with payload: %v\n", err)
		fmt.Println("--------------------------------------------------")
	} else {
		fmt.Println("--------------------------------------------------")
		log.Printf("\tMessage posted successfully.\n")
		fmt.Println("--------------------------------------------------")
	}
}
