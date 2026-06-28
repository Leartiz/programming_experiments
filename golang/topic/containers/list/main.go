package main

import (
	"container/list"
	"fmt"
)

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v -> ", e.Value.(string))
	}
	fmt.Printf("nil\n")
}

func main() {
	{
		l := list.New()

		l.PushBack("middle")
		l.PushFront("first")
		l.PushBack("last")

		fmt.Println("len:", l.Len())
		fmt.Println("list:")
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}

		fmt.Printf("front value: %v\n", l.Front().Value)
		fmt.Printf("back value: %v\n", l.Back().Value)

		for e := l.Front(); e != nil; e = e.Next() {
			if e.Value == "middle" {
				l.Remove(e)
				break
			}
		}

		fmt.Println("len:", l.Len())
		fmt.Println("after remove:")
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
	}

	// ***

	{
		l := list.New()
		elementA := l.PushBack("A")
		l.PushBack("B")
		l.PushBack("C")
		printList(l)

		l.MoveToBack(elementA)
		printList(l)
	}

	{
		l := list.New()

		elementA := l.PushBack("A")
		l.PushBack("B")
		l.PushBack("C")
		printList(l)

		l1 := list.New()
		elementA1 := l1.PushBack("A")
		l1.PushBack("B")
		l1.PushBack("C")
		printList(l1)

		_ = elementA
		l.MoveToBack(elementA1) // e.list != l
		printList(l)
	}
}
