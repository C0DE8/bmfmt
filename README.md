# bmfmt Golang [b]eautify [m]ap display [fmt] for Humans

## Install

``` shell
# Stable version
go get -u -v gopkg.in/bmfmt.v1

# Latest version
go get -u -v github.com/c0de8/bmfmt
```

## Usage

[API Documentation](https://godoc.org/github.com/c0de8/bmfmt)

[Examples](https://github.com/c0de8/bmfmt/blob/master/examples/main.go)

``` golang
package main

import (
	bmfmt "github.com/c0de8/bmfmt"
)

func main() {
	example()
}

func example() {
	
	m := map[string][]string{
		"Warc-Type": { "response" }
	}
	
	fmt.Println(m) // fmt the default formatting.
	/*
		map[  ]
	*/

	bmfmt.Beautify(m) // More friendly formatting.
	/*
[ "Warc-Type"            string(  9) ]: "response"                                                                                           string(  8)
[ "Warc-Date"            string(  9) ]: "2014-08-02T09:52:13Z"                                                                               string( 20)
[ "Content-Length"       string( 14) ]: "43428"                                                                                              string(  5)
[ "Warc-Payload-Digest"  string( 19) ]: "sha1:M63W6MNGFDWXDSLTHF7GWUPCJUH4JK3J"                                                              string( 37)
[ "Warc-Block-Digest"    string( 17) ]: "sha1:YHKQUSBOS4CLYFEKQDVGJ457OAPD6IJO"                                                              string( 37)
[ "Warc-Truncated"       string( 14) ]: "length"                                                                                             string(  6)
[ "Warc-Record-Id"       string( 14) ]: "<urn:uuid:ffbfb0c0-6456-42b0-af03-3867be6fc09f>", "<urn:uuid:ffbfb0c0-6456-42b0-af03-xxxxxxxxxxx>"  string( 97)
[ "Content-Type"         string( 12) ]: "application/http; msgtype=response"                                                                 string( 34)
[ "Warc-Warcinfo-Id"     string( 16) ]: "<urn:uuid:3169ca8e-39a6-42e9-a4e3-9f001f067bdf>"                                                    string( 47)
[ "Warc-Concurrent-To"   string( 18) ]: "<urn:uuid:d99f2a24-158a-4c77-bb0a-3cccd40aad56>"                                                    string( 47)
[ "Warc-Ip-Address"      string( 15) ]: "212.58.244.61"                                                                                      string( 13)
[ "Warc-Target-Uri"      string( 15) ]: "http://news.bbc.co.uk/2/hi/africa/3414345.stm"                                                      string( 45)

	*/

}


```

## License

bmfmt is licensed under the MIT License. See [LICENSE](https://github.com/go-ffmt/ffmt/blob/master/LICENSE) for the full license text.
