package summator

import "fmt"

type Summator struct {
	base float64
}

func (s *Summator) Add(inc float64) float64 {
	s.base += inc
	return s.base
}

func (s Summator) String() string {
	return fmt.Sprintf("%f", s.base)
}
