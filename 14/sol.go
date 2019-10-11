package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs)==0{
		return ""
	}
	minlen:=len(strs[0])
	LowStr:=strs[0]
	for i:=0;i<len(strs);i++{
		if minlen> len(strs[i]){
			minlen=len(strs[i])
			LowStr=strs[i]
		}
	}
	ans := 0
	for i:=0;i<minlen;i++{
		match:=true
		for j:=0;j<len(strs);j++{
			if strs[j][i]!=LowStr[i]{
				match=false
			}
		}
		if match==true{
			ans = ans + 1
		} else {
			break
		}
	}
	return LowStr[:ans]
}

func main(){
	strs:=[]string {}
	fmt.Println(longestCommonPrefix(strs))
}
