package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	//     for i:=0; i<len(nums);i++{

	//         for j:=1; j<=k ;j++{
	//             if i+j < len(nums)  {
	//                 if nums[i] == nums[i+j]{
	//                     return true
	//                 }
	//         }
	//         }
	//     }
	//     return false

	m := make(map[int]int)

	for i:=0; i<len(nums); i++{
		index,ok:= m[nums[i]]
		if ok{
			if  i-index <= k{
				return true
			}else{
				m[nums[i]] = i
			}
		}else{
			m[nums[i]] = i
		}
	}
	return false
}
func main(){
	//nums = [1,0,1,1], k = 1
	nums := []int{1,2,3,1,2,3}
	k := 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}
