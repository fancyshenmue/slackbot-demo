package middleware

import (
	"context"
	"fmt"
	"log"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"go.mongodb.org/mongo-driver/bson"

	customDatabaseApp "slackbot-demo/cmd/database"
	customDatabase "slackbot-demo/pkg/database"
)

func Interactive(evt *socketmode.Event, client *socketmode.Client) {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("\t-- middlewareInteractive --\n")
	fmt.Println("----------------------------------------------------------------------")

	callback, ok := evt.Data.(slack.InteractionCallback)
	if !ok {
		fmt.Printf("Ignored %+v\n", evt)
		return
	}

	var payload interface{}

	switch callback.Type {
	case slack.InteractionTypeBlockActions:
		fmt.Println("----------------------------------------------------------------------")
		fmt.Printf("\t-- slack.InteractionTypeBlockActions --\n")
		fmt.Println("----------------------------------------------------------------------")

		// get trade strategy
		var (
			actionID          = callback.ActionCallback.BlockActions[0].ActionID
			selectOptionValue = callback.ActionCallback.BlockActions[0].SelectedOption.Value
			filter            = bson.M{
				"idx": "00000000",
			}
			mongoClient = customDatabaseApp.MongoClient
			database    = customDatabase.MongoDBDatabase
			collection  = customDatabase.MongoDBCollectionBotOperator
		)

		// MongoDB Client
		c := &customDatabase.MongoOperatorStruct{
			Client:     mongoClient,
			Collection: collection,
		}

		// find documents
		var res OptionsStruct
		coll := c.Client.Database(customDatabase.MongoDBDatabase).Collection(customDatabase.MongoDBCollectionBotOperator)
		err := coll.FindOne(context.TODO(), filter).Decode(&res)
		if err != nil {
			log.Printf("Find Documents Error: %v", err)
		}

		switch {
		// operator Terraform
		case actionID == "operator":
			fmt.Println("----------------------------------------------------------------------------------------------------")
			fmt.Printf("\tactionID == \"%s\" && \"selectOptionValue\" == \"%s\"\n", actionID, selectOptionValue)
			fmt.Println("----------------------------------------------------------------------------------------------------")

			// insert documents
			operator := customDatabaseApp.DataInsertStruct{
				Client:     mongoClient,
				Database:   database,
				Collection: collection,
				Field:      actionID,
				Data:       selectOptionValue,
				Filter:     filter,
			}
			operator.DataInsert()

			msg := slack.NewBlockMessage(
				SlackBlockKitHeader(selectOptionValue),
				SlackBlockKitDiv(),
				OperatorOptions(selectOptionValue),
				DryRun(),
				SlackBlockKitDiv(),
				SubmitButton(),
			)
			// Post Message
			_, _, err := client.Client.PostMessage(callback.Channel.ID, slack.MsgOptionBlocks(msg.Msg.Blocks.BlockSet...))
			if err != nil {
				fmt.Println("--------------------------------------------------")
				fmt.Printf("\nfailed posting message with payload: %v\n", err)
				fmt.Println("--------------------------------------------------")
			} else {
				fmt.Println("--------------------------------------------------")
				log.Printf("\tMessage posted successfully.\n")
				fmt.Println("--------------------------------------------------")
			}

			client.Ack(*evt.Request, nil)

			// submit
		case actionID == "submit" && callback.ActionCallback.BlockActions[0].BlockID == "submit_button":
			// debug
			fmt.Println("--------------------------------------------------")
			fmt.Printf("\t-- submit --\n")
			fmt.Println("--------------------------------------------------")

			// // insert documents
			// var m interface{}
			// json.Unmarshal(callback.RawState, &m)
			// coll := c.Client.Database(customDatabase.MongoDBDatabase).Collection(customDatabase.MongoDBCollectionBotTransaction)
			// opts := options.Update().SetUpsert(true)
			// updateField := make(bson.M)
			// updateField["param"] = m.(map[string]interface{})["values"].(map[string]interface{})["param"].(map[string]interface{})["param_input"].(map[string]interface{})["value"]
			// updateField["value"] = m.(map[string]interface{})["values"].(map[string]interface{})["value"].(map[string]interface{})["value_input"].(map[string]interface{})["value"]
			// documents := bson.M{
			// 	"$set": updateField,
			// }
			// _, err := coll.UpdateOne(context.TODO(), filter, documents, opts)
			// if err != nil {
			// 	log.Fatalf("Insert Documents Failed: %v", err)
			// }

			// find documents
			err = coll.FindOne(context.TODO(), filter).Decode(&res)
			if err != nil {
				log.Printf("Find Documents Error: %v", err)
			}

			res = DataGenerator(res)

			InteractiveSubmitResponseOperator(
				client,
				callback,
				OptionsStruct{
					Operator: res.Operator,
					Action:   res.Action,
					DryRun:   res.DryRun,
				},
			)

			if res.DryRun == "false" {
				coll.DeleteOne(context.TODO(), filter)
			}

		default:
			fmt.Println("----------------------------------------------------------------------")
			fmt.Printf("\tdefault\n")
			fmt.Println("----------------------------------------------------------------------")
			// insert documents
			operator := customDatabaseApp.DataInsertStruct{
				Client:     mongoClient,
				Database:   database,
				Collection: collection,
				Field:      actionID,
				Data:       selectOptionValue,
				Filter:     filter,
			}
			operator.DataInsert()

		}

		// Delete operator message
		if actionID == "operator" {
			_, _, err = client.Client.DeleteMessage(callback.Channel.ID, callback.Message.Timestamp)
			if err != nil {
				fmt.Println("--------------------------------------------------")
				fmt.Printf("\tfailed deleting message with payload: %v\n", err)
				fmt.Println("--------------------------------------------------")
			} else {
				fmt.Println("--------------------------------------------------")
				log.Printf("\tDelete message successfully.\n")
				fmt.Println("--------------------------------------------------")
			}
		}

	case slack.InteractionTypeShortcut:
		fmt.Println("--------------------------------------------------")
		fmt.Printf("\tInteractionTypeShortcut: %+v\n", slack.InteractionTypeShortcut)
		fmt.Println()

	case slack.InteractionTypeViewSubmission:
		fmt.Println("--------------------------------------------------")
		fmt.Printf("\tInteractionTypeViewSubmission: %+v\n", slack.InteractionTypeViewSubmission)
		fmt.Println("--------------------------------------------------")

	case slack.InteractionTypeDialogSubmission:
		fmt.Println("--------------------------------------------------")
		fmt.Printf("InteractionTypeDialogSubmission: %+v\n", slack.InteractionTypeDialogSubmission)
		fmt.Println("--------------------------------------------------")

	default:
		fmt.Println("--------------------------------------------------")
		fmt.Printf("\tdefault\n")
		fmt.Println("--------------------------------------------------")

	}

	client.Ack(*evt.Request, payload)
}
