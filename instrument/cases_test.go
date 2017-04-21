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
			order:        TuningOrder{'e', 'b', 'g', 'd', 'a', 'E'},
			numOfStrings: 6},
	},
	{
		"bass",
		&Bass{fretBoard: Fretboard{
			'E': "---",
			'a': "---",
			'd': "---",
			'g': "---"},
			order:        TuningOrder{'g', 'd', 'a', 'E'},
			numOfStrings: 4},
	},
	{
		"ukulele",
		&Ukulele{fretBoard: Fretboard{
			'G': "---",
			'c': "---",
			'e': "---",
			'a': "---",
		},
			order:        TuningOrder{'a', 'e', 'c', 'G'},
			numOfStrings: 4},
	},
	{
		"guitar-seven",
		&GuitarSeven{fretBoard: Fretboard{
			'B': "---",
			'E': "---",
			'a': "---",
			'd': "---",
			'g': "---",
			'b': "---",
			'e': "---"},
			order:        TuningOrder{'e', 'b', 'g', 'd', 'a', 'E', 'B'},
			numOfStrings: 7},
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
			order:        TuningOrder{'e', 'b', 'g', 'd', 'a', 'E'},
			numOfStrings: 6},
	},
}

var tuneCases = []struct {
	input     string
	newTuning string
	wantErr   bool
}{
	{
		"guitar",
		"Dadgbe",
		false,
	},
	{
		"bass",
		"Dadgb",
		true,
	},
	{
		"guitar",
		"Eadgbz",
		true,
	},
}

var parseFingerBoardCases = []struct {
	input            string
	instrumentString byte
	fret             string
	wantErr          bool
}{
	{
		"a1",
		'a',
		"1",
		false,
	},
	{
		"E12",
		'E',
		"12",
		false,
	},
	{
		"z10",
		0,
		"-",
		true,
	},
	{
		"D2222",
		0,
		"-",
		true,
	},
}
