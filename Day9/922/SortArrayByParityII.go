package main

import "fmt"

func sortArrayByParityII(A []int) []int {
	odd:=[]int{}
	even:=[]int{}
	for i:=0;i<len(A);i++{
		if A[i]%2==0{
			even=append(even,A[i])
		} else{
			odd=append(odd,A[i])
		}
	}
	A=[]int{}
	for i:=0;i<len(even);i++{
		A=append(A,even[i])
		A=append(A,odd[i])
	}
	return A
}
func main(){
	A:=[]int{4,2,5,7}
	fmt.Println(sortArrayByParityII(A))
}
