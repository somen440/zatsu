package structure

type Structure struct {
	Up string
}

func NewStructure(up string) *Structure {
	return &Structure{
		Up: up,
	}
}

func (s *Structure) GetUp() string {
	return s.Up
}
