package middleware

import (
	"fmt"

	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func AppMentionEvent(evt *socketmode.Event, client *socketmode.Client) {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("\t-- middlewareAppMentionEvent --\n")
	fmt.Println("----------------------------------------------------------------------")

	eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
	if !ok {
		fmt.Printf("Ignored %+v\n", evt)
		return
	}
	client.Ack(*evt.Request)

	ev, ok := eventsAPIEvent.InnerEvent.Data.(*slackevents.AppMentionEvent)
	if !ok {
		fmt.Printf("Ignored %+v\n", ev)
		return
	}
}
