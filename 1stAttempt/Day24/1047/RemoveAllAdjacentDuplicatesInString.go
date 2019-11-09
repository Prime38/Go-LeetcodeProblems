package main

import (
	"fmt"
)
func removeDuplicates(S string) string{
	S="."+S
	ans:=S[:1]

	pos:=0
	for i:=1;i<len(S);i++{
		if S[i]==ans[pos]{
			if pos>0{
				ans=ans[:pos]
				pos--
			} else {
				ans=""
			}
		} else{
			ans+=string(S[i])
			pos++
		}
	}
	return ans[1:]
}
func main(){
	st:="aabbcc"
	fmt.Println(removeDuplicates(st))
}
