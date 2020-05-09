package main

import (
	"fmt"
	bmfmt "github.com/c0de8/bmfmt"
)

func main() {
	example()
}

func example() {

	m := map[string][]string{
		"some-key":         {"response"},
		"Another-Hash-Key": {"first value", "second value"},
	}

	fmt.Println(m) // fmt the default formatting
	/*
	   map[some-key:[response] Another-Hash-Key:[first value second value]]
	*/

	err := bmfmt.Beautify(m) // significant more friendly formatting
	if err != nil {
		fmt.Println("ERROR (bmfmt.Beautify): " + err.Error())
	}
	/*
	   [ "some-key"          string(  8) ]: "response"                     string(  8)
	   [ "Another-Hash-Key"  string( 16) ]: "first value", "second value"  string( 23)
	*/

}
