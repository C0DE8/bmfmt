# bmfmt Golang [b]eautify [m]ap display [fmt] for Humans

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/c0de8/bmfmt/graphs/commit-activity)
[![Github all releases](https://img.shields.io/github/downloads/c0de8/bmfmt/total.svg)](https://GitHub.com/c0de8/bmfmt/releases/)

## Install

``` shell
# stable version
go get -u -v gopkg.in/c0de8/bmfmt.v0

# latest version (may be unstable)
go get -u -v github.com/c0de8/bmfmt
```

## Supported Structures

- map[string]string
- map[string][]string

More structures are in implementation.

## Usage

[API Documentation](https://pkg.go.dev/github.com/c0de8/bmfmt?tab=doc)

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

bmfmt is licensed under the MIT License. See [LICENSE](https://github.com/c0de8/bmfmt/blob/master/LICENSE) for the full license text.
