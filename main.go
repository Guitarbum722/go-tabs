// This will be the command line app that uses the tablature/instrument package api.
// It is a WIP and is not yet intended for use.
package main

import (
	"bufio"
	"fmt"
	"github.com/Guitarbum722/tablature/instrument"
	"log"
	"os"
	"strings"
)

func main() {
	var player instrument.Instrument = instrument.NewInstrument("guitar")
	err := player.Tune("Dadgbe")
	if err != nil {
		log.Fatal("error tuning guitar")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the string, then fret number (ie e7 or g2)")
	fmt.Println("****************")

	fmt.Print(instrument.StringifyCurrentTab(player))
	for {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1) // Windows; BOOO!

		guitarString, fret, err := instrument.ParseFingerBoard(input)
		if err != nil {
			log.Printf("invalid entry: %s", err)
		} else {
			instrument.UpdateCurrentTab(player, guitarString, fret)
		}

		fmt.Print(instrument.StringifyCurrentTab(player))

	}
}
