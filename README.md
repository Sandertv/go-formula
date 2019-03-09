# go-formula
A simple and fast formula parser and evaluator.

## Getting Started

### Usage
```go
package main

import (
	"github.com/sandertv/go-formula"
	"log"
)

func main() {
    f, err := formula.New("17*x + pow(z*3, 3)")
    if err != nil {
        log.Print(err)
        return
    }
    x := formula.Var("x", 4.5)
    z := formula.Var("z", 5)
    log.Printf("Formula output: %v", f.Eval(x, z))
}
```

### Documentation
https://godoc.org/github.com/Sandertv/go-formula