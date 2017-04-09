package instrument

import (
	"fmt"
)

const testVersion = 1

// Instrument defines behavior that is common across the instruments used in this package, such as setting the tuning.
type Instrument interface {
	Tune(tuning string)
	Fretboard() Fretboard
	Order() TuningOrder
}

// TuningOrder defines the order in which the strings would be physically on the guitar.  TuningOrder can be used to display the strings properly and consistently.
type TuningOrder []byte

// Fretboard represents the string/fretnumber relationship that is a single note or an entire chord.
type Fretboard map[byte]string

// Guitar represents a standard 6 string guitar with a default tuning of Eadgbe.  Tuning can be changed by calling Tune()
type Guitar struct {
	fretBoard    Fretboard
	order        TuningOrder
	numOfStrings int
}

// Bass represents a standard 4 string bass guitar with a default tuning of Eadg.  Tuning can be changed by calling Tune()
type Bass struct {
	fretBoard    Fretboard
	order        TuningOrder
	numOfStrings int
}

// NewInstrument returns the configured instrument in its default tuning.  A Guitar in standard tuning will be returned by default if the desired instrument is not available.
func NewInstrument(i string) Instrument {
	var instrument Instrument

	switch i {
	case "guitar":
		instrument = newGuitar()
	case "bass":
		instrument = newBass()
	default:
		instrument = newGuitar()
	}

	return instrument
}

// Tune updates the tuning configuration of the current Guitar.  The order of the strings will also be the order in which the tuning was input to Tune()
func (g *Guitar) Tune(tuning string) {

}

// Fretboard returns the current map which represents the Guitar pointers tuning.
func (g *Guitar) Fretboard() Fretboard {
	return g.fretBoard
}

func (g *Guitar) orderTuning() {

}

// Order returns the current tuning order for the current Guitar
func (g *Guitar) Order() TuningOrder {
	return g.order
}

// returns a pointer to a guitar with standard tuning by default.
func newGuitar() *Guitar {
	return &Guitar{fretBoard: Fretboard{
		'E': "---",
		'a': "---",
		'd': "---",
		'g': "---",
		'b': "---",
		'e': "---"},
		order:        TuningOrder{'E', 'a', 'd', 'g', 'b', 'e'},
		numOfStrings: 6,
	}
}

// Tune updates the tuning configuration of the current Bass.  The order of the strings will also be the order in which the tuning was input to Tune()
func (b *Bass) Tune(tuning string) {

}

// Order returns the current tuning order for the current Bass
func (b *Bass) Order() TuningOrder {
	return b.order
}

// Fretboard returns the current map which represents the Bass pointers tuning.
func (b *Bass) Fretboard() Fretboard {
	return b.fretBoard
}

func (b *Bass) orderTuning() {

}

// returns a pointer to a guitar with a standard tuning by default.
func newBass() *Bass {
	return &Bass{fretBoard: Fretboard{
		'E': "---",
		'a': "---",
		'd': "---",
		'g': "---"},
		order:        TuningOrder{'E', 'a', 'd', 'g'},
		numOfStrings: 4,
	}
}

// StringifyCurrentTab converts the current fretBoard configuration to a string.
func StringifyCurrentTab(i Instrument) string {
	orderOfStrings := i.Order()
	fretBoard := i.Fretboard()
	var result string
	for _, v := range orderOfStrings {
		result += fmt.Sprintf("%s : %s\n", string(v), fretBoard[v])
	}
	return result
}

// UpdateCurrentTab accepts the guitarString and fret to be updated on the Instrument.
func UpdateCurrentTab(i Instrument, guitarString byte, fret string) {
	i.Fretboard()[guitarString] = fmt.Sprintf("-%s-", fret)
}
