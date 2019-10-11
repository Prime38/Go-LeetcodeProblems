package main

import (
	"fmt"
	"time"
)

func singleNumber(nums []int) int {
	//sol 1 time->O(n) space -> O(n)
	//m:=make(map[int]int)
	//min:=-199999999
	//for k,v:=range nums{
	//	if m[v]!=min{
	//		if m[v]==0{
	//			m[v]=k+1
	//		} else {
	//			nums[m[v]-1]=min
	//			nums[k]=min
	//		}
	//
	//	}
	//}
	//for i:=0;i<len(nums);i++{
	//	if nums[i]!=min{
	//		return nums[i]
	//	}
	//}
	//return 0

	////sol 2 time->O(n) space -> O(n)
	//m:=make(map[int]bool)
	//total:=0
	//twices:=0
	//for i:=0;i<len(nums);i++{
	//	total+=nums[i]
	//	if m[nums[i]]==true{
	//		twices+=2*nums[i]
	//	} else{
	//		m[nums[i]]=true
	//	}
	//}
	//return total-twices

	//sol 3 time->O(n) ,space -> O(1)
	ans:=0

	for i:=0;i<len(nums);i++{
		ans^=nums[i]
	}
	return ans


}


func main(){
	start := time.Now()
	nums:=[]int{4,1,2,1,2}
	fmt.Println(singleNumber(nums))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
