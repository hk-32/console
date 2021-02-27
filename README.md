# console
This is a lightweight replacement for the fmt package, reduces the binary size by roughly 400kb in a hello world program.

*Please note: This package achieves smaller binaries primarily by not relying on reflection. So if your program does use reflection itself then this won't benefit you much.*

## Getting Started

### Installing

To start using wire, install Go and run `go get`:

```sh
$ go get -u github.com/hk-32/console
```

This will retrieve the library. Specifically the v1.0.0 right now. Works perfecly fine with modules.

### Samples
Hello World:

```go
package main

import "github.com/hk-32/console"

func main() {
    console.WriteLine("Hello World")
}
```

Or get some user input:

```go
package main

import "github.com/hk-32/console"

func main() {
    // name := console.ReadLine()
    name := console.Input("Please enter your name: ")
    console.WriteLine("Hello", name)
}
```

## Contact

Hassan Khan: HK.32@outlook.com

## License
`console` source code is available under the MIT [License](/LICENSE).
