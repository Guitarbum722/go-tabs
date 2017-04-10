package main

import (
	"bufio"
	"fmt"
	"github.com/Guitarbum722/tablature/instrument"
	"os"
	"strings"
)

func main() {
	player := instrument.NewInstrument("bass")
	err := player.Tune("Dadgbe")
	fmt.Println(err)
	fmt.Println(player.Order())

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the string, then fret number (ie e7 or g2)")
	fmt.Println("****************")

	fmt.Print(instrument.StringifyCurrentTab(player))
	for {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1) // Windows; BOOO!

		guitarString, fret := parseFingerBoard(input)

		instrument.UpdateCurrentTab(player, guitarString, fret)

		fmt.Print(instrument.StringifyCurrentTab(player))

	}
}

func parseFingerBoard(i string) (byte, string) {

	var guitarString byte
	var fret string

	// Guitar string number plus frets should not be less than 2 and not more than 3 characters in length.  This allots for up to 99 frets.
	if len(i) < 2 || len(i) > 3 {
		fmt.Print("Invalid entry; make sure the format is [string#][fret#]")
	}

	// confirms that the string input by the user is valid
	switch i[0] {
	case 'E':
		guitarString = i[0]
	case 'a':
		guitarString = i[0]
	case 'd':
		guitarString = i[0]
	case 'g':
		guitarString = i[0]
	case 'b':
		guitarString = i[0]
	case 'e':
		guitarString = i[0]
	}

	fret = i[1:]

	return guitarString, fret
}
