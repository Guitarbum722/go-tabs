package instrument

// Instrument defines behavior that is common across the instruments used in this package, such as setting the tuning.
type Instrument interface {
	Tune(tuning string)
	Fretboard() map[byte]string
	orderTuning()
}

// Guitar represents a standard 6 string guitar with a default tuning of Eadgbe.  Tuning can be changed by calling Tune()
type Guitar struct {
	fretBoard map[byte]string
}

// Bass represents a standard 4 string bass guitar with a default tuning of Eadg.  Tuning can be changed by calling Tune()
type Bass struct {
	fretBoard map[byte]string
}

// NewInstrument returns the configured instrument in its default tuning.
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

// Tune updates the tuning configuration of the current Guitar.
func (g *Guitar) Tune(tuning string) {

}

// Fretboard returns the current map which represents the Guitar pointers tuning.
func (g *Guitar) Fretboard() map[byte]string {
	return g.fretBoard
}

func (g *Guitar) orderTuning() {

}

// returns a pointer to a guitar with standard tuning by default.
func newGuitar() *Guitar {
	return &Guitar{fretBoard: map[byte]string{
		'E': "---",
		'a': "---",
		'd': "---",
		'g': "---",
		'b': "---",
		'e': "---"},
	}
}

// Tune updates the tuning configuration of the current Bass.
func (b *Bass) Tune(tuning string) {

}

// Fretboard returns the current map which represents the Bass pointers tuning.
func (b *Bass) Fretboard() map[byte]string {
	return b.fretBoard
}

func (b *Bass) orderTuning() {

}

// returns a pointer to a guitar with a standard tuning by default.
func newBass() *Bass {
	return &Bass{fretBoard: map[byte]string{
		'E': "---",
		'a': "---",
		'd': "---",
		'g': "---"},
	}
}
