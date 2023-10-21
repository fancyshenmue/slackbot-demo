package middleware

import (
	"fmt"

	"github.com/slack-go/slack/socketmode"
)

func SlashCommand(evt *socketmode.Event, client *socketmode.Client) {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("\t-- middleSlashCommand --\n")
	fmt.Println("----------------------------------------------------------------------")

	// cmd, ok := evt.Data.(slack.SlashCommand)
	// if !ok {
	// 	fmt.Printf("Ignored %+v\n", evt)
	// 	return
	// }

	// client.Debugf("Slash command received: %+v", cmd)

	// payload := map[string]interface{}{
	// 	"blocks": []slack.Block{
	// 		slack.NewSectionBlock(
	// 			&slack.TextBlockObject{
	// 				Type: slack.MarkdownType,
	// 				Text: "foo",
	// 			},
	// 			nil,
	// 			slack.NewAccessory(
	// 				slack.NewButtonBlockElement(
	// 					"",
	// 					"somevalue",
	// 					&slack.TextBlockObject{
	// 						Type: slack.PlainTextType,
	// 						Text: "bar",
	// 					},
	// 				),
	// 			),
	// 		),
	// 	}}
	// client.Ack(*evt.Request, payload)
}
