package tabio

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Guitarbum722/go-tabs/instrument"
	"io"
)

// TablatureWriter embeds a buffered writer
// The wrapPosition can be used to choose how far the tablature section will reach before wrapping to the next.
// 20 is the default.
type TablatureWriter struct {
	*bufio.Writer
	WrapPosition int
	tb           tablatureBuilder
}

type tablatureBuilder struct {
	builder map[byte][]byte
}

// NewTablatureWriter creates a buffered writer to be used for staging tablature
func NewTablatureWriter(w io.Writer, pos int) *TablatureWriter {
	return &TablatureWriter{
		bufio.NewWriter(w),
		20,
		tablatureBuilder{
			make(map[byte][]byte),
		},
	}
}

// StageTablature writes the Instrument's current tablature to w for buffering.
// The purpose is only to stage or buffer the current tablature but it does not
// write the tablature to a file.
func StageTablature(i instrument.Instrument, w *TablatureWriter) {

	for k, v := range i.Fretboard() {

		// labels each instrument's "string" with a prefix
		if len(w.tb.builder[k]) == 0 {
			w.tb.builder[k] = append(w.tb.builder[k], k, ':', ' ')
		}

		// adds the first newline after each line of tablature.
		// the newline is added to the end of each line of tablature
		if !bytes.Contains(w.tb.builder[k], []byte("\n")) {
			w.tb.builder[k] = append(w.tb.builder[k], []byte(v)...)
			w.tb.builder[k] = append(w.tb.builder[k], '\n')
		} else {
			w.tb.builder[k] = append(w.tb.builder[k][:len(w.tb.builder[k])-1], []byte(v)...)
			w.tb.builder[k] = append(w.tb.builder[k], '\n')
		}
	}

	return
}

// ExportTablature will flush the bufferred writer to the io.Writer of which it was initialized
func ExportTablature(i instrument.Instrument, w *TablatureWriter) error {

	for _, v := range i.Order() {

		fmt.Fprint(w, string(w.tb.builder[v]))
	}

	w.Flush()

	return nil
}
