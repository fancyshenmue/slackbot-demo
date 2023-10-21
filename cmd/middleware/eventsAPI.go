package middleware

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func EventsAPI(evt *socketmode.Event, client *socketmode.Client) {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("\t-- middlewareEventsAPI --\n")
	fmt.Println("----------------------------------------------------------------------")

	eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
	if !ok {
		fmt.Printf("Ignored %+v\n", evt)
		return
	}

	fmt.Printf("Event received: %+v\n", eventsAPIEvent)

	client.Ack(*evt.Request)

	switch eventsAPIEvent.Type {
	case slackevents.CallbackEvent:
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			// Build Message with blocks created above
			msg := slack.NewBlockMessage(
				SlackBlockKitHeader("TopLevel"),
				SlackBlockKitDiv(),
				Operator(),
			)

			// Post Message
			_, _, err := client.Client.PostMessage(ev.Channel, slack.MsgOptionBlocks(msg.Msg.Blocks.BlockSet...))
			if err != nil {
				fmt.Println()
				fmt.Printf("failed posting message with payload: %v", err)
				fmt.Println()
			} else {
				fmt.Println()
				log.Println("Message posted successfully.")
				fmt.Println()
			}

			client.Ack(*evt.Request, nil)

		case *slackevents.MemberJoinedChannelEvent:
			fmt.Printf("user %q joined to channel %q", ev.User, ev.Channel)
		}

	default:
		client.Debugf("unsupported Events API event received")
	}
}
