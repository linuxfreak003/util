# util
[![Documentation](https://godoc.org/github.com/linuxfreak003/util?status.svg)](http://godoc.org/github.com/linuxfreak003/util)

Generic utility library in Go

## Library

### Installation

```bash
go get -u github.com/linuxfreak003/util
```

or use `go mod`

### Example

```go
package main

import (
    "github.com/linuxfreak003/util"
    "github.com/linuxfreak003/util/slice"
)

func main() {
    _ = util.Round(1.2345, 3) // Returns 1.235

    _ = slice.Map([]int{1,2,3,4}, func(i int) float64{
        return float64(i * i)
    }) // will return []int{1.0, 4.0, 9.0, 16.0}
}
```

The available functions are listed in documentation. Another good place to look for an example is in `util_test.go`.
