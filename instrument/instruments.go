package instrument

import (
	"errors"
	"fmt"
)

const testVersion = 1

// Instrument defines behavior that is common across the instruments used in this package, such as setting the tuning.
type Instrument interface {
	Tune(tuning string) error
	Fretboard() Fretboard
	Order() TuningOrder
	NumOfStrings() int
}

// TuningOrder defines the order in which the strings would be physically on the guitar.  TuningOrder can be used to display the
// strings properly and consistently.
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

// NewInstrument returns the configured instrument in its default tuning.  A Guitar in standard tuning will be returned by
// default if the desired instrument is not available.
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

// Tune updates the tuning configuration of the current Guitar.  The order of the strings will also be the order in which the
// tuning was input to Tune().  Validation of the input will occur to make sure that the number of instrument strings is the same
// as the count of the input.
func (g *Guitar) Tune(tuning string) error {
	for _, v := range tuning {
		if valid := validMusicNote(v); !valid {
			return errors.New("one or more of the note provided in the requested tuning is invalid")
		}
	}
	if ok := validCount(g, tuning); !ok {
		errMsg := fmt.Sprintf("attempted to reconfigure the Guitar with %d strings which does not match the allowed %d number of strings",
			len(tuning),
			g.NumOfStrings())
		return errors.New(errMsg)
	}
	for k := range g.fretBoard {
		delete(g.fretBoard, k)
	}
	for _, v := range tuning {
		g.fretBoard[byte(v)] = "---"
	}
	g.order = TuningOrder(tuning)
	return nil
}

// Fretboard returns the current map which represents the Guitar pointers tuning.
func (g *Guitar) Fretboard() Fretboard {
	return g.fretBoard
}

// NumOfStrings returns the number of strings that the instrument has.
func (g *Guitar) NumOfStrings() int {
	return g.numOfStrings
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

// Tune updates the tuning configuration of the current Bass.  The order of the strings will also be the order in which the
// tuning was input to Tune().  Validation of the input will occur to make sure that the number of instrument strings is the same
// as the count of the input.
func (b *Bass) Tune(tuning string) error {
	for _, v := range tuning {
		if ok := validMusicNote(v); !ok {
			return errors.New("one or more of the note provided in the requested tuning is invalid")
		}
	}
	if ok := validCount(b, tuning); !ok {
		errMsg := fmt.Sprintf("attempted to reconfigure the Bass with %d strings which does not match the allowed %d number of strings",
			len(tuning),
			b.NumOfStrings())
		return errors.New(errMsg)
	}
	return nil
}

// Order returns the current tuning order for the current Bass
func (b *Bass) Order() TuningOrder {
	return b.order
}

// Fretboard returns the current map which represents the Bass pointers tuning.
func (b *Bass) Fretboard() Fretboard {
	return b.fretBoard
}

// NumOfStrings returns the number of strings that the instrument has.
func (b *Bass) NumOfStrings() int {
	return b.numOfStrings
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
func UpdateCurrentTab(i Instrument, instrumentString byte, fret string) {
	i.Fretboard()[instrumentString] = fmt.Sprintf("-%s-", fret)
}

func validMusicNote(note rune) bool {
	return 'a' <= note && note <= 'g' || 'A' <= note && note <= 'G'
}

func validCount(i Instrument, s string) bool {
	return i.NumOfStrings() == len(s)
}
