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
	"errors"

	"github.com/QuentinPerez/serror"
)

func main() {
	fmt.Println(serror.NewString("foo"))
	fmt.Println(serror.NewError(errors.New("bar")))
}
```
> Output
```console
main.go, line 11: main()
	➜ foo
main.go, line 12: main()
	➜ bar
```

## Changelog

### master (unreleased)

* Add `NewString`
* Add `NewError`

## Development

Feel free to contribute :smiley::beers:


## License

MIT
