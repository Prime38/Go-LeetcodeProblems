package main

import "fmt"

func twoSum(nums []int, target int) []int {
	ans:=[]int{}
	m:=make(map[int]int)
	for k,v:=range nums{
		a, ok:=m[target - v]
		fmt.Println(a,ok)
		if ok{
			ans= append(ans, a)
			ans= append(ans, k)
			return ans
		}
		m[v]=k
	}
	return nil
}

func main(){
	nums:=[]int {6,2,7,11,15}
	target:=9
	fmt.Println(twoSum(nums,target))

}
