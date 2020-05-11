package bmfmt

import (
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"testing"
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

// TestBeautifyOutputSimpleMapWithStringKeysAndStringValues test a simple map with string keys and string values
func TestBeautifyOutputSimpleMapWithStringKeysAndStringValues(t *testing.T) {

	var err error
	candidateSingleValue := map[string]string{
		"a key": "a short value",
	}

	expectedOutputSingleValue := `[ "a key"  string(  5) ]: "a short value"  string( 13)
`
	output := capturer.CaptureStdout(func() {
		err = Beautify(candidateSingleValue)
	})

	assert.NoError(t, err, "nil is expected")
	assert.Equal(t, expectedOutputSingleValue, output, "beautified map output expected (1)")
}

// TestBeautifyOutputSimpleMapWithStringKeysAndSliceOfStringValues test a valid map[string][]string output
func TestBeautifyOutputSimpleMapWithStringKeysAndSliceOfStringValues(t *testing.T) {

	var err error
	candidateSingleValue := map[string][]string{
		"a key": {"a single value"},
	}

	// due to the fact, that maps are NOT sorted, the output order may vary, so we can test 1 element in map only
	expectedOutputSingleValue := `[ "a key"  string(  5) ]: "a single value"  string( 14)
`
	output := capturer.CaptureStdout(func() {
		err = Beautify(candidateSingleValue)
	})

	assert.NoError(t, err, "nil is expected")
	assert.Equal(t, expectedOutputSingleValue, output, "beautified map output expected (1)")

	candidateMultipleValue := map[string][]string{
		"next key": {"value one", "value two"},
	}

	expectedOutputMultipleValue := `[ "next key"  string(  8) ]: "value one", "value two"  string( 22)
`
	output = capturer.CaptureStdout(func() {
		err = Beautify(candidateMultipleValue)
	})

	assert.NoError(t, err, "nil is expected")
	assert.Equal(t, expectedOutputMultipleValue, output, "beautified map output expected (2)")
}
