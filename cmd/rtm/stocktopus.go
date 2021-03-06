package main

import (
	"fmt"
	"log"
	"os"

	iex "github.com/thorfour/iex/pkg/api"
	"github.com/thorfour/stocktopus/pkg/slack"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: stocktopus: slack-bot-token")
		return
	}

	// Open connection
	slackBot, err := slack.NewRTMClient(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := slackBot.Receive()
		if err != nil {
			log.Fatal(err)
		}

		if len(msg) != 0 {
			quote, err := iex.Price(msg)
			if err != nil {
				continue
			}

			// Post the quote
			slackBot.Send(fmt.Sprintf("Current Price: %v", quote))
		}
	}
}
