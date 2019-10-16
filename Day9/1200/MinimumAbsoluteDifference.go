package main

import (
	"fmt"
	"sort"
)

func abs(x int) int{
	if x<0{
		return -x
	}
	return x
}
func minimumAbsDifference(arr []int) [][]int {

	sort.Ints(arr)
	ans:=[][]int {}
	mini:=2000002
	for i:=0;i<len(arr)-1;i++{
		sub:=arr[i+1]-arr[i]
		if abs(sub) < mini{
			mini=abs(sub)
		}
	}
	for i:=0;i<len(arr)-1;i++{
		sub:=arr[i+1]-arr[i]
		if abs(sub) == mini{
			ans=append(ans,[]int{arr[i],arr[i+1]})
		}
	}
	return ans
}
func main(){
	arr:=[]int{4,2,1,3}
	fmt.Println(minimumAbsDifference(arr))
}
