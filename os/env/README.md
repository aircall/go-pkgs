# Environment variables

## Usage

Get the package:

```console
go get -u github.com/aircall/go-pkgs/os/env
```

Use it:

```go
package main

import (
	"log"

	"github.com/aircall/go-pkgs/os/env"
)

func main() {
  // default value
  fmt.Println(env.Get("AWS_REGION", "us-west-1"))

  // mandatory value, should panic if not set
  fmt.Println(env.Get("DB_NAME"))
}
```
