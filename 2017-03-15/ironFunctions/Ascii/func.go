package main

import (
	"encoding/json"
	"github.com/common-nighthawk/go-figure"
	"os"
)

type Ascii struct {
	Text string
	Font string
}

func main() {

	//Set default values
	ascii := &Ascii{Text: "Hello World!", Font: "smkeyboard"}

	//Decode JSON input
	json.NewDecoder(os.Stdin).Decode(ascii)

	//Draw ascii
	figure.NewFigure(ascii.Text, ascii.Font, false).Print()
}
