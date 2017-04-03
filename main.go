package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the string, then fret number (ie e7 or g2)")
	fmt.Println("****************")

	// TODO(guitarbum722) create default current tablature for multiple instruments 2017-04-02T18:54 4
	var current = []string{"---", "---", "---", "---", "---", "---"}

	fmt.Println(strings.Join(current, "\n"))
	for {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1) // Windows; BOOO!

		// Guitar string number plus frets should not be less than 2 and not more than 3 characters in length.  This allots for up to 99 frets.
		if len(input) < 2 || len(input) > 3 {
			fmt.Print("Invalid entry; make sure the format is [string#][fret#]")
		}

		// TODO(guitarbum722) parse input and reconstruct the values of current for display 2017-04-02T18:54 4
		fmt.Println(strings.Join(current, "\n"))
	}
}
