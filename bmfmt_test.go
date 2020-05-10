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

// TestBeautifyOutputASimpleMap test a valid map[string][]string output
func TestBeautifyOutputASimpleMap(t *testing.T) {

	var err error
	candidateSingleValue := map[string][]string{
		"a key": {"a single value"},
	}

	// due to the fact, that maps are NOT sorted, the output order may vary, so we ned to test both orderings
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
