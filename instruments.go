package main

// Instrument defines behavior that is common across the instruments used in this package, such as setting the tuning.
type Instrument interface {
	Tune(tuning string)
}

// Guitar represents a standard 6 string guitar with a default tuning of Eadgbe.  Tuning can be changed by calling Tuning()
type Guitar struct {
	fretBoard map[byte]string
}

// NewInstrument returns the configured instrument in its default tuning.
func NewInstrument(i string) Instrument {
	var instrument Instrument

	switch i {
	case "guitar":
		instrument = &Guitar{fretBoard: map[byte]string{
			'E': "---",
			'a': "---",
			'd': "---",
			'g': "---",
			'b': "---",
			'e': "---"},
		}
	default:
		instrument = &Guitar{fretBoard: map[byte]string{
			'E': "---",
			'a': "---",
			'd': "---",
			'g': "---",
			'b': "---",
			'e': "---"},
		}
	}

	return instrument
}

// Tune updates the tuning configuration of the current Guitar.
func (g *Guitar) Tune(tuning string) {

}
