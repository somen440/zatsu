package hoge

import (
	"github.com/somen440/zatsu/golang/interface/foo"
)

type Hoge struct {
	bar *Bar
}

func New(bar *Bar) *Hoge {
	return &Hoge{
		bar: bar,
	}
}

func (h *Hoge) HogeMethod() string {
	return "hoge method."
}

func (h *Hoge) Bar() foo.Bar {
	return h.bar
}

type Bar struct{}

func NewBar() *Bar {
	return &Bar{}
}

func (bar *Bar) BarMethod() string {
	return "bar method."
}
