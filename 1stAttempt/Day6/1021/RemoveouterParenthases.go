package main

import "fmt"

func removeOuterParentheses(S string) string {
	res:=""
	ans,st:=0,0
	for i:=0;i<len(S);i++{
		if S[i]=='('{
			ans++

		} else {
			ans--
		}
		if ans==0{

			res+=S[st+1:i]
			st=i+1
		}
	}
	return res

}
func main()  {
	S:="(()())(())(()(()))"
	fmt.Println(removeOuterParentheses(S))

}
