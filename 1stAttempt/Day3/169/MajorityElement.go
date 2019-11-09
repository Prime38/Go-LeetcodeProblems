package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	//Sol 1 -> using sort t-> O(nlogn) sp-> o(n) / O(1)
	l:=len(nums)
	sort.Ints(nums)
	return nums[l/2]
	//Sol 2 -> using map t-> O(n) sp-> o(n)
	m:=make(map[int]int)
	max:=len(nums)/2
	var  ans int
	for i:=0;i<len(nums);i++{
		m[nums[i]]++
		if max<m[nums[i]]{
			max=m[nums[i]]
			ans=nums[i]
		}
	}
	return ans





}
func main(){
	nums:=[]int {2,2,1,1,1,2,2}
	fmt.Println(majorityElement(nums))
}
