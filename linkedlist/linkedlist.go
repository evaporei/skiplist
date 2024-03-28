package linkedlist

type LinkedList[T interface{}] struct {
    value T
    next *LinkedList[T]
}

func New[T interface{}](value T) *LinkedList[T] {
    return &LinkedList[T] {
        value: value,
        next: nil,
    }
}
