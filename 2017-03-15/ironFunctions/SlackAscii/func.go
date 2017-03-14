package main

import (
	"encoding/json"
	"os"
	"github.com/common-nighthawk/go-figure"
	"fmt"



	"github.com/nlopes/slack"

)

type Avatar struct {
	Username   string
	Gender     string

}

type Input struct {
	Token     string
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

	input := &Input{Token: ""}
	json.NewDecoder(os.Stdin).Decode(input)


	api := slack.New(input.Token)


	rtm := api.NewRTM()
	go rtm.ManageConnection()


	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:

			rows := textToAscii(ev.Text, "larry3d")
			for _, printRow := range rows {
				message:=fmt.Sprintf("`%s`",printRow)
				rtm.SendMessage(rtm.NewOutgoingMessage(message, ev.Channel))
			}
		}
	}
}

func textToAscii(text string, font string) []string {

		//Draw ascii
		return figure.NewFigure(text, font,false).Slicify()
}
