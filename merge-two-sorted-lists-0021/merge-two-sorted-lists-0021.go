package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	head := &ListNode{
//		Val:  0,
//		Next: nil,
//	}
//	cur := head
//	for l1 != nil && l2 != nil{
//		if l1.Val <= l2.Val{
//			tmp := &ListNode{
//				Val:  l1.Val,
//				Next: nil,
//			}
//			cur.Next = tmp
//			cur = tmp
//			l1 = l1.Next
//		}else{
//			tmp := &ListNode{
//				Val:  l2.Val,
//				Next: nil,
//			}
//			cur.Next = tmp
//			cur = tmp
//			l2 = l2.Next
//		}
//	}
//	if l1 != nil{
//		cur.Next = l1
//	}
//	if l2 != nil{
//		cur.Next = l2
//	}
//	return head.Next
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func main() {
	a := &ListNode{
		Val:  4,
		Next: nil,
	}
	b := &ListNode{
		Val:  2,
		Next: a,
	}
	c := &ListNode{
		Val:  1,
		Next: b,
	}

	aa := &ListNode{
		Val:  5,
		Next: nil,
	}
	//bb := &ListNode{
	//	Val:  3,
	//	Next: aa,
	//}
	//cc := &ListNode{
	//	Val:  1,
	//	Next: bb,
	//}
	q := mergeTwoLists(c, aa)
	for q != nil {
		fmt.Println(q)
		q = q.Next
	}
}
