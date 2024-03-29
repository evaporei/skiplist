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
    node := new(value)
    return &LinkedList[T] {
        head: node,
    }
}

// O(n)
func (l *LinkedList[T]) Push(value T) {
    curr := l.head

    for curr.next != nil {
        curr = curr.next
    }

    newNode := new(value)
    curr.next = newNode
}

// O(n)
func (l *LinkedList[T]) Pop() {
    curr := l.head
    var prev *Node[T] = nil

    for curr.next != nil {
        prev = curr
        curr = curr.next
    }

    prev.next = nil
}

func (l *LinkedList[T]) Value() T {
    return l.head.value
}

func (l *LinkedList[T]) Next() *LinkedList[T] {
    return &LinkedList[T] {
        head: l.head.next,
    }
}

// O(n)
func (l *LinkedList[T]) Last() *LinkedList[T] {
    curr := l.head

    for curr.next != nil {
        curr = curr.next
    }

    return &LinkedList[T] {
        head: curr,
    }
}
