package hoge

import (
	"github.com/somen440/zatsu/golang/interface/bar"
)

type Bar interface {
	BarMethod() string
}

type Hoge struct {
	bar Bar
}

func New(bar Bar) *Hoge {
	return &Hoge{
		bar: bar,
	}
}

func (h *Hoge) HogeMethod() string {
	return "hoge method."
}

func (h *Hoge) Bar() *bar.Bar {
	return h.bar.(*bar.Bar)
}
