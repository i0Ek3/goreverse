<h1 align="center">goreverse 👋</h1>
<p>
</p>

> Reverse multi values in Go.

## Getting Started

### Installation

`go get https://github.com/i0Ek3/goreverse`

### Import

```Go
import (
    gr "github.com/i0Ek3/goreverse"
)

func main() {
    b := "test"
    gr.ReverseWithInterval([]byte(b), 0, 2)

    res := []int{1, 2, 3}
    gr.Reverse(res)
}

```

## TODO

- support reverse any type data, just like interface{}


## Contributing

PRs and Issues are also welcome.


## Show your support

Give a ⭐️ if this project helped you!

