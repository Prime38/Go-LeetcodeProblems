package main

import (
	"fmt"
)
type fs struct{
	pos int
	char uint8
}
func minWindow(s string, t string) string {
	if len(s)<len(t){
		return ""
	}
	Ans:=s
	m:=[130]int{}
	ms:=[130]int{}
	for i:=0;i<len(t);i++{
		m[t[i]]++
	}
	arr:=[]fs{}
	for i:=0;i<len(s);i++  {
		ms[s[i]]++
		if m[s[i]]>0{
			temp:=new(fs)
			temp.pos=i
			temp.char=s[i]
			arr=append(arr,*temp)
		}
	}

	for k,_:=range (m){
		if ms[k]<m[k]{
			return ""
		}
	}

	for i:=0;i<len(arr);i++{
		mk:=[130]int{}
		for k,val:=range m{
			mk[k]=val
		}
		window:=""
		cnt:=0
		st:=arr[i].pos
		end:=arr[i].pos
		for j:=i;j<len(arr);j++{
			if mk[arr[j].char]>0{
				mk[arr[j].char]--
				cnt++
				end=arr[j].pos
			}
			if cnt==len(t){
				window=s[st:end+1]
				break
			}
		}
		if cnt==len(t)&&len(window)<len(Ans){
			Ans=window
		}
	}
	return Ans
}

func main(){
	s:="a"
	t:="b"
	fmt.Println(minWindow(s,t))
}
