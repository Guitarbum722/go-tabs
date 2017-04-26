package tabio

import (
	"bufio"
	"github.com/Guitarbum722/go-tabs/instrument"
	"github.com/pkg/errors"
	"io"
)

// TablatureWriter embeds a buffered writer
// The wrapPosition can be used to choose how far the tablature section will reach before wrapping to the next.
// 20 is the default.
type TablatureWriter struct {
	*bufio.Writer
	wrapPosition int
	totalLength  int
	tb           tablatureBuilder
}

type tablatureBuilder struct {
	builder map[byte][]byte
}

// UpdateWrapPosition will update the wrapPosition of the TablatureWriter
// which will have a default value of 20 if the input is less than 20 for
// readability
func (tw *TablatureWriter) UpdateWrapPosition(pos int) {
	if pos < 20 {
		pos = 20
	}
	tw.wrapPosition = pos
}

// NewTablatureWriter creates a buffered writer to be used for staging tablature
func NewTablatureWriter(w io.Writer, pos int) *TablatureWriter {
	if pos < 20 {
		pos = 20
	}
	return &TablatureWriter{
		bufio.NewWriter(w),
		pos,
		0,
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

		w.tb.builder[k] = append(w.tb.builder[k], []byte(v)...)
		w.totalLength = len(w.tb.builder[k])
	}

	return
}

// ExportTablature will flush the bufferred writer to the io.Writer of which it was initialized
func ExportTablature(i instrument.Instrument, w *TablatureWriter) error {

	var done int

	for ; done < w.totalLength; done += w.wrapPosition {

		for _, v := range i.Order() {

			w.Write([]byte{v, ':', ' '})

			if (done + w.wrapPosition) < w.totalLength {
				if _, err := w.Write(w.tb.builder[v][done:(done + w.wrapPosition)]); err != nil {
					return errors.Wrap(err, "write to buffer failed")
				}
			} else {
				if _, err := w.Write(w.tb.builder[v][done:(w.totalLength)]); err != nil {
					return errors.Wrap(err, "write to buffer failed")
				}
			}
			w.WriteByte('\n')
		}
		w.WriteByte('\n')
	}

	w.Flush()

	return nil
}
