package main

import "fmt"

/**
 * Definition for singly-linked list.

 */
type ListNode struct {
	Val int
	Next *ListNode
}
func (l *ListNode)addAtTail(){

}
func middleNode(head *ListNode) *ListNode {
	nums:=0
	cur:=head
	for cur!=nil{
		nums++
		cur=cur.Next
	}
	mid:=(nums/2)
	nums=0
	cur=head
	for cur!=nil{
		if nums==mid{
			return cur
		}
		nums++
		cur=cur.Next
	}
	return nil
}
func main(){
	listNode:=new(ListNode)
	fmt.Println(*listNode)

}
