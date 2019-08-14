package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	L "github.com/frankielearns/intersectproject/cmd/database"
	"github.com/nlopes/slack"
	"github.com/piquette/finance-go/quote"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func main() {

	token := os.Getenv("SLACKTOKEN")
	api := slack.New(token)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)
				fmt.Println(info.User)
				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
					respond(rtm, ev, prefix)
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}
}

func respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
	var response string
	text := msg.Text
	text = strings.TrimPrefix(text, prefix)
	text = strings.TrimSpace(text)
	re := regexp.MustCompile(`<http://.*\||>`)
	text = re.ReplaceAllString(text, " ")
	stock, err := quote.Get(text)
	L.Databaseconnect(text, stock.RegularMarketPrice, "test")

	if err == nil {
		response = fmt.Sprintln("The stock", stock.ShortName, "is at", FloatToString(stock.RegularMarketPrice))
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	}
}
