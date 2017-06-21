package instrument

var newInstrumentCases = []struct {
	inputType    string
	expectedType Instrument
}{
	{
		"guitar",
		&Guitar{fretBoard: Fretboard{
			"E": openPlayerString,
			"a": openPlayerString,
			"d": openPlayerString,
			"g": openPlayerString,
			"b": openPlayerString,
			"e": openPlayerString,
		},
			order:        TuningOrder{"e", "b", "g", "d", "a", "E"},
			numOfStrings: 6},
	},
	{
		"bass",
		&Bass{fretBoard: Fretboard{
			"E": openPlayerString,
			"a": openPlayerString,
			"d": openPlayerString,
			"g": openPlayerString,
		},
			order:        TuningOrder{"g", "d", "a", "E"},
			numOfStrings: 4},
	},
	{
		"bass-five",
		&BassFive{fretBoard: Fretboard{
			"B": openPlayerString,
			"E": openPlayerString,
			"a": openPlayerString,
			"d": openPlayerString,
			"g": openPlayerString,
		},
			order:        TuningOrder{"g", "d", "a", "E", "B"},
			numOfStrings: 5},
	},
	{
		"ukulele",
		&Ukulele{fretBoard: Fretboard{
			"G": openPlayerString,
			"c": openPlayerString,
			"e": openPlayerString,
			"a": openPlayerString,
		},
			order:        TuningOrder{"a", "e", "c", "G"},
			numOfStrings: 4},
	},
	{
		"guitar-seven",
		&GuitarSeven{fretBoard: Fretboard{
			"B": openPlayerString,
			"E": openPlayerString,
			"a": openPlayerString,
			"d": openPlayerString,
			"g": openPlayerString,
			"b": openPlayerString,
			"e": openPlayerString,
		},
			order:        TuningOrder{"e", "b", "g", "d", "a", "E", "B"},
			numOfStrings: 7},
	},
	{
		"mandolin",
		&Mandolin{fretBoard: Fretboard{
			"G": openPlayerString,
			"d": openPlayerString,
			"a": openPlayerString,
			"e": openPlayerString,
		},
			order:        TuningOrder{"e", "a", "d", "G"},
			numOfStrings: 4},
	},
	{
		"lap-steel",
		&LapSteel{fretBoard: Fretboard{
			"C": openPlayerString,
			"E": openPlayerString,
			"g": openPlayerString,
			"a": openPlayerString,
			"c": openPlayerString,
			"e": openPlayerString,
		},
			order:        TuningOrder{"e", "c", "a", "g", "E", "C"},
			numOfStrings: 6,
		},
	},
	{
		"unKnoWN",
		&Guitar{fretBoard: Fretboard{
			"E": openPlayerString,
			"a": openPlayerString,
			"d": openPlayerString,
			"g": openPlayerString,
			"b": openPlayerString,
			"e": openPlayerString,
		},
			order:        TuningOrder{"e", "b", "g", "d", "a", "E"},
			numOfStrings: 6},
	},
}

var tuneCases = []struct {
	input     string
	newTuning []string
	wantErr   bool
}{
	{
		"guitar",
		[]string{"D", "a", "d", "g", "b", "e"},
		false,
	},
	{
		"bass",
		[]string{"D", "a", "d", "g", "b"},
		true,
	},
	{
		"guitar",
		[]string{"E", "a", "d", "g", "b", "z"},
		true,
	},
	{
		"guitar",
		[]string{"D", "a", "d", "g", "b", "e", "A"},
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
		"a:0",
		'a',
		"0",
		false,
	},
	{
		"a:1",
		'a',
		"1",
		false,
	},
	{
		"E:12",
		'E',
		"12",
		false,
	},
	{
		"z:10",
		0,
		"-",
		true,
	},
	{
		"D:2222",
		0,
		"-",
		true,
	},
}
