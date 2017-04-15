// This will be the command line app that uses the tablature/instrument package api.
// It is a WIP and is not yet intended for use.
package main

import (
	"bufio"
	"fmt"
	"github.com/Guitarbum722/tablature/instrument"
	"github.com/Guitarbum722/tablature/tabio"
	"log"
	"os"
	"strings"
)

func main() {
	var player instrument.Instrument = instrument.NewInstrument("guitar")
	// err := player.Tune("Dadgbe")
	// if err != nil {
	// 	log.Fatal("error tuning guitar")
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the string, then fret number (ie e7 or g2)")
	fmt.Println("****************")
	fmt.Print(instrument.StringifyCurrentTab(player))

	w := tabio.NewTablatureWriter(os.Stdout, 20)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1) // Windows; BOOO!

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
}
