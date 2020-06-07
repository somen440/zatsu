package bar

import (
	"github.com/somen440/zatsu/golang/interface/di"
)

type Bar struct{}

var _ di.Bar = &Bar{}

func New() *Bar {
	return &Bar{}
}

func (bar *Bar) BarMethod() string {
	return "bar method."
}
