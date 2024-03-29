package main

import (
    "fmt"
    "skiplist/linkedlist"
)

func main() {
    l := linkedlist.New(2)
    l.Push(4)
    fmt.Println("first", l)
    fmt.Println("next", l.Next())

    fmt.Println("Pop()")
    l.Pop()
    fmt.Println("first", l)
    fmt.Println("next", l.Next())
}
