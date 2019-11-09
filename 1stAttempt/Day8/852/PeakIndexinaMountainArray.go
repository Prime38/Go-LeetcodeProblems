package main

import "fmt"

func peakIndexInMountainArray(A []int) int {
	low:=0
	high:=len(A)-1
	for low<high{
		mid:=low+(high-low)/2
		if A[mid]<A[mid+1]{
			low=mid+1
		} else{
			high=mid
		}
	}
	return low
}

func main(){
	A:=[]int{0,2,3,4,3,1,0}
	fmt.Println(peakIndexInMountainArray(A))
}
