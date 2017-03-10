package main

//Split these to two ironFunctions, one to fetch next meetup, and one to communicate with slack? maybe 3, one to listen to slacka nd one to send message

import (
	"fmt"
	"log"

	"github.com/nlopes/slack"
	"encoding/json"
	"os"
)

type Input struct {
	Token     string
	Channel   string
	Message   string
}

func getStdIn() (s Input, err error) {
	s = Input{}

	err = json.NewDecoder(os.Stdin).Decode(&s)

	if (err != nil){
		return
	}

	if(s.Token=="") {
		err = fmt.Errorf("Please define Slack Token")
	} else if(s.Channel == "") {
		err = fmt.Errorf("Please define Slack Channel")
	} else if(s.Message == "") {
		err = fmt.Errorf("Please define Slack Message")
	}

	return
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

	rtm.SendMessage(rtm.NewOutgoingMessage(input.Message, input.Channel))

}