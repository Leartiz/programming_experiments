package n4

type Stack[T any] []T

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() T {
	length := len(*s)
	value := (*s)[length-1]
	*s = (*s)[:length-1]
	return value
}
