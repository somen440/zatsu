package di

type Hoge interface {
	HogeMethod() string
	Bar() Bar
}

type Bar interface {
	BarMethod() string
}
