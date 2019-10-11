package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	ans:=0
	arr:=[128]bool{}
	for i:=0;i<len(J);i++{
		arr[int(J[i])-65]=true
	}
	for i:=0;i<len(S);i++{
		if arr[int(S[i])-65]==true{
			ans++
		}
	}
	return ans
}

func main(){
	J:="aAb"
	S:="aAAaAbbbb"
	fmt.Println(numJewelsInStones(J,S))

}
