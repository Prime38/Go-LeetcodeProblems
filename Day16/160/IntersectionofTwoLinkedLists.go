package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */



func getIntersectionNode(headA, headB *ListNode) *ListNode {
	indexA:=0
	indexB:=0
	curA:=headA
	curB:=headB
	for curA!=nil{
		indexA++
		curA=curA.Next
	}
	for curB!=nil{
		indexB++
		curB=curB.Next
	}
	curA=headA
	curB=headB
	cnt:=0
	if indexA>indexB{
		dif:=indexA-indexB
		for cnt<dif{
			curA=curA.Next
			cnt++
		}
	} else if indexA<indexB{
		dif:=indexB-indexA
		for cnt<dif{
			curB=curB.Next
			cnt++
		}
	}
	for curA!=curB{
		curA=curA.Next
		curB=curB.Next
	}
	return curA
}