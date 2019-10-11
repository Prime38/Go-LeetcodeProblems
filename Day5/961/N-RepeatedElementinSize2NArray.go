package main

import (
	"fmt"
)

func repeatedNTimes(A []int) int {
	// sol 1
	sort.Ints(A)
	l:=len(A)/2
	if A[l-1]==A[l]{
		return A[l]
	} else {
		if A[l-1]==A[l-2]{
			return A[l-1]
		} else if A[l]==A[l+1] {
			return A[l]
		}
	}
	return 0
	// sol 1 using map
	m:=make(map[int]int)
	for i:=0;i<len(A);i++{
		if m[A[i]]>0{
			return A[i]
		} else {
			m[A[i]]++
		}
	}
	return 0
	//sol 2 using a trick
	for i:=1;i<4;i++{
		for j:=0;j<len(A)-i;j++{
			if A[j]==A[i+j]{
				return A[j]
			}
		}
	}
	return 0
}

func main(){
	A:=[]int{2,1,5,3,2,2}
	fmt.Println(repeatedNTimes(A))

}
