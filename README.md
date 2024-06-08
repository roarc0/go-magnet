# 🧲 Go Magnet

[![Go Reference](https://pkg.go.dev/badge/image)](https://pkg.go.dev/github.com/roarc0/go-magnet)
[![Go Report](https://goreportcard.com/badge/github.com/roarc0/go-magnet)](https://goreportcard.com/report/github.com/roarc0/go-magnet)
[![Go Coverage](https://github.com/roarc0/go-magnet/wiki/coverage.svg)](https://raw.githack.com/wiki/roarc0/go-magnet/coverage.html)
![go workflow](https://github.com/roarc0/go-magnet/actions/workflows/go.yml/badge.svg)

This is a zero external dependencies magnet link parser that can be used to extract information from magnet links and create magnet links.

## Install 💾

```bash
go get github.com/roarc0/go-magnet
```

## Usage 🏄

```go
package main

import (
  "fmt"
  "github.com/roarc0/go-magnet"
)

func main() {
  magnet, err := magnet.Parse("magnet:?...")
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(magnet)

  // Back to a string
  fmt.Println(magnet.String())
}
```

## Credits :star:

- [Alessandro Rosetti](https://github.com/roarc0)

## License :scroll:

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
