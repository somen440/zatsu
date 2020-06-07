package main

import (
	"fmt"

	"github.com/somen440/zatsu/golang/interface/bar"
	"github.com/somen440/zatsu/golang/interface/foo"
	"github.com/somen440/zatsu/golang/interface/hoge"
)

func main() {
	fmt.Println("interface test --")
	b := bar.New()
	h := hoge.New(b)
	foo.Exec(h)
}
