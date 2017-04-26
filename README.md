## ----GO--------
## -------TABS---

### A minimal and easy to use package for creating tablature for stringed / fretted instruments

#### What's included?
* A very ugly, yet easy to use command line interface
* COMING SOON! - A (nearly) just as ugly web interface
* A simple package that can be quickly and easily implemented by your own application

#### Usage
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

#### Current instruments supported

| Instrument      | Strings | Default Tuning |
| :---------      | ------- | -------------: |
| Guitar          | 6       | Eadgbe         |
| Bass Guitar     | 4       | Eadg           |
| Ukulele         | 4       | Gcea           |
| 7 String Guitar | 7       | BEadgbe        |
| Mandolin        | 4       | Gdea           |

#### Contributing

If you have an instrument that you would like added to this package, please get in touch or open an issue or even a pull request.

I would be extremely enthused if you have any suggestions or contribute.  Please fork and open a pull request.

