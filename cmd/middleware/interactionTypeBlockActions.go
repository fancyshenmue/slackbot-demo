package middleware

import (
	"fmt"

	"github.com/slack-go/slack/socketmode"
)

func InteractionTypeBlockActions(evt *socketmode.Event, client *socketmode.Client) {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("\t-- middlewareInteractionTypeBlockActions --\n")
	fmt.Println("----------------------------------------------------------------------")
}
