package main

import (
	"fmt"
	"github.com/gusanmaz/capdec"
	"log"
)

func main() {
	dest := "bird_captioned.png"
	captions := make([]string, 2)
	captions[0] = "Blue-footed bird."
	captions[1] = "What a fancy bird!"
	codes := make([]string, 2)
	codes[0] = `document.querySelector('figure').style.borderWidth = "10px"`
	codes[1] = `document.querySelector('figure').style.borderColor = "blue" ;`

	err := capdec.Caption("bird.jpeg", captions, dest, codes)
	if err != nil {
		log.Fatal("Terminating because of an error. Error: " + err.Error())
	} else {
		fmt.Printf("Captioned image successfully saved at %v\n", dest)
	}
}
