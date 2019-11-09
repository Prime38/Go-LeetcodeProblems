package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct{
	Val int
	Next *ListNode
}
//iterative Solution
func reverseList(head *ListNode) *ListNode {
	prev:=new(ListNode)
	prev=nil
	cur:=head
	for cur!=nil{
		nextTemp:=cur.Next
		cur.Next=prev
		prev=cur
		cur=nextTemp
	}
	return prev
}
