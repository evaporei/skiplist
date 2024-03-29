package linkedlist

type LinkedList[T any] struct {
    value T
    next *LinkedList[T]
}

func New[T any](value T) *LinkedList[T] {
    return &LinkedList[T] {
        value: value,
        next: nil,
    }
}

func (l *LinkedList[T]) Push(value T) {
    oldNode := New(l.value)
    l.value = value
    l.next = oldNode
}

func (l *LinkedList[T]) Pop() {
    *l = *l.next
}

func (l *LinkedList[T]) Value() T {
    return l.value
}

func (l *LinkedList[T]) Next() *LinkedList[T] {
    return l.next
}
