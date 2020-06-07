package hoge

type Hoge struct{}

func New() *Hoge {
	return &Hoge{}
}

func (h *Hoge) HogeMethod() string {
	return "hoge method."
}
