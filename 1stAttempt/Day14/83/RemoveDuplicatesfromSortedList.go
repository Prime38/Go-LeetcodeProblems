package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
type ListNode struct {
    Val int
     Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head==nil{
		return head
	}
	ans:=head
	for head.Next!=nil{
		if head.Val==head.Next.Val{
			head.Next=head.Next.Next
		}
		if head.Next==nil{
			break
		}
		if head.Val!=head.Next.Val {
			head=head.Next

		}
	}
	return ans

}