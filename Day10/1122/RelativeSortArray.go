package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m:=make(map[int]int,len(arr1))
	for i:=0;i<len(arr1);i++{
		m[arr1[i]]++
	}
	arr1=[]int{}
	for i:=0;i<len(arr2);i++{
		for m[arr2[i]]>0{
			arr1=append(arr1,arr2[i])
			m[arr2[i]]--
		}
	}
	arr2=[]int{}
	for i,v:=range m{
		for v>0{
			arr2=append(arr2,i)
			v--
		}
	}
	sort.Ints(arr2)
	arr1=append(arr1,arr2...)
	return arr1
}


func main(){
	arr1:=[]int{2,3,1,3,2,4,6,7,9,2,19}
	arr2:=[]int{2,1,4,3,9,6}
	fmt.Println(relativeSortArray(arr1,arr2))

}
