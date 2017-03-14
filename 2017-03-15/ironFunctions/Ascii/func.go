package main

import (
	"encoding/json"
	"os"

	"github.com/common-nighthawk/go-figure"

)

type Ascii struct {
	Text   string
	Font   string

}

func main() {

	//Set default values
	ascii := &Ascii{Text: "Hello World!", Font: "smkeyboard"}

	//Decode JSON input
	json.NewDecoder(os.Stdin).Decode(ascii)

	//Draw ascii
	figure.NewFigure(ascii.Text, ascii.Font, false).Scroll(20000,200, "right")
}
