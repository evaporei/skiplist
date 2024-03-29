package main

import (
    "fmt"
    "skiplist/linkedlist"
)

func main() {
    l := linkedlist.New(2)
    l.Push(4)
    fmt.Println("first", l) // { value: 4, next: { value: 2, next: nil } }
    fmt.Println("next", l.Next()) // { value: 2, next: nil }

    fmt.Println("Pop()")
    l.Pop()
    fmt.Println("first", l) // { value: 2, next: nil }
    fmt.Println("next", l.Next()) // nil
}
