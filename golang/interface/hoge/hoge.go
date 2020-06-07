package hoge

import (
	"github.com/somen440/zatsu/golang/interface/di"
)

type Hoge struct {
	bar di.Bar
}

func New(bar di.Bar) *Hoge {
	return &Hoge{
		bar: bar,
	}
}

func (h *Hoge) HogeMethod() string {
	return "hoge method."
}

func (h *Hoge) Bar() di.Bar {
	return h.bar
}
