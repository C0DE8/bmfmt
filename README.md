# bmfmt Golang [b]eautify [m]ap display [fmt] for Humans

## Install

``` shell
# Stable version
go get -u -v gopkg.in/c0de8/bmfmt.v0

# Latest version
go get -u -v github.com/c0de8/bmfmt
```

## Usage

[API Documentation](https://godoc.org/github.com/c0de8/bmfmt)

[Examples](https://github.com/c0de8/bmfmt/blob/master/examples/main.go)

``` golang
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
		"some-key": { "response" },
		"Another-Hash-Key": { "first value", "second value" },
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

```

## License

bmfmt is licensed under the MIT License. See [LICENSE](https://github.com/go-ffmt/ffmt/blob/master/LICENSE) for the full license text.
