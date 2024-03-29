package linkedlist

type LinkedList[T any] struct {
    Value T
    Next *LinkedList[T]
}

func New[T any](value T) *LinkedList[T] {
    return &LinkedList[T] {
        Value: value,
        Next: nil,
    }
}

func (l *LinkedList[T]) Push(value T) {
    oldNode := New(l.Value)
    l.Value = value
    l.Next = oldNode
}

func (l *LinkedList[T]) Pop() {
    *l = *l.Next
}
