package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	a:=[]int{}
	a=append(a,heights...)
	ans:=0
	sort.Ints(heights)
	for i:=0;i<len(heights);i++{
		if a[i]!=heights[i]{
			ans++
		}
	}
	return ans

}
func  main()  {
	heights:=[]int{1,1,4,2,1,3}
	fmt.Println(heightChecker(heights))
}
