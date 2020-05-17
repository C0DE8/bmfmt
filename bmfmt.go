// Package bmfmt provides a function to output a map in a human (beautified) readable way
package bmfmt

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var keyTypePadLen = 3
var valueTypePadLen = 3

var keyDelim = map[string]string{
	"open":  "[",
	"close": "]",
}
var keyValueSep = ": "
var valDelim = map[string]string{
	"open":  "",
	"close": "",
}

var rowTmpl = keyDelim["open"] + " %v  %v " + keyDelim["close"] +
	keyValueSep +
	valDelim["open"] + "%v  %v" + valDelim["close"] + "\n"

// Beautify prints a given map (any key and value type) to  better readable format
func Beautify(a interface{}) error {

	// prevent argument a NOT to be a map
	if getArgType(a) != reflect.Map {
		return errors.New("given argument is NOT a map")
	}

	keyType := getMapKeyType(a)
	valueType := getMapValueType(a)

	switch {

	// map structure is => map[string]string
	case keyType == reflect.String && // is key of type <string>
		valueType == reflect.String: // is value of type <string>

		var targetType map[string]string
		var convertedArg = reflect.ValueOf(a).Convert(reflect.TypeOf(targetType)).Interface().(map[string]string)

		printMapKeyStringWithStrings(convertedArg)

	// map structure is => map[string][]string
	case keyType == reflect.String && // is key of type <string>
		valueType == reflect.Slice && // is value of type <slice>
		reflect.TypeOf(a).Elem() == reflect.TypeOf([]string{}): // is slice value if type <string>

		var targetType map[string][]string
		var convertedArg = reflect.ValueOf(a).Convert(reflect.TypeOf(targetType)).Interface().(map[string][]string)

		printMapKeyStringWithSliceOfStrings(convertedArg)

	default:
		// handle unknown map type/structure
		return errors.New(fmt.Sprintf("(currently) unknown map structure: map[%v]%v", keyType, valueType))
	}

	return nil
}

// printMapKeyStringWithString prints out a map with string keys and strings as value
func printMapKeyStringWithStrings(aMap map[string]string) {
	maxKeyLength, maxValueLength := getMaxLengthOfKeyStringValueString(aMap)

	for key, value := range aMap {
		maxKeyStringLen := len("string(" + string(maxKeyLength) + ")")
		keyType := fmt.Sprintf("string(%v)", leftPad(strconv.Itoa(len(key)), " ", keyTypePadLen))
		maxValueStringLen := len("string(" + string(maxValueLength) + ")")
		valueType := fmt.Sprintf("string(%v)", leftPad(strconv.Itoa(len(value)), " ", valueTypePadLen))
		fmt.Printf(
			rowTmpl,
			rightPad("\""+key+"\"", " ", maxKeyLength+2),
			rightPad(keyType, " ", maxKeyStringLen+1),
			rightPad("\""+value+"\"", " ", maxValueLength+2),
			leftPad(valueType, " ", maxValueStringLen+1),
		)
	}
}

// getMaxLengthOfKeyStringValueString detects the max. length of the key, value and value count
func getMaxLengthOfKeyStringValueString(aMap map[string]string) (int, int) {
	var maxKeyLength = 0
	var maxValueLength = 0

	for key, value := range aMap {
		if len(key) > maxKeyLength {
			maxKeyLength = len(key)
		}
		if len(value) > maxValueLength {
			maxValueLength = len(value)
		}
	}

	return maxKeyLength, maxValueLength
}

// printMapKeyStringWithSliceOfStrings prints out a map with string keys and slice of strings as value
func printMapKeyStringWithSliceOfStrings(aMap map[string][]string) {
	maxKeyLength, maxValueLength, maxValueCount := getMaxLengthOfKeyStringSliceOfStrings(aMap)

	for key, value := range aMap {
		stringValue := strings.Join(value, "\", \"")
		maxKeyStringLen := len("string(" + string(maxKeyLength) + ")")
		keyTypeString := fmt.Sprintf("string(%v)", leftPad(strconv.Itoa(len(key)), " ", keyTypePadLen))
		valueType := fmt.Sprintf("string(%v)", leftPad(strconv.Itoa(len(stringValue)), " ", valueTypePadLen))
		fmt.Printf(
			rowTmpl,
			rightPad("\""+key+"\"", " ", maxKeyLength+2),
			rightPad(keyTypeString, " ", maxKeyStringLen+1),
			rightPad("\""+stringValue+"\"", " ", maxValueLength+(maxValueCount*2)),
			leftPad(valueType, " ", valueTypePadLen),
		)
	}
}

// getMaxLengthOfKeyStringSliceOfStrings detects the max. length of the key, value and value count
func getMaxLengthOfKeyStringSliceOfStrings(aMap map[string][]string) (int, int, int) {
	var maxKeyLength = 0
	var maxValueLength = 0
	var maxValueCount = 0

	for key, valueSlice := range aMap {
		if len(key) > maxKeyLength {
			maxKeyLength = len(key)
		}
		value := strings.Join(valueSlice, ", ")
		if len(value) > maxValueLength {
			maxValueLength = len(value)
		}
		if len(valueSlice) > maxValueCount {
			maxValueCount = len(valueSlice)
		}
	}

	return maxKeyLength, maxValueLength, maxValueCount
}

// leftPad just repeat the padStr the indicated
// number of times
func leftPad(s string, padStr string, pLen int) string {
	padLen := 0
	if len(s) < pLen {
		padLen = pLen - len(s)
	}
	return strings.Repeat(padStr, padLen) + s
}

// rightPad just repeat the padStr the indicated
// number of times
func rightPad(s string, padStr string, pLen int) string {
	padLen := 0
	if len(s) < pLen {
		padLen = pLen - len(s)
	}
	return s + strings.Repeat(padStr, padLen)
}

// getArgType return the argument type (Kind) [type: reflect.Kind]
func getArgType(a interface{}) reflect.Kind {
	return reflect.TypeOf(a).Kind()
}

// getMapKeyType return the map key type (Kind) [type: reflect.Kind]
func getMapKeyType(a interface{}) reflect.Kind {
	return reflect.TypeOf(a).Key().Kind()
}

// getMapValueType return the map element type (Kind) [type: reflect.Kind]
func getMapValueType(a interface{}) reflect.Kind {
	return reflect.TypeOf(a).Elem().Kind()
}
