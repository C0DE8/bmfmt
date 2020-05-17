package bmfmt

import (
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

// TestBeautifyCauseAnErrorOnNonMap test the "Beautify" function to return an error on given argument is not a map
func TestBeautifyCauseAnErrorOnNonMap(t *testing.T) {

	candidate := "not a map, just a string"

	err := Beautify(candidate)

	if !assert.Error(t, err, "error is expected") {
		t.FailNow()
	}
	assert.Equal(t, "given argument is NOT a map", err.Error())
}

// TestBeautifyCauseAnErrorOnUnknownMapStructure test the "Beautify" function with an unknown map structure
func TestBeautifyCauseAnErrorOnUnknownMapStructure(t *testing.T) {

	candidate := map[string]struct{}{}

	err := Beautify(candidate)

	if !assert.Error(t, err, "error is expected") {
		t.FailNow()
	}
	assert.Equal(t, "(currently) unknown map structure: map[string]struct", err.Error())
}

// TestBeautifyOutputSimpleMapWithStringKeysAndStringValues test a simple map with string keys and string values
func TestBeautifyOutputSimpleMapWithStringKeysAndStringValues(t *testing.T) {
	tests := map[string]struct {
		in  map[string]string
		out string
	}{
		"empty key/value": {
			in: map[string]string{
				"a key": "",
			},
			out: `[ "a key"  string(  5) ]: ""  string(  0)
`,
		},
		"simple key/value": {
			in: map[string]string{
				"a key": "a short value",
			},
			out: `[ "a key"  string(  5) ]: "a short value"  string( 13)
`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var err error
			output := capturer.CaptureStdout(func() {
				err = Beautify(tc.in)
			})

			assert.NoError(t, err, "nil is expected")
			assert.Equal(t, tc.out, output, "beautified map output expected")
		})
	}
}

// TestBeautifyOutputSimpleMapWithStringKeysAndSliceOfStringValues test a valid map[string][]string output
func TestBeautifyOutputSimpleMapWithStringKeysAndSliceOfStringValues(t *testing.T) {
	tests := map[string]struct {
		in  map[string][]string
		out string
	}{
		"single value": {
			in: map[string][]string{
				"a key": {"a single value"},
			},
			out: `[ "a key"  string(  5) ]: "a single value"  string( 14)
`,
		},
		"multiple values": {
			in: map[string][]string{
				"next key": {"value one", "value two"},
			},
			out: `[ "next key"  string(  8) ]: "value one", "value two"  string( 22)
`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var err error
			output := capturer.CaptureStdout(func() {
				err = Beautify(tc.in)
			})

			assert.NoError(t, err, "nil is expected")
			assert.Equal(t, tc.out, output, "beautified map output expected")
		})
	}
}
