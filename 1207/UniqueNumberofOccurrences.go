package main

import (
	"fmt"
	"sort"
)

func uniqueOccurrences(arr []int) bool {
	m:=make(map[int]int)
	for i:=0;i<len(arr);i++{
		m[arr[i]]++
	}
	arr=[]int{}
	for _,v:=range m{
		arr=append(arr,v)
	}
	sort.Ints(arr)

	set:=[]int{}
	set=append(set,arr[0])
	for i:=1;i<len(arr);i++{
		if arr[i]!=set[len(set)-1]{
			set=append(set,arr[i])
		}
	}
	return (len(arr)==len(set))
}
func main(){
	arr:=[]int{-3,0,1,-3,1,1,1,-3,10,0}
	fmt.Println(uniqueOccurrences(arr))

}
