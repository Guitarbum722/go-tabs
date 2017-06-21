package instrument

// Represents all valid music notes of the chromatic scale
const (
	LowA       = "A"
	LowASharp  = "A♯"
	LowBFlat   = "B♭"
	LowB       = "B"
	LowC       = "C"
	LowCSharp  = "C♯"
	LowDFlat   = "D♭"
	LowD       = "D"
	LowDSharp  = "D♯"
	LowEFlat   = "E♭"
	LowE       = "E"
	LowF       = "F"
	LowFSharp  = "F♯"
	LowGFlat   = "G♭"
	LowG       = "G"
	LowGSharp  = "G♯"
	LowAFlat   = "A♭"
	HighA      = "a"
	HighASharp = "a♯"
	HighBFlat  = "b♭"
	HighB      = "b"
	HighC      = "c"
	HighCSharp = "c♯"
	HighDflat  = "d♭"
	HighD      = "d"
	HighDSharp = "d♯"
	HighEFlat  = "e♭"
	HighE      = "e"
	HighF      = "f"
	HighFSharp = "f♯"
	HighGFlat  = "g♭"
	HighG      = "g"
	HighGSharp = "g♯"
	HighAFlat  = "a♭"
)

func validMusicNote(note string) bool {
	switch note {
	case "A":
		return true
	case "A♯":
		return true
	case "B♭":
		return true
	case "B":
		return true
	case "C":
		return true
	case "C♯":
		return true
	case "D♭":
		return true
	case "D":
		return true
	case "D♯":
		return true
	case "E♭":
		return true
	case "E":
		return true
	case "F":
		return true
	case "F♯":
		return true
	case "G♭":
		return true
	case "G":
		return true
	case "G♯":
		return true
	case "A♭":
		return true
	case "a":
		return true
	case "a♯":
		return true
	case "b♭♯":
		return true
	case "b":
		return true
	case "c":
		return true
	case "c♯":
		return true
	case "d♭":
		return true
	case "d":
		return true
	case "d♯":
		return true
	case "e♭":
		return true
	case "e":
		return true
	case "f":
		return true
	case "f♯":
		return true
	case "g♭":
		return true
	case "g":
		return true
	case "g♯":
		return true
	case "a♭":
		return true
	default:
		return false
	}
}
