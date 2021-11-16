# Serial - continouse range tool

To download, run:

```bash
go get -u github.com/gocurr/serial
```

If it doesn't work, try to remove `go/pkg/sumdb`

Import it in your program as:

```go
import "github.com/gocurr/serial"
```

It requires Go 1.11 or later due to usage of Go Modules.

- Usage:

```go
data := []interface{}{-2, -1, 0, 1, 2, 3, 4, 5, 6, 3, 3, 2, 3, 3, 3, 3, 4}
ranges := serial.Ranges(data, 2, func (v interface{}) bool {
return v.(int) > 2 && v.(int) < 5
})
fmt.Println(ranges)
```
