package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	//ans:=[]int{}
	//Sol 1
	//for i:=0;i<len(nums);i++{
	//	if nums[i]==0{
	//		continue
	//	}
	//	next:=nums[i]-1
	//	cur:=i
	//	for nums[next]>0{
	//
	//		cur=next
	//		next=nums[next]-1
	//		nums[cur]=0
	//		fmt.Println(cur,next)
	//	}
	//	fmt.Println(nums)
	//}
	// sol 2
	//for i:=0;i<len(nums);i++{
	//	if nums[i]!=0{
	//		ans=append(ans,i+1)
	//	}
	//}

	//for i:=0;i<len(nums);i++{
	//	t:=int(math.Abs(float64(nums[i]))-1)
	//	//fmt.Println(t)
	//	nums[t]= int(-math.Abs(float64(nums[t])))
	//}
	//fmt.Println(nums)
	//for i:=0;i<len(nums);i++{
	//	if nums[i]>0{
	//		f:=i+1
	//		ans=append(ans,f)
	//	}
	//}
	//return ans
	//sol 3
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums)
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			fmt.Println(nums)
		}
	}

	res := make([]int, 0, len(nums))

	for i, n := range nums {
		if n != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}
func main(){
	nums:=[]int{8,6,4,2,1,3,5,7}
	//for i:=0; i <len(nums); i++ {
	//	print(nums[i])
	//}
	x:=[]int{}
	x = findDisappearedNumbers(nums)
	fmt.Println(x)


}