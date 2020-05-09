package bmfmt

import (
	"errors"
	"github.com/c0de8/bmfmt"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"testing"
)

var notAMapError = errors.New("given argument is NOT a map")

// TestBeautifyCauseAnErrorOnNonMap test the "Beautify" function to return an error on given argument is not a map
func TestBeautifyCauseAnErrorOnNonMap(t *testing.T) {

	candidate := "not a map, just a string"
	expectedError := notAMapError

	err := bmfmt.Beautify(candidate)
	assert.Equal(t, expectedError, err)
}

// TestBeautifyOutputASimpleMap test a valid map[string][]string output
func TestBeautifyOutputASimpleMap(t *testing.T) {

	candidate := map[string][]string{
		"a key":    {"a single value"},
		"next key": {"value one", "value two"},
	}

	// due to the fact, that maps NOT sorted, the output order may vary
	expectedOutput1 := `[ "a key"     string(  5) ]: "a single value"          string( 14)
[ "next key"  string(  8) ]: "value one", "value two"  string( 22)
`
	expectedOutput2 := `[ "next key"  string(  8) ]: "value one", "value two"  string( 22)
[ "a key"     string(  5) ]: "a single value"          string( 14)
`
	var err error

	output := capturer.CaptureStdout(func() {
		err = bmfmt.Beautify(candidate)
	})

	assert.Nil(t, err)

	firstTry := assert.Equal(t, expectedOutput1, output)

	if firstTry != true {
		assert.Equal(t, expectedOutput2, output)
	}
}
