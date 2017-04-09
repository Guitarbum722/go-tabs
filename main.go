// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"github.com/Guitarbum722/tablature/instrument"
// 	"os"
// 	"strings"
// )

// // TODO(guitarbum722) define structure that will display the instrument map in order 2017-04-06T18:00 5

// func main() {
// 	instrument := instrument.NewInstrument("guitar")

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter the string, then fret number (ie e7 or g2)")
// 	fmt.Println("****************")

// 	orderOfStrings := instrument.Order()
// 	fretBoard := instrument.Fretboard()
// 	for _, v := range orderOfStrings {
// 		fmt.Printf("%s : %s\n", string(v), fretBoard[v])
// 	}

// 	for {
// 		input, _ := reader.ReadString('\n')
// 		input = strings.Replace(input, "\n", "", -1) // Windows; BOOO!

// 		guitarString, fret := parseFingerBoard(input)

// 		fretBoard[guitarString] = fmt.Sprintf("-%s-", fret)

// 		// TODO(guitarbum722) parse input and reconstruct the values of currentTab for display 2017-04-02T18:54 4
// 		for _, v := range orderOfStrings {
// 			fmt.Printf("%s : %s\n", string(v), fretBoard[v])
// 		}
// 	}
// }

// func parseFingerBoard(i string) (byte, string) {

// 	var guitarString byte
// 	var fret string

// 	// Guitar string number plus frets should not be less than 2 and not more than 3 characters in length.  This allots for up to 99 frets.
// 	if len(i) < 2 || len(i) > 3 {
// 		fmt.Print("Invalid entry; make sure the format is [string#][fret#]")
// 	}

// 	// confirms that the string input by the user is valid
// 	switch i[0] {
// 	case 'E':
// 		guitarString = i[0]
// 	case 'a':
// 		guitarString = i[0]
// 	case 'd':
// 		guitarString = i[0]
// 	case 'g':
// 		guitarString = i[0]
// 	case 'b':
// 		guitarString = i[0]
// 	case 'e':
// 		guitarString = i[0]
// 	}

// 	fret = i[1:]

// 	return guitarString, fret
// }
