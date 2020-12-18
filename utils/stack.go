package utils

type Stack struct {
	values []interface{}
}

func (s *Stack) Pop() interface{} {
	top := len(s.values) - 1
	val := s.values[top]
	s.values = s.values[:top]
	return val
}
func (s *Stack) Push(val interface{}) {
	s.values = append(s.values, val)
}
func (s *Stack) Peek() interface{} {
	if len(s.values) == 0 {
		return nil
	}
	return s.values[len(s.values)-1]
}
