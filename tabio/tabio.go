package tabio

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Guitarbum722/tablature/instrument"
	"io"
)

// TablatureWriter embeds a buffered writer
type TablatureWriter struct {
	*bufio.Writer
}

// NewTablatureWriter creates a buffered writer to be used for staging tablature
func NewTablatureWriter(w io.Writer) *TablatureWriter {
	return &TablatureWriter{
		bufio.NewWriter(w),
	}
}

// StageTablature writes the Instrument's current tablature to w for buffering.
// The purpose is only to stage or buffer the current tablature but it does not
// write the tablature to a file.
func StageTablature(i instrument.Instrument, w *TablatureWriter) error {
	_, err := w.WriteString(instrument.StringifyCurrentTab(i))
	if err != nil {
		errMsg := fmt.Sprintf("there was an error writing to the bufferred writer: %s", err)
		return errors.New(errMsg)
	}
	// w.Flush()
	return nil
}

// ExportTablature will flush the bufferred writer to the io.Writer of which it was initialized
func ExportTablature(w *TablatureWriter) {
	w.Flush()
}
