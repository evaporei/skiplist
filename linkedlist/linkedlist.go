package linkedlist

type Node[T any] struct {
    value T
    next *Node[T]
}

func new[T any](value T) *Node[T] {
    return &Node[T] {
        value: value,
        next: nil,
    }
}

type LinkedList[T any] struct {
    last *Node[T]
}

func New[T any](value T) *LinkedList[T] {
    return &LinkedList[T] {
        last: new(value),
    }
}

func (l *LinkedList[T]) Push(value T) {
    oldNode := new(l.last.value)
    l.last.value = value
    l.last.next = oldNode
}

func (l *LinkedList[T]) Pop() {
    *l.last = *l.last.next
}

func (l *LinkedList[T]) Value() T {
    return l.last.value
}

func (l *LinkedList[T]) Next() *LinkedList[T] {
    return &LinkedList[T] {
        last: l.last.next,
    }
}
