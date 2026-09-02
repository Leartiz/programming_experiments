package main

import (
	"fmt"
	"unsafe"
)

// https://leetcode.com/problems/linked-list-cycle-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycleByAddr(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// The number of the nodes in the list is in the range [0, 10^4].
	// хеш тейбл, поможет посчитать сколько раз адрес повторяется...

	// map[*ListNode]int{}
	// map[<addr>] += 1

	a := head
	b := head.Next

	for b != nil {
		addrA := uintptr(unsafe.Pointer(a))
		addrB := uintptr(unsafe.Pointer(b))

		if addrA > addrB {
			return b
		}

		a = head.Next
		b = head.Next
	}

	return nil
}

func detectCycleByScan(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// The number of the nodes in the list is in the range [0, 10^4].
	// хеш тейбл, поможет посчитать сколько раз адрес повторяется...

	// map[*ListNode]int{}
	// map[<addr>] += 1

	a := head
	b := head.Next // !

	if b == nil {
		return nil
	}

	// 1. понять, есть ли вообще цикл

	// 2. завести еще один указатель, будет индикатором начала цикла
	//    идея такая, ставим на старт его. Как обнаружили цикл, смещаем его на 1,
	//    делаем полный прохол a и b, если пересеклись с mbStart, то он уже стоит на старте цикла.
	//    Иначе, mbStart = mbStart.Next, и повторяем прогон с a и b.

	for a != b {
		a = a.Next // точно не nil, так как b ушёл вперед

		b = b.Next
		if b == nil {
			return nil
		}

		b = b.Next // два шага
		if b == nil {
			return nil
		}
	}

	// цикл ТОЧНО найден
	// значит a и b точно не nil

	maybeStart := head
	copyA := a

	for {
		a = copyA
		a = a.Next

		isFound := a == maybeStart

		// полный цикл
		for !isFound && a != copyA {
			if a == maybeStart {
				isFound = true
				break
			}

			a = a.Next // 1 шаг
		}

		if isFound {
			break
		}

		maybeStart = maybeStart.Next
	}

	return maybeStart // ?
}

// detectCycleFloyd - Floyd: O(n) time, O(1) mem.
//
//	head ----μ----> entrance -------> ... -------> meeting
//	                   ^                              |
//	                   |<------------ λ-a ------------|
//	                   +-------- a (уже прошли в цикле)
//
//	пример [3,2,0,-4], вход = 2:
//
//	   3 ---> 2 ---> 0 ---> -4
//	          ^______________|
//	          entrance
//
//	фаза 1: a(+1), b(+2) с head -> встречаются внутри цикла (meeting)
//	        при этом μ == (шаги от meeting до entrance по циклу)
//	фаза 2: p с head и a с meeting, оба +1 -> первая встреча = entrance
func detectCycleFloyd(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// фаза 1: оба с head;
	// шаг, потом сравнение (классика Floyd)
	a, b := head, head

	for b != nil && b.Next != nil {
		a = a.Next
		b = b.Next.Next // всегда дальше?
		if a == b {
			break
		}
	}
	if b == nil || b.Next == nil {
		return nil // цикла нет
	}

	// фаза 2: вход в цикл;
	// с head и с точки встречи оба +1 -> первая встреча = start
	p := head
	for p != a {
		p = p.Next
		a = a.Next
	}
	return p
}

func main() {
	/*
		Input: head = [3,2,0,-4], pos = 1
		Output: tail connects to node index 1

		Input: head = [1,2], pos = 0
		Output: tail connects to node index 0

		Input: head = [1], pos = -1
		Output: no cycle
	*/
	{
		// [1], pos = 0
		n0 := &ListNode{Val: 1}
		n0.Next = n0 // !
		got := detectCycleFloyd(n0)
		fmt.Println(got == n0) // true
	}
	{
		// [3,2,0,-4], pos = 1
		n3 := &ListNode{Val: -4}
		n2 := &ListNode{Val: 0, Next: n3}
		n1 := &ListNode{Val: 2, Next: n2}
		n0 := &ListNode{Val: 3, Next: n1}
		n3.Next = n1 // cycle to index 1

		got := detectCycleFloyd(n0)
		fmt.Println(got == n1) // true
	}
	{
		// [1,2], pos = 0
		n1 := &ListNode{Val: 2}
		n0 := &ListNode{Val: 1, Next: n1}
		n1.Next = n0

		got := detectCycleFloyd(n0)
		fmt.Println(got == n0) // true
	}
	{
		// [1], pos = -1
		n0 := &ListNode{Val: 1}
		got := detectCycleFloyd(n0)
		fmt.Println(got == nil) // true
	}
}
