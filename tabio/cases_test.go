package tabio

var newTablatureWriterCases = []struct {
	inWrapPos   int
	expectedPos int
}{
	{
		20,
		20,
	},
	{
		18,
		20,
	},
	{
		110,
		110,
	},
}
