package tabio

import (
	"os"
	"testing"
)

func TestNewTablatureWriter(t *testing.T) {
	for _, tt := range newTablatureWriterCases {
		if got := NewTablatureWriter(os.Stdout, tt.inWrapPos); got.wrapPosition != tt.expectedPos {
			t.Fatalf("NewTablatureWriter(os.Stdout, %d) = %d and got %d.", tt.inWrapPos, got.wrapPosition, tt.expectedPos)
		}
	}
}
