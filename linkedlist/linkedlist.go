package linkedlist

type LinkedList[T interface{}] struct {
    Value T
    Next *LinkedList[T]
}

func New[T interface{}](value T) *LinkedList[T] {
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
