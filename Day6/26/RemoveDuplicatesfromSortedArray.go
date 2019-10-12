package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums)==0{
		return 0
	}
	i:=0
	for j:=1;i<len(nums);j++{
		if nums[j]!=nums[i]{
			i++
			nums[i]=nums[j]
		}
	}
	return i+1
}
func main(){
	nums:=[]int{1,2}
	fmt.Println(removeDuplicates(nums))

}
