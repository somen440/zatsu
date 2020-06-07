package foo

import "fmt"

type Hoge interface {
	HogeMethod() string
	Bar() Bar
}

type Bar interface {
	BarMethod() string
}

func Exec(h Hoge) {
	fmt.Println("\t" + h.HogeMethod())
	fmt.Println("\t" + h.Bar().BarMethod())
}
