# optional

![License](https://img.shields.io/github/license/kristofaranyos/optional)
![Release](https://img.shields.io/github/v/release/kristofaranyos/optional)
[![Go test](https://github.com/kristofaranyos/optional/actions/workflows/test.yml/badge.svg)](https://github.com/kristofaranyos/optional/actions/workflows/test.yml)
[![Go Coverage](https://github.com/kristofaranyos/optional/wiki/coverage.svg)](https://raw.githack.com/wiki/kristofaranyos/optional/coverage.html)
[![Go Report Card](https://goreportcard.com/badge/github.com/kristofaranyos/optional)](https://goreportcard.com/report/github.com/kristofaranyos/optional)

[comment]: # (TODO: add CI job for build and codecov, then add badges for both)

Optional is a go library that adds a wrapper for representing values that might be omitted.

## Installation

`$ go get -u github.com/kristofaranyos/optional@latest`

## Usage

```go
package main

import (
	"fmt"
	"github.com/kristofaranyos/optional"
)

func main() {
	filledOptional := optional.New("Hello world!")
	emptyOptional := optional.Empty[string]() // Empty requires explicit generic type initialization

	printWithOptional(filledOptional)
	printWithOptional(emptyOptional)
}

func printWithOptional(o optional.T[string]) {
	val, ok := o.Get()
	if !ok {
		fmt.Println("Optional is not set.")
		return
	}

	fmt.Println("Optional is set:", val)
}
```

Try it on the [Go Playground](https://go.dev/play/p/W7UfH8G9PqK).

There are multiple mutator and getter methods defined on the optional type, you can check them in the source code for
now (it's simple :) ).

## Compatibility

### JSON

The `optional.T` type implements `json.Marshaler` and `json.Unmarshaler` for convenient conversion between Go types and
JSON.  
One caveat to this is that the wrapped type also needs to be marshallable in order for this to work.

### Databases

TODO
