# go-formula [![GoDoc](http://godoc.org/github.com/Sandertv/go-formula?status.svg)](https://godoc.org/github.com/Sandertv/go-formula/v2)
A simple and fast formula parser and evaluator.

## Getting Started

### Usage
Formulas may be parsed using the formula.New() function. The function returns a formula that may be evaluated
an unlimited amount of times. Note that parsing formulas is generally heavier than evaluating them, so it is
recommended to parse once and evaluate the same formula multiple times where applicable.

```go
package main

import (
	"github.com/sandertv/go-formula/v2"
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
    log.Printf("Formula output: %v", f.MustEval(x, z))
}
```

### Documentation
https://godoc.org/github.com/Sandertv/go-formula/v2