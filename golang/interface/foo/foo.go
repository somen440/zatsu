package foo

import (
	"fmt"

	"github.com/somen440/zatsu/golang/interface/bar"
)

type Hoge interface {
	HogeMethod() string
	Bar() *bar.Bar
}

func Exec(h Hoge) {
	fmt.Println("\t" + h.HogeMethod())
	fmt.Println("\t" + h.Bar().BarMethod())
}
