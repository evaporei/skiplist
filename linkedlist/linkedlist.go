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
    head *Node[T]
}

func New[T any](value T) *LinkedList[T] {
    return &LinkedList[T] {
        head: new(value),
    }
}

func (l *LinkedList[T]) Push(value T) {
    oldNode := new(l.head.value)
    l.head.value = value
    l.head.next = oldNode
}

func (l *LinkedList[T]) Pop() {
    *l.head = *l.head.next
}

func (l *LinkedList[T]) Value() T {
    return l.head.value
}

func (l *LinkedList[T]) Next() *LinkedList[T] {
    return &LinkedList[T] {
        head: l.head.next,
    }
}
