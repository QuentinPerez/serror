# Smart Error

:wrench: Smart Error (written in Golang)

## Install

Using go

- `go get github.com/QuentinPerez/serror`

## Usage
```go
package main

import (
	"fmt"

	"github.com/QuentinPerez/serror"
)

func main() {
	fmt.Println(serror.Errorf("Smart Error"))
}

```
> Output
```console
main.go:10 main() Smart Error
```

## Changelog

### master (unreleased)

* Add `Errorf`

## Development

Feel free to contribute :smiley::beers:


## License

MIT
