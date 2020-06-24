package main

import (
	"fmt"
	"container/list"
)

func main() {
	fmt.Println("package list")
	l := list.New()
	l.PushBack(2)
	l.PushFront(1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}