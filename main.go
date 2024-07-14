package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	app := pocketbase.New()

	// fires only for "texts" collections
	app.OnRecordBeforeCreateRequest("texts").Add(func(e *core.RecordCreateEvent) error {
		accountSid := goDotEnvVariable("TWILIO_ACCOUNT_SID")
		authToken := goDotEnvVariable("TWILIO_AUTH_TOKEN")
		var body string = e.Record.Get("message").(string)
		var to string = e.Record.Get("to").(string)
		var twilioNumber string = goDotEnvVariable("TWILIO_NUMBER")

		client := twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		})

		params := &twilioApi.CreateMessageParams{}
		params.SetTo(to)
		params.SetFrom(twilioNumber)
		params.SetBody(body)

		resp, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println("Error sending SMS message: " + err.Error())
		} else {
			response, _ := json.Marshal(*resp)
			fmt.Println("Response: " + string(response))
		}

		log.Println()
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
