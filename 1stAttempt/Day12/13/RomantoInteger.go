package main

import "fmt"

func romanToInt(s string) int {
	m:=make(map[byte]int)
	ans:=0
	m['I']=1
	m['V']=5
	m['X']=10
	m['L']=50
	m['C']=100
	m['D']=500
	m['M']=1000
	for i:=0;i<len(s);i++{
		if i==len(s)-1{
			ans+=m[s[i]]
		} else{
			if m[s[i]]<m[s[i+1]]{
				ans-=m[s[i]]
			} else{
				ans+=m[s[i]]
			}
		}
	}
	return ans
}

func main(){


	s:="LVIII"
	fmt.Println(romanToInt(s))

}
