// This will be the command line app that uses the tablature/instrument package api.
// It is a WIP and is not yet intended for use.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Guitarbum722/go-tabs/instrument"
	"github.com/Guitarbum722/go-tabs/tabio"
)

const (
	guitar      = "guitar"
	bass        = "bass"
	ukulele     = "ukulele"
	guitarSeven = "guitar-seven"
	mandolin    = "mandolin"
	bassFive    = "bass-five"
)

// TODO(guitarbum722) define more valid characters such as hammer-on/pull-off and dead string 2017-04-15T18:00 4
func main() {
	// var player = instrument.NewInstrument("mandolin")
	// err := player.Tune("Dadgbe")
	// if err != nil {
	// 	log.Fatal("error tuning guitar")
	// }

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Which instrument?")
	fmt.Println("1 - Guitar")
	fmt.Println("2 - Bass Guitar")
	fmt.Println("3 - Ukulele")
	fmt.Println("4 - Seven String Guitar")
	fmt.Println("5 - Mandolin")
	fmt.Println("6 - Five String Bass Guitar")

	userChoice, err := reader.ReadString('\n')
	userChoice = strings.Replace(userChoice, "\n", "", -1)

	if err != nil {
		log.Fatalf("error with input %v", err)
	}

	var player instrument.Instrument

	switch userChoice {
	case "1":
		player = instrument.NewInstrument(guitar)
	case "2":
		player = instrument.NewInstrument(bass)
	case "3":
		player = instrument.NewInstrument(ukulele)
	case "4":
		player = instrument.NewInstrument(guitarSeven)
	case "5":
		player = instrument.NewInstrument(mandolin)
	case "6":
		player = instrument.NewInstrument(bassFive)
	default:
		log.Printf("Your choice of %v is not valid.\n", userChoice)
		os.Exit(1)
	}

	fmt.Println("Enter the string, then fret number (ie e7 or g2)")
	fmt.Println("****************")
	fmt.Print(instrument.StringifyCurrentTab(player))

	// w := tabio.NewTablatureWriter(os.Stdout, 20)
	f, err := os.OpenFile("guitar_tab.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Printf("error opening file %s", err)
	}
	w := tabio.NewTablatureWriter(f, 80)

	for z := 0; z < 1; {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		switch input {
		case "stage":
			tabio.StageTablature(player, w)
			for k := range player.Fretboard() {
				instrument.UpdateCurrentTab(player, k, "-")
			}
			fmt.Print(instrument.StringifyCurrentTab(player))
		case "export":
			if err := tabio.ExportTablature(player, w); err != nil {
				log.Fatalf("there was an error exporting the tablature::: %s\n", err)
			}
			for k := range player.Fretboard() {
				instrument.UpdateCurrentTab(player, k, "-")
			}
		case "quit":
			z++
		default:
			guitarString, fret, err := instrument.ParseFingerBoard(input)
			if err != nil {
				log.Printf("invalid entry: %s", err)
			} else {
				instrument.UpdateCurrentTab(player, guitarString, fret)
			}
			fmt.Print(instrument.StringifyCurrentTab(player))
		}
	}
	if err := f.Close(); err != nil {
		log.Printf("error closing the file %s", err)
		os.Exit(1)
	}
	os.Exit(0)
}
