package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head:=new(ListNode)
	if l1!=nil && l2==nil{
		return l1

	} else if l1==nil && l2!=nil{

		return l2

	} else if l1==nil && l2==nil{
		return l1

	}
	if l1.Val<=l2.Val{
		head=l1

		l1=l1.Next
	} else{
		head=l2
		l2=l2.Next
	}
	ans:=head
	for l1!=nil && l2!=nil{
		if l1.Val<=l2.Val{
			head.Next=l1
			l1=l1.Next
			head=head.Next
		} else{
			head.Next=l2
			l2=l2.Next
			head=head.Next
		}
	}
	if l1!=nil{
		head.Next=l1
		l1=l1.Next
		head=head.Next
	} else if l2!=nil{
		head.Next=l2
		l2=l2.Next
		head=head.Next
	}

	return ans

}
