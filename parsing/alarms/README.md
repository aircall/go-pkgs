# alarms parsing

## Usage

Get the package:

```console
go get -u github.com/aircall/go-pkgs/parsing/alarms
```

Use it:

```go
package main

import (
	"log"

	"github.com/aircall/go-pkgs/parsing/alarms"
)

func main() {
	alarmName := "InstanceHealth_production_aircall-web-42_aircall-web-42_i-0ab77c46031bd8260"
	a, err := alarms.Parse(alarmName)

	if err != nil {
		log.Fatal("Error parsing alarm with: ", err.Error())
	}

	log.Println(a.Name, a.Project, a.Component, a.Environment)
}
```
