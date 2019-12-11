package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices)<2{
		return 0
	}
	ans:=0
	for i:=0;i<len(prices)-1;{
		for i<len(prices)-1&& (prices[i+1]<=prices[i]){
			i++
		}
		if i==len(prices)-1{
			break
		}
		 buy:=i
		 i++
		 for i<len(prices)&&(prices[i]>=prices[i-1]){
		 	i++
		 }
		 sell:=i-1
		ans+=(prices[sell]-prices[buy])
	}
	return ans
}
func main(){
	prices:=[]int{7,6,4,3,1}
	fmt.Println(maxProfit(prices))
}
