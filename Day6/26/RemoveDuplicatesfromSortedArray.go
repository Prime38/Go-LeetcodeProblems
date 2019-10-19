package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums)==0{
		return 0
	}
	//sol 1
	//now:=nums[0]
	//for i:=1;i<len(nums);i++{
	//	fmt.Println("i =",i)
	//	for nums[i]==now{fmt.Println(nums[i],now)
	//		if i+1<len(nums){
	//			nums=append(nums[:i], nums[i+1:]...)
	//			fmt.Println("inside if",nums)
	//		} else{
	//			nums=nums[:i]
	//			fmt.Println("inside else",nums)
	//			break
	//		}
	//	}
	//	if i<len(nums){
	//		now=nums[i]
	//	}
	//}
	//return len(nums)

	//sol 2
	j:=0
	for i:=1;i<len(nums);i++{
		if nums[i]!=nums[j]{
			j++
			nums[j]=nums[i]
		}
	}
	nums=nums[:j+1]
	fmt.Println(nums)
	return j+1
}
func main(){
	nums:=[]int{0,0,1,1,1,2,2,3,3,4,4}
	fmt.Println(removeDuplicates(nums))

}
