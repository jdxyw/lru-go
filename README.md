# lru-go

[![Go](https://github.com/jdxyw/lru-go/actions/workflows/go.yml/badge.svg)](https://github.com/jdxyw/lru-go/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jdxyw/lru-go)](https://goreportcard.com/report/github.com/jdxyw/lru-go)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/jdxyw/lru-go/main/LICENSE)
[![codecov](https://codecov.io/gh/jdxyw/lru-go/branch/master/graph/badge.svg?token=e7NNqTMSYi)](https://codecov.io/gh/jdxyw/lru-go)

# Install

Install this package through `go get`.

```
go get github.com/jdxyw/lru-go
```

# Simple Usage

```go
package main

import (
	"fmt"
	"github.com/jdxyw/lru-go"
)

func main() {
	cache := lru.NewCache(100)
	cache.Add("Go", 1)
	val, _ := cache.Get("Go")

	fmt.Printf("The value for Go is %v.", val)
}
```