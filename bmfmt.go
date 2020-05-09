package bmfmt

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Beautify prints a given map (any key and value type) to a better readable format
func Beautify(a interface{}) error {

	// prevent argument a NOT to be a map
	if getArgType(a) != reflect.Map {
		return errors.New("given argument is NOT a map")
	}

	keyType := getMapKeyType(a)
	valueType := getMapValueType(a)

	switch {

	// map structure is => map[string][]string
	case keyType == reflect.String && // is key of type <string>
		valueType == reflect.Slice && // is value of type <slice>
		reflect.TypeOf(a).Elem() == reflect.TypeOf([]string{}): // is slice value if type <string>

		var targetType map[string][]string
		var convertedArg = reflect.ValueOf(a).Convert(reflect.TypeOf(targetType)).Interface().(map[string][]string)

		printMapKeyStringWithSliceOfStrings(convertedArg)

	default:
		// handle unknown type
		return errors.New(fmt.Sprintf("(currently) unknown map structure: map[%v]%v", keyType, valueType))
	}

	return nil
}

// printMapKeyStringWithSliceOfStrings prints out a map with string keys and slice of strings as value
func printMapKeyStringWithSliceOfStrings(aMap map[string][]string) {

	var stringPadLen = 3

	maxKeyLength, maxValueLength, maxValueCount := getMaxLength(aMap)

	for key, value := range aMap {
		stringValue := strings.Join(value, "\", \"")
		maxKeyStringLen := len("string(" + string(maxKeyLength) + ")")
		keyTypeString := fmt.Sprintf("string(%v)", leftPad(strconv.Itoa(len(key)), " ", stringPadLen))
		fmt.Printf(
			"[ %v  %v ]: %v  string(%v)\n",
			rightPad("\""+key+"\"", " ", maxKeyLength+2),
			rightPad(keyTypeString, " ", maxKeyStringLen+1),
			rightPad("\""+stringValue+"\"", " ", maxValueLength+(maxValueCount*2)),
			leftPad(strconv.Itoa(len(stringValue)), " ", stringPadLen),
		)
	}
}

func getMaxLength(aMap map[string][]string) (int, int, int) {
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

func getArgType(a interface{}) reflect.Kind {
	return reflect.TypeOf(a).Kind()
}

func getMapKeyType(a interface{}) reflect.Kind {
	return reflect.TypeOf(a).Key().Kind()
}

func getMapValueType(a interface{}) reflect.Kind {
	return reflect.TypeOf(a).Elem().Kind()
}
