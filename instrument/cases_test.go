package instrument

var newInstrumentCases = []struct {
	inputType    string
	expectedType Instrument
}{
	{
		"guitar",
		&Guitar{fretBoard: Fretboard{
			'E': "---",
			'a': "---",
			'd': "---",
			'g': "---",
			'b': "---",
			'e': "---"},
			order:        TuningOrder{'E', 'a', 'd', 'g', 'b', 'e'},
			numOfStrings: 6},
	},
	{
		"bass",
		&Bass{fretBoard: Fretboard{
			'E': "---",
			'a': "---",
			'd': "---",
			'g': "---"},
			order:        TuningOrder{'E', 'a', 'd', 'g'},
			numOfStrings: 4},
	},
	{
		"unKnoWN",
		&Guitar{fretBoard: Fretboard{
			'E': "---",
			'a': "---",
			'd': "---",
			'g': "---",
			'b': "---",
			'e': "---"},
			order:        TuningOrder{'E', 'a', 'd', 'g', 'b', 'e'},
			numOfStrings: 6},
	},
}
