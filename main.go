package main

import (
    "fmt"
    "skiplist/linkedlist"
)

func main() {
    l := linkedlist.New(2)
    l.Push(4)
    l.Push(6)
    fmt.Println("first", l.Value()) // 6
    fmt.Println("next", l.Next().Value()) // 4
    fmt.Println("last", l.Next().Next().Value()) // 2

    fmt.Println("Pop()")
    l.Pop()
    fmt.Println("first", l.Value()) // 4
    fmt.Println("next", l.Next().Value()) // 2
    fmt.Println("last", l.Next().Next()) // nil
}
