<h1 align="center">go-findroot</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Completed-100%25-green.svg?style=flat-square">
</p>

<p align="center">
  Locate the root directory of a project using Git via the command line
</p>

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/integralist/go-findroot/find"
)

func main() {
  root, err := find.Repo()
  if err != nil {
    log.Fatalf("Error: %s", err.Error())
  }

  fmt.Printf("%+v", root)
  // {Name:go-findroot Path:/Users/M/Projects/golang/src/github.com/integralist/go-findroot}
}
```

## Tests

```go
go test -v ./...
```

## Licence

[The MIT License (MIT)](http://opensource.org/licenses/MIT)

Copyright (c) 2016 [Mark McDonnell](http://twitter.com/integralist)
