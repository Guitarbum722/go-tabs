package instrument

import (
	"errors"
	"fmt"
	"reflect"
)

const testVersion = 2

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

// Ukulele represents a standard 4 string ukulele with a default tuning of Gcea.  Tuning can be changed by calling Tune()
type Ukulele struct {
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
	case "ukulele":
		instrument = newUkulele()
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
		errMsg := tuningLengthError(g, tuning)
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

// Fretboard returns the current map which represents the Guitar pointer's tuning.
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
		order:        TuningOrder{'e', 'b', 'g', 'd', 'a', 'E'},
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
		errMsg := tuningLengthError(b, tuning)
		return errors.New(errMsg)
	}
	return nil
}

// Order returns the current tuning order for the current Bass
func (b *Bass) Order() TuningOrder {
	return b.order
}

// Fretboard returns the current map which represents the Bass pointer's tuning.
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
		order:        TuningOrder{'g', 'd', 'a', 'E'},
		numOfStrings: 4,
	}
}

func newUkulele() *Ukulele {
	return &Ukulele{fretBoard: Fretboard{
		'G': "---",
		'c': "---",
		'e': "---",
		'a': "---"},
		order:        TuningOrder{'a', 'e', 'c', 'G'},
		numOfStrings: 4,
	}
}

// Tune updates the tuning configuration of the current Ukulele.  The order of the strings will also be the order in which the
// tuning was input to Tune().  Validation of the input will occur to make sure that the number of instrument strings is the same
// as the count of the input.
func (u *Ukulele) Tune(tuning string) error {
	for _, v := range tuning {
		if ok := validMusicNote(v); !ok {
			return errors.New("one or more of the note provided in the requested tuning is invalid")
		}
	}
	if ok := validCount(u, tuning); !ok {
		errMsg := tuningLengthError(u, tuning)
		return errors.New(errMsg)
	}
	return nil
}

// Fretboard returns the current map which represents the Ululele pointer's tuning.
func (u *Ukulele) Fretboard() Fretboard {
	return u.fretBoard
}

// NumOfStrings returns the number of strings that the instrument has.
func (u *Ukulele) NumOfStrings() int {
	return u.numOfStrings
}

// Order returns the current tuning order for the current Ukulele
func (u *Ukulele) Order() TuningOrder {
	return u.order
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
	switch len(fret) {
	case 1:
		if fret == "0" {
			i.Fretboard()[instrumentString] = fmt.Sprintf("--%s", "-")
		} else {
			i.Fretboard()[instrumentString] = fmt.Sprintf("--%s", fret)
		}
	case 2:
		i.Fretboard()[instrumentString] = fmt.Sprintf("-%s", fret)
	case 3:
		i.Fretboard()[instrumentString] = fmt.Sprintf("%s", fret)
	}
}

func validMusicNote(note rune) bool {
	return 'a' <= note && note <= 'g' || 'A' <= note && note <= 'G'
}

func validCount(i Instrument, s string) bool {
	return i.NumOfStrings() == len(s)
}

// validates that the fret number is numeric
func validFretCount(s string) bool {
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

func tuningLengthError(i Instrument, tuning string) string {
	return fmt.Sprintf("attempted to reconfigure the %s with %d strings which does not match the allowed %d number of strings",
		reflect.TypeOf(i),
		len(tuning),
		i.NumOfStrings())
}

// ParseFingerBoard validates input for the next tab fingering, validates it and returns the parsed values.
func ParseFingerBoard(i string) (byte, string, error) {

	var instrumentString byte

	// Guitar string number plus frets should not be less than 2 and not more than 3 characters in length.  This allots for up to 99 frets.
	if len(i) < 2 || len(i) > 4 {
		return 0, "-", errors.New("invalid entry: make sure the format is [string#][fret#] and the fret number is <= 999")
	}

	if !validMusicNote(rune(i[0])) {
		return 0, "-", errors.New("invalid entry: make sure the string is a valid music note")
	}

	if !validFretCount(i[1:]) {
		return 0, "-", errors.New("invalid entry: make sure the fret number is numeric")
	}

	// confirms that the string input by the user is valid
	switch i[0] {
	case 'A':
		instrumentString = i[0]
	case 'B':
		instrumentString = i[0]
	case 'C':
		instrumentString = i[0]
	case 'D':
		instrumentString = i[0]
	case 'E':
		instrumentString = i[0]
	case 'F':
		instrumentString = i[0]
	case 'G':
		instrumentString = i[0]
	case 'a':
		instrumentString = i[0]
	case 'b':
		instrumentString = i[0]
	case 'c':
		instrumentString = i[0]
	case 'd':
		instrumentString = i[0]
	case 'e':
		instrumentString = i[0]
	case 'f':
		instrumentString = i[0]
	case 'g':
		instrumentString = i[0]
	}

	fret := i[1:]

	return instrumentString, fret, nil
}
