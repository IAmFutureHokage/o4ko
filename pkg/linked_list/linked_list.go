package linkedlist

// Нода
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

//Заголовок
type LinkedList[T any] struct {
	head *Node[T]
	size int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Lenght() int {
	return l.size
}

func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.size++
}

func (l *LinkedList[T]) Prepend(value T) {
	l.head = &Node[T]{Value: value, Next: l.head}
	l.size++
}

func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.size = 0
}

func (l *LinkedList[T]) ToSlice() []T {
	slice := make([]T, 0, l.size)
	current := l.head
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}
	return slice
}

func (l *LinkedList[T]) Get(index int) T {
	if index < 0 || index >= l.size {
		panic("index out of bounds")
	}
	current := l.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (l *LinkedList[T]) Delete(index int) {
	if index < 0 || index >= l.size {
		panic("index out of bounds")
	}
	if index == 0 {
		l.head = l.head.Next
		l.size--
		return
	}
	current := l.head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
	l.size--
}
