package main

import "testing"

func TestDetectCycleFloyd(t *testing.T) {
	{
		n0 := &ListNode{Val: 1}
		n0.Next = n0
		if got := detectCycleFloyd(n0); got != n0 {
			t.Fatal("self-loop")
		}
	}
	{
		n3 := &ListNode{Val: -4}
		n2 := &ListNode{Val: 0, Next: n3}
		n1 := &ListNode{Val: 2, Next: n2}
		n0 := &ListNode{Val: 3, Next: n1}
		n3.Next = n1
		if got := detectCycleFloyd(n0); got != n1 {
			t.Fatal("pos=1")
		}
	}
	{
		n1 := &ListNode{Val: 2}
		n0 := &ListNode{Val: 1, Next: n1}
		n1.Next = n0
		if got := detectCycleFloyd(n0); got != n0 {
			t.Fatal("pos=0")
		}
	}
	{
		n0 := &ListNode{Val: 1}
		if got := detectCycleFloyd(n0); got != nil {
			t.Fatal("no cycle")
		}
	}
}
