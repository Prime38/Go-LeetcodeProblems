package main

import (
	"fmt"
)

func sortArrayByParity(A []int) []int {
	//sol 1 concating two array
	//odd:=[]int{}
	//even:=[]int{}
	//for i:=0;i<len(A);i++{
	//	if A[i]%2==0{
	//		even= append(even, A[i])
	//	} else{
	//		odd=append(odd,A[i])
	//	}
	//}
	//return append(even,odd...)
	//sol 2 sorting
	//for i:=0;i<len(A);i++{
	//		if A[i]%2==0{
	//			A[i]=-A[i]
	//		}
	//	}
	//	sort.Ints(A)
	//for i:=0;i<len(A);i++{
	//	if A[i]<0{
	//		A[i]=-A[i]
	//	} else {
	//		break
	//	}
	//}
	//return A
}
func main (){
	A:=[]int{3,1,2,4}
	fmt.Println(sortArrayByParity(A))

}
