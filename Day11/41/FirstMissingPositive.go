package main

import "fmt"

func firstMissingPositive(nums []int) int {
		for i:=0;i<len(nums);i++{
			if nums[i]<0 || nums[i]>len(nums){
				nums[i]=-1
			}
		}
		fmt.Println(nums)
		for i:=0;i<len(nums);i++{
			if nums[i]>0{
				cur:=nums[i]-1
				for nums[cur]>0{
					next:=nums[cur]-1
					nums[cur]=-1
					cur=next
				}
			}
		}
		fmt.Println(nums)
		for i:=0;i<len(nums);i++{
			if nums[i]>-1 {
				return i+1
			}
		}
		all:=true

	for i:=0;i<len(nums);i++{
		if nums[i]>-1 {
			all=false
		}
	}
	if all==true{
		return 1
	}

	return 0
}

func main(){
	nums:=[]int{7,8,9,11,12}
	fmt.Println(firstMissingPositive(nums))


}
