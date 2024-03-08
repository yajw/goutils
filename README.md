# goutils
a collection of golang utility for make your life easier

requires go version `>=1.18`

# Examples

```go
import (
    "github.com/yajw/goutils/val"
)

func eg() {
	var p *string = nil
	println(nilOr(p, "hello"))  // => hello

	x := "olleh"
	p = &x
	println(nilOr(p, "hello"))  // => olleh

	var q *int = nil
	println(nilOr(q, int(1)))   // => 1

	y := 123
	q = &y
	println(nilOr(q, int(1)))  // => 123
}
```
