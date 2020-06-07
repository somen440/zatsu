package foo

import "fmt"

type Hoge interface {
	HogeMethod() string
}

func Exec(h Hoge) {
	fmt.Println("\t" + h.HogeMethod())
}
