package linkedlist

type LinkedList[T any] struct {
    Value T
    next *LinkedList[T]
}

func New[T any](value T) *LinkedList[T] {
    return &LinkedList[T] {
        Value: value,
        next: nil,
    }
}

func (l *LinkedList[T]) Push(value T) {
    oldNode := New(l.Value)
    l.Value = value
    l.next = oldNode
}

func (l *LinkedList[T]) Pop() {
    *l = *l.next
}

func (l *LinkedList[T]) Next() *LinkedList[T] {
    return l.next
}
