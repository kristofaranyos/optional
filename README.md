# optional

![License](https://img.shields.io/github/license/kristofaranyos/optional?style=flat-square)
![Release](https://img.shields.io/github/v/release/kristofaranyos/optional?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/kristofaranyos/optional?style=flat-square)](https://goreportcard.com/report/github.com/kristofaranyos/optional)

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

There are multiple mutator and getter methods defined on the optional type, you can check them in the source code for now (it's simple :) ).
