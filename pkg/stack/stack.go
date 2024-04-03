package stack

type ListInterface[T any] interface {
	Lenght() int
	Prepend(value T)
	Get(index int) T
	Delete(index int)
}

type Stack[T any] struct {
	list ListInterface[T]
}

func NewStack[T any](list ListInterface[T]) *Stack[T] {
	return &Stack[T]{list: list}
}

func (s *Stack[T]) Push(value T) {
	s.list.Prepend(value)
}

func (s *Stack[T]) Pop() T {
	if s.list.Lenght() == 0 {
		panic("pop from an empty stack")
	}
	val := s.list.Get(0)
	s.list.Delete(0)
	return val
}

func (s *Stack[T]) Length() int {
	return s.list.Lenght()
}

// добавление по указанному индексу (перед элементом)
//удаление по указанному индексу (с возвращением элемента)
// Не, нафиг стек тогда нужен?
