### ----GO--------
### -🎸-----TABS---

[![Build Status](https://travis-ci.org/Guitarbum722/go-tabs.svg?branch=master)](https://travis-ci.org/Guitarbum722/go-tabs)

### A minimal and easy to use package for creating tablature for stringed / fretted instruments

#### What's included?
* A very ugly, yet easy to use command line interface
* COMING SOON! - A (nearly) just as ugly web interface
* A simple package that can be quickly and easily implemented by your own application

#### Usage

##### _*In a nutshell, the basic usage is as follows:*_
> 1. Initialize an _Instrument_
> 2. Initialize a _TablatureWriter_
> 3. Tune the _instrument_ (only if desired tuning is different than the default)
> 4. Provide input for the 'current' tablature.  This is how you input a single note on the fretboard
or build an entire chord.
> 5. 'Stage' the current tablature.
> 6. Keep building 'current' tablature and staging
> 7. Export the staged tablature to the provided writer (typically a file) 

``` sh
$ go get -u github.com/Guitarbum722/go-tabs
```

Initialize an instrument
```go
	var player = instrument.NewInstrument("guitar")
```
or...
```go
	var player = instrument.NewInstrument("bass")
```

Also initialize a TablatureWriter so that the tablature which is WIP, or "staged" can be properly buffered,
organized, and prepared to be exported to the underlying writer.  In this case, it is Stdout.
But a more common use case would be a file.
```go
    w := tabio.NewTablatureWriter(os.Stdout, 100) // provide io.Writer and a wrap position
```
*The wrap position is important because it determines where each section of the tablature will wrap to the next.*
_Example: a wrap position of 10_
```
e: --2-----------------
b: -----222222222222222
g: --------------------
d: --------------------
a: --------------------
E: --------------------

e: --2-
b: ----
g: ----
d: ----
a: ----
E: ----
```

_a wrap position of 100_
```
e: --5-----------------------------------------------------------------------------
b: --------------------------------------------------------2-----------------------
g: -----------------------------------------------------------------1--------------
d: -----------------------------------------------------------2--1--2--------------
a: --2--------------------------------------------------------------2--------------
E: --5-----------------------------------------------------------------------------

e: -------------------------------
b: -------------------------------
g: -------------------------------
d: -------------------------------
a: -------------------------------
E: -------------------------------
```

```go
// update the wrap
w.UpdateWrapPosition(110) // if input is less than 20, the default is 20
```

Enter an instrument's string and the fret number.  Example [ E12 ] or [ a7 ]
```go
	guitarString, fret, err := instrument.ParseFingerBoard(input) // (input == "E12") && true == true
	if err != nil {
		log.Printf("invalid entry: %s", err)
	} else {
		instrument.UpdateCurrentTab(player, guitarString, fret) // update the instrument's 'current' tablature

	}
	fmt.Print(instrument.StringifyCurrentTab(player))

```
Above: validate and parse the input to refresh the 'current' tablature (which could be a single note or a chord).

If it is determined that the current tablature is correct, then it needs to get 'staged' which is really just a map that
buffers each of the instrument's tablature by string.

```go
	tabio.StageTablature(player, w) // adds the current tablature to the staging buffer of the TablatureWriter
	for k := range player.Fretboard() {
		instrument.UpdateCurrentTab(player, k, "-") // refreshes the current tablature with no fret markers
	}
	fmt.Print(instrument.StringifyCurrentTab(player))

```

Now export all of the staged tablature to the TablatureWriter's underlying buffered writer that it was
initialized with (in the case of this document, os.Stdout).  The tablature will wrap appropriately based
on the configured wrap position.
```go
	if err := tabio.ExportTablature(player, w); err != nil {
		log.Fatalf("there was an error exporting the tablature::: %s\n", err)
	}
```
```
e: --5-----------------------------------------------------------------------------
b: --------------------------------------------------------2-----------------------
g: -----------------------------------------------------------------1--------------
d: -----------------------------------------------------------2--1--2--------------
a: --2--------------------------------------------------------------2--------------
E: --5-----------------------------------------------------------------------------
```

### What if I want a different tuning?

You can tune the instrument to whatever you want!  If two strings are tuned to the same note, then
make the 'lower' octave string an uppercase letter.  The tuning that is input will
be validated against the number of strings the instrument should have.
```go
if err := player.Tune("Dadgbe"); err != nil {
	log.Fatal("error tuning guitar")
}
```


#### Current instruments supported

| Instrument      | Strings | Default Tuning |
| :---------      | ------- | -------------: |
| Guitar          | 6       | Eadgbe         |
| Bass Guitar     | 4       | Eadg           |
| Ukulele         | 4       | Gcea           |
| 7 String Guitar | 7       | BEadgbe        |
| Mandolin        | 4       | Gdea           |
| 5 String Bass   | 5       | BEadg          |

#### Contributing

If you have an instrument that you would like added to this package, please get in touch or open an issue or even a pull request.

I would be extremely enthused if you have any suggestions or contribute.  Please fork and open a pull request.

