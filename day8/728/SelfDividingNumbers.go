package main

import (
	"fmt"
	"strconv"
)

func selfDividingNumbers(left int, right int) []int {
	ans:=[]int{}
	for i:=left;i<=right;i++{
		if i%10==0{
			continue
		}
		num:=(strconv.Itoa(i))
		self:=true
		for j:=0;j<len(num);j++{
			dig:=int(num[j]-48)
			if dig==0{
				self=false
				break
			}

			if i % dig!=0  {
				self=false
				break
			}
		}
		if self==true{
			ans=append(ans,i)
		}
	}
	return ans
}

func main(){
	left:=1
	right:=156
	fmt.Println(selfDividingNumbers(left,right))

}
