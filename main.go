package main

import (
    "fmt"
    "skiplist/linkedlist"
)

func main() {
    l := linkedlist.New(2)
    l.Push(4)
    fmt.Println("first", l.Value()) // 4
    fmt.Println("next", l.Next().Value()) // 2

    fmt.Println("Pop()")
    l.Pop()
    fmt.Println("first", l.Value()) // 2
    fmt.Println("next", l.Next()) // nil
}
