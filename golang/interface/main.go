package main

import (
	"fmt"

	"github.com/somen440/zatsu/golang/interface/foo"
	"github.com/somen440/zatsu/golang/interface/hoge"
)

func main() {
  fmt.Println("interface test --")
	foo.Exec(hoge.New(hoge.NewBar()))
}
