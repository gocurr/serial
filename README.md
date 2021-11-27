# Serial - continuous range tool

To download, run:

```bash
go get -u github.com/gocurr/serial
```

Import it in your program as:

```go
import "github.com/gocurr/serial"
```

It requires Go 1.11 or later due to usage of Go Modules.

- Usage:

```go
data := []interface{}{nil, 1, V{Val: 5}, V{Val: 3}}
ranges := Ranges(data, 2, func(i interface{}) bool {
    v, ok := i.(V)
    if !ok {
        return false
    }
    return v.Val > 2
})
fmt.Println(ranges)
```
