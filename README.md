# ntp-go

> :hourglass: Simple Network Time Protocol Implementation in Golang.

## Installation

```shell
~$ go get github.com/song940/ntp-go
```

## Example

### client

```go
package main

import (
	"log"

	"github.com/song940/ntp-go/ntp"
)

func main() {
	client := ntp.NewClient(
		"time.apple.com:123",
	)
	time, err := client.GetTime()
	log.Println(time, err)
}
```

### server

TODO: implement

## Node.js Implementation

> ⌛ simple network time protocol implementation for node.js 
>
> https://github.com/song940/node-ntp

## License

This project is licensed under the MIT License.