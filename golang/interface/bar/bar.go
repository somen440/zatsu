package bar

type Bar struct{}

func New() *Bar {
	return &Bar{}
}

func (bar *Bar) BarMethod() string {
	return "bar method."
}
