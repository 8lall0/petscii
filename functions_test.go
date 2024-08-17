package petscii

import (
	"bytes"
	"strings"
	"testing"
)

func TestAsciiToPetscii(t *testing.T) {
	var tests = []struct {
		in  string
		out []byte
	}{
		{"a", []byte{asciiToPetscii['A']}},
	}

	for _, tt := range tests {
		bb, err := AsciiToPetscii(tt.in)
		if err != nil {
			t.Errorf("asciiToPetscii failed: %v", err)
		}
		if bytes.Compare(bb, tt.out) != 0 {
			t.Errorf("asciiToPetscii failed: got %v, want %v", bb, tt.out)
		}
	}
}

func TestPetsciiToAscii(t *testing.T) {
	var tests = []struct {
		in  []byte
		out string
	}{
		{[]byte{asciiToPetscii['A']}, "a"},
	}

	for _, tt := range tests {
		bb, err := PetsciiToAscii(tt.in)
		if err != nil {
			t.Errorf("asciiToPetscii failed: %v", err)
		}
		if strings.Compare(bb, tt.out) != 0 {
			t.Errorf("asciiToPetscii failed: got %v, want %v", bb, tt.out)
		}
	}
}
