package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		head = head.Next
		return head
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}
func main() {
	a := &ListNode{
		Val:  5,
		Next: nil,
	}
	b := &ListNode{
		Val:  4,
		Next: a,
	}
	//c := &ListNode{
	//	Val:  3,
	//	Next: b,
	//}
	//d := &ListNode{
	//	Val:  2,
	//	Next: c,
	//}
	//e := &ListNode{
	//	Val:  1,
	//	Next: d,
	//}
	p := b
	for p != nil {
		fmt.Println(p)
		p = p.Next
	}
	fmt.Println("remove")
	h := removeNthFromEnd(b, 2)
	q := h
	for q != nil {
		fmt.Println(q)
		q = q.Next
	}
}
