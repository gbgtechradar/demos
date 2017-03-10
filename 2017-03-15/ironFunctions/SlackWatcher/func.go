package main

//Split these to two ironFunctions, one to fetch next meetup, and one to communicate with slack? maybe 3, one to listen to slacka nd one to send message

import (
	"fmt"
	"log"
	"strings"

	"github.com/nlopes/slack"
	"encoding/json"
	"os"
)

type Input struct {
	Token     string
	Trigger   string
}

func getStdIn() (s Input, err error) {
	s = Input{}

	err = json.NewDecoder(os.Stdin).Decode(&s)

	if(err == nil && s.Token=="") {
		err = fmt.Errorf("Please define Slack Token")
	}

	return s, err
}

func main() {
	input, err := getStdIn()

	if(err != nil){
		//Invalid input, bail out
		log.Panic(err)
	}

	api := slack.New(input.Token)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:
			if strings.Contains(ev.Text, input.Trigger) {
        address := strings.SplitAfterN(ev.Text, input.Trigger,1)
        log.Printf("Scanning address: %s", address)
        //Call function portscan
			}

		}
	}
}
