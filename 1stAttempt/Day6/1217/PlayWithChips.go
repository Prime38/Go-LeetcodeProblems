package main

import (
	"fmt"
	"math"
)

func minCostToMoveChips(chips []int) int {
	odds,evens:=0,0
	for _,v:=range chips{
		if v%2==0{
			evens++
		} else {
			odds++
		}
	}
	return int(math.Min(float64(odds), float64(evens)))
}

func main()  {
	chips:=[]int{2,2,2,3,3}
	fmt.Println(minCostToMoveChips(chips))

}
