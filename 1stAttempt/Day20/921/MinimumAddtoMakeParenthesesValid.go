package main

import (

	"fmt"
)

func minAddToMakeValid(S string) int {
	if len(S)<2{
		return len(S)
	}

	st:=new()
	for i:=0;i<len(S) ;i++  {
		if S[i]==')'&

	}

	return 0
}
func main(){
	str:="()))(("
	fmt.Println(minAddToMakeValid(str))
}
