package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head==nil{
		return nil
	}
	for head.Val==val{
		if head.Next!=nil{
			head=head.Next
		}else{
			return nil
		}
	}
	cur:=head
	for cur.Next!=nil{
		for cur.Next!=nil &&cur.Next.Val==val {
			cur.Next=cur.Next.Next
		}
		if cur.Next!=nil{
			cur=cur.Next
		}
	}
	return head

}