package main

import "fmt"

func balancedStringSplit(s string) int {
	ans:=0
	cur:=0
	for i:=0;i<len(s);i++{
		if s[i]=='R'{
			cur++
		} else {
			cur--
		}
		if cur==0{
			ans++
		}
	}
	return ans
}

func main(){
	s:="RLRRLLRLRL"
	fmt.Println(balancedStringSplit(s))

}
