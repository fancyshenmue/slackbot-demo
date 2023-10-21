package middleware

import (
	"fmt"
	"os"

	"github.com/slack-go/slack/socketmode"
)

func Default(evt *socketmode.Event, client *socketmode.Client) {
	fmt.Println("------------------------------------------------------------------------------------------")
	fmt.Fprintf(os.Stderr, "\tUnexpected event type received: %s\n", evt.Type)
	fmt.Println("------------------------------------------------------------------------------------------")
}
