你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
